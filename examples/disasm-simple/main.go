package main

import (
	"bytes"
	"fmt"
	"unsafe"

	"github.com/zyantific/zydis-go"
)

func main() {
	data := []byte{
		0x51, 0x8D, 0x45, 0xFF, 0x50, 0xFF, 0x75, 0x0C, 0xFF, 0x75,
		0x08, 0xFF, 0x15, 0xA0, 0xA5, 0x48, 0x76, 0x85, 0xC0, 0x0F,
		0x88, 0xFC, 0xDA, 0x02, 0x00,
	}

	// The runtime address (instruction pointer) was chosen arbitrarily here in order to better
	// visualize relative addressing. In your actual program, set this to e.g. the memory address
	// that the code being disassembled was read from.
	runtimeAddress := uintptr(0x007FFFFFFF400000)

	// Loop over the instructions in our buffer.
	offset := 0
	insn := zydis.DisassembledInstruction{}
	for offset < len(data) {
		status := zydis.DisassembleIntel(
			zydis.MACHINE_MODE_LONG_64,
			uint64(runtimeAddress),
			unsafe.Pointer(&data[offset]),
			uint64(len(data)-offset),
			&insn,
		)
		if !zydis.Ok(status) {
			break
		}

		textEnd := bytes.IndexByte(insn.Text[:], 0)
		fmt.Printf("%016X  %s\n", runtimeAddress, string(insn.Text[:textEnd]))
		offset += int(insn.Info.Length)
		runtimeAddress += uintptr(insn.Info.Length)
	}
}
