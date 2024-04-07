package main

import (
	"fmt"
	"os"
	"time"
	"unsafe"

	"github.com/can1357/zydis-go"
)

const DISAS_BENCH_NO_FORMAT = true
const DISAS_BENCH_DECODE_MINIMAL = true

func main() {
	fmt.Printf("Zydis Version: %x\n", zydis.GetVersion())

	formatter := zydis.Formatter{}
	if !DISAS_BENCH_NO_FORMAT {
		if zydis.Failed(formatter.Init(zydis.FORMATTER_STYLE_INTEL)) {
			fmt.Println("Unable to initialize instruction formatter")
			os.Exit(1)
		}
	}

	loopCount := 4
	code, err := os.ReadFile("xul.dll")
	if err != nil {
		fmt.Println("Unable to read input file")
		os.Exit(1)
	}

	dec := zydis.Decoder{}
	dec.Init(zydis.MACHINE_MODE_LONG_64, zydis.STACK_WIDTH_64)

	if DISAS_BENCH_DECODE_MINIMAL {
		dec.EnableMode(zydis.DECODER_MODE_MINIMAL, zydis.Bool(1))
	}

	numValidInsns := 0
	numBadInsn := 0
	readOffs := 0
	startTime := time.Now()
	for i := 0; i < loopCount; i++ {
		readOffs = 0

		info := zydis.DecodedInstruction{}
		ctx := zydis.DecoderContext{}
		ops := [5]zydis.DecodedOperand{}
		printBuffer := [256]byte{}
		for readOffs < len(code) {
			status := dec.DecodeInstruction(
				&ctx,
				unsafe.Pointer(&code[readOffs]), uint64(len(code)-readOffs),
				&info,
			)
			if status == zydis.STATUS_OUT_OF_RANGE {
				break
			}
			if !zydis.Ok(status) {
				readOffs++
				numBadInsn++
				continue
			}

			readOffs += int(info.Length)
			numValidInsns++

			if !DISAS_BENCH_NO_FORMAT {
				res := formatter.FormatInstruction(
					&info,
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
	mbPerMs := float64(len(code)*loopCount) / (1024 * 1024) / float64(endTime.Sub(startTime).Milliseconds())
	fmt.Printf("Speed: %.2f MB/s\n", mbPerMs*1000)
}
