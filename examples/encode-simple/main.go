package main

import (
	"fmt"
	"unsafe"

	"github.com/zyantific/zydis-go"
)

func main() {
	req := zydis.EncoderRequest{
		Mnemonic:     zydis.MNEMONIC_MOV,
		MachineMode:  zydis.MACHINE_MODE_LONG_64,
		OperandCount: 2,
		Operands: [5]zydis.EncoderOperand{
			{Type: zydis.OPERAND_TYPE_REGISTER},
			{Type: zydis.OPERAND_TYPE_IMMEDIATE},
		},
	}
	req.Operands[0].Reg.Value = zydis.REGISTER_RAX
	req.Operands[1].Imm.SetU(0x1337)

	encodedInsn := [16]byte{}
	encodedLength := uint64(len(encodedInsn))
	if !zydis.Ok(req.EncodeInstruction(unsafe.Pointer(&encodedInsn[0]), &encodedLength)) {
		fmt.Println("Failed to encode instruction")
		return
	}

	for i := uint64(0); i < encodedLength; i++ {
		fmt.Printf("%02X ", encodedInsn[i])
	}
	fmt.Println("")
}
