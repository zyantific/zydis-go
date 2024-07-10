package io

import (
	"fmt"
	"io"
	"unsafe"

	"github.com/zyantific/zydis-go"
)

type (
	Encoder struct {
		writer io.Writer
	}
)

func NewEncoder(writer io.Writer) *Encoder {
	return &Encoder{
		writer: writer,
	}
}

func (encoder *Encoder) Encode(from *zydis.EncoderRequest) error {
	buf := make([]byte, zydis.MAX_INSTRUCTION_LENGTH)
	len := zydis.MAX_INSTRUCTION_LENGTH

	status := from.EncodeInstruction(unsafe.Pointer(unsafe.SliceData(buf)), &len)
	if !zydis.Ok(status) {
		return fmt.Errorf("%w: %d", ErrWrongStatus, status)
	}

	_, err := encoder.writer.Write(buf[:len])
	if err != nil {
		return err
	}

	return nil
}
