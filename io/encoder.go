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
	buffer := [zydis.MAX_INSTRUCTION_LENGTH]byte{}
	length := uint64(len(buffer))

	status := from.EncodeInstruction(unsafe.Pointer(&buffer), &length)
	if !zydis.Ok(status) {
		return fmt.Errorf("%w: %d", ErrWrongStatus, status)
	}

	_, err := encoder.writer.Write(buffer[:length])
	if err != nil {
		return err
	}

	return nil
}

func AppendRequest(to []byte, from *zydis.EncoderRequest) ([]byte, error) {
	buffer := [zydis.MAX_INSTRUCTION_LENGTH]byte{}
	length := uint64(len(buffer))

	status := from.EncodeInstruction(unsafe.Pointer(&buffer), &length)
	if !zydis.Ok(status) {
		return nil, fmt.Errorf("%w: %d", ErrWrongStatus, status)
	}

	return append(to, buffer[:length]...), nil
}
