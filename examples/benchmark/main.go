package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
	"unsafe"

	"github.com/zyantific/zydis-go"
)

const DISAS_BENCH_NO_FORMAT = true
const DISAS_BENCH_DECODE_MINIMAL = true

var (
	fProfile    = flag.Bool("cpuprofile", false, "write cpu profile to file")
	fLoopCount  = flag.Int("loops", 5, "number of loops")
	fInputFile  = flag.String("input", "xul.dll", "input file")
	fCodeOffset = flag.Int("offset", 0x400, "code offset")
	fCodeSize   = flag.Int("size", 0x2460400, "code size")
)

func readFile() []byte {
	file, err := os.Open(*fInputFile)
	if err != nil {
		log.Fatalf("Unable to open input file: %v", err)
	}
	defer file.Close()
	if _, err := file.Seek(int64(*fCodeOffset), 0); err != nil {
		log.Fatalf("Unable to seek input file: %v", err)
	}
	data := make([]byte, *fCodeSize)
	if _, err := file.Read(data); err != nil {
		log.Fatalf("Unable to read input file: %v", err)
	}
	return data
}

func main() {
	fmt.Printf("Zydis Version: %x\n", zydis.GetVersion())

	flag.Parse()
	if *fProfile {
		profile, _ := os.Create("default.pgo")
		pprof.StartCPUProfile(profile)
		defer pprof.StopCPUProfile()
		defer profile.Close()
	}

	code := readFile()

	formatter := zydis.Formatter{}
	if !DISAS_BENCH_NO_FORMAT {
		if zydis.Failed(formatter.Init(zydis.FORMATTER_STYLE_INTEL)) {
			fmt.Println("Unable to initialize instruction formatter")
			os.Exit(1)
		}
	}

	dec := zydis.Decoder{}
	dec.Init(zydis.MACHINE_MODE_LONG_64, zydis.STACK_WIDTH_64)
	if DISAS_BENCH_DECODE_MINIMAL {
		dec.EnableMode(zydis.DECODER_MODE_MINIMAL, zydis.Bool(1))
	}

	numValidInsns := 0
	numBadInsn := 0
	startTime := time.Now()
	loopCount := *fLoopCount
	for i := 0; i < loopCount; i++ {
		readOffs := 0

		info := &zydis.DecodedInstruction{}
		ctx := &zydis.DecoderContext{}
		ops := [5]zydis.DecodedOperand{}
		printBuffer := [256]byte{}
		for {
			bytesLeft := len(code) - readOffs
			if bytesLeft <= 0 {
				break
			}
			status := dec.DecodeInstruction(
				ctx,
				unsafe.Pointer(&code[readOffs]),
				uint64(bytesLeft),
				info,
			)
			if status == zydis.STATUS_NO_MORE_DATA {
				break
			}
			if !zydis.Ok(status) {
				readOffs++
				numBadInsn++
				continue
			}
			numValidInsns++
			readOffs += int(info.Length)

			if !DISAS_BENCH_NO_FORMAT {
				res := formatter.FormatInstruction(
					info,
					&ops[0], 5,
					&printBuffer[0], 256,
					uint64(readOffs),
					nil,
				)
				if zydis.Failed(res) {
					fmt.Println("Unable to format instruction")
					os.Exit(1)
				}
			}
		}
	}

	endTime := time.Now()
	fmt.Printf("Disassembled %d instructions (%d valid, %d bad), %.2f ms\n",
		numValidInsns+numBadInsn, numValidInsns, numBadInsn,
		endTime.Sub(startTime).Seconds()*1000)

	// Calculate MB/s
	dataSizeMb := float64(len(code)*loopCount) / (1024 * 1024)
	timeSpentSec := float64(endTime.Sub(startTime)) / float64(time.Second)
	fmt.Printf("Speed: %.2f MB/s\n", dataSizeMb/timeSpentSec)
}
