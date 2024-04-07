package zydis

import (
	"fmt"
	"unsafe"

	"github.com/can1357/zydis-go/zyembed"
)

func init() {
	if zyembed.Data != nil {
		GengoLibrary.LoadEmbed(zyembed.Data)
	} else {
		fmt.Println("zydis-go: zyembed.Data is nil")
	}
}

func (s ShortString) String() string {
	return unsafe.String((*byte)(unsafe.Pointer(s.Data())), s.Size())
}

func Ok(status uint32) bool {
	return status&0x80000000 == 0
}
func Failed(status uint32) bool {
	return status&0x80000000 != 0
}
func MakeStatus(err uint32, module uint32, code uint32) uint32 {
	return (err & 0x01 << 31) | (module & 0x7FF << 20) | (code & 0xFFFFF)
}
func StatusModule(status uint32) uint32 {
	return (status >> 20) & 0x7FF
}
func StatusCode(status uint32) uint32 {
	return status & 0xFFFFF
}

const ZYAN_TRUE = 1
const ZYAN_FALSE = 0

const (
	MODULE_ZYCORE   = 0x001
	MODULE_ARGPARSE = 0x003
	MODULE_USER     = 0x3FF
)

var (
	STATUS_SUCCESS                  = MakeStatus(0, MODULE_ZYCORE, 0x00)
	STATUS_FAILED                   = MakeStatus(1, MODULE_ZYCORE, 0x01)
	STATUS_TRUE                     = MakeStatus(0, MODULE_ZYCORE, 0x02)
	STATUS_FALSE                    = MakeStatus(0, MODULE_ZYCORE, 0x03)
	STATUS_INVALID_ARGUMENT         = MakeStatus(1, MODULE_ZYCORE, 0x04)
	STATUS_INVALID_OPERATION        = MakeStatus(1, MODULE_ZYCORE, 0x05)
	STATUS_ACCESS_DENIED            = MakeStatus(1, MODULE_ZYCORE, 0x06)
	STATUS_NOT_FOUND                = MakeStatus(1, MODULE_ZYCORE, 0x07)
	STATUS_OUT_OF_RANGE             = MakeStatus(1, MODULE_ZYCORE, 0x08)
	STATUS_INSUFFICIENT_BUFFER_SIZE = MakeStatus(1, MODULE_ZYCORE, 0x09)
	STATUS_NOT_ENOUGH_MEMORY        = MakeStatus(1, MODULE_ZYCORE, 0x0A)
	STATUS_BAD_SYSTEMCALL           = MakeStatus(1, MODULE_ZYCORE, 0x0B)
	STATUS_OUT_OF_RESOURCES         = MakeStatus(1, MODULE_ZYCORE, 0x0C)
	STATUS_MISSING_DEPENDENCY       = MakeStatus(1, MODULE_ZYCORE, 0x0D)
	STATUS_ARG_NOT_UNDERSTOOD       = MakeStatus(1, MODULE_ARGPARSE, 0x00)
	STATUS_TOO_FEW_ARGS             = MakeStatus(1, MODULE_ARGPARSE, 0x01)
	STATUS_TOO_MANY_ARGS            = MakeStatus(1, MODULE_ARGPARSE, 0x02)
	STATUS_ARG_MISSES_VALUE         = MakeStatus(1, MODULE_ARGPARSE, 0x03)
	STATUS_REQUIRED_ARG_MISSING     = MakeStatus(1, MODULE_ARGPARSE, 0x04)
)
