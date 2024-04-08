package zydis

import (
	"fmt"
	"unsafe"

	"github.com/zyantific/zydis-go/zyembed"
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

const (
	TRUE  Bool = 1
	FALSE Bool = 0
)

const (
	CODE_OKAY       Status = 0x00000000
	CODE_FAIL       Status = 0x80000000
	CODE_MASK       Status = 0x80000000
	MODULE_ZYCORE   Status = (0x001) << 20
	MODULE_ZYDIS    Status = (0x002) << 20
	MODULE_ARGPARSE Status = (0x003) << 20
	MODULE_USER     Status = (0x3FF) << 20
	MODULE_MASK     Status = (0x7FF) << 20
)

func Ok(status uint32) bool {
	return int32(status) >= 0
}
func Failed(status Status) bool {
	return int32(status) < 0
}

const (
	STATUS_SUCCESS                  Status = 0x00 | CODE_OKAY | MODULE_ZYCORE
	STATUS_FAILED                   Status = 0x01 | CODE_FAIL | MODULE_ZYCORE
	STATUS_TRUE                     Status = 0x02 | CODE_OKAY | MODULE_ZYCORE
	STATUS_FALSE                    Status = 0x03 | CODE_OKAY | MODULE_ZYCORE
	STATUS_INVALID_ARGUMENT         Status = 0x04 | CODE_FAIL | MODULE_ZYCORE
	STATUS_INVALID_OPERATION        Status = 0x05 | CODE_FAIL | MODULE_ZYCORE
	STATUS_ACCESS_DENIED            Status = 0x06 | CODE_FAIL | MODULE_ZYCORE
	STATUS_NOT_FOUND                Status = 0x07 | CODE_FAIL | MODULE_ZYCORE
	STATUS_OUT_OF_RANGE             Status = 0x08 | CODE_FAIL | MODULE_ZYCORE
	STATUS_INSUFFICIENT_BUFFER_SIZE Status = 0x09 | CODE_FAIL | MODULE_ZYCORE
	STATUS_NOT_ENOUGH_MEMORY        Status = 0x0A | CODE_FAIL | MODULE_ZYCORE
	STATUS_BAD_SYSTEMCALL           Status = 0x0B | CODE_FAIL | MODULE_ZYCORE
	STATUS_OUT_OF_RESOURCES         Status = 0x0C | CODE_FAIL | MODULE_ZYCORE
	STATUS_MISSING_DEPENDENCY       Status = 0x0D | CODE_FAIL | MODULE_ZYCORE
	STATUS_ARG_NOT_UNDERSTOOD       Status = 0x00 | CODE_FAIL | MODULE_ARGPARSE
	STATUS_TOO_FEW_ARGS             Status = 0x01 | CODE_FAIL | MODULE_ARGPARSE
	STATUS_TOO_MANY_ARGS            Status = 0x02 | CODE_FAIL | MODULE_ARGPARSE
	STATUS_ARG_MISSES_VALUE         Status = 0x03 | CODE_FAIL | MODULE_ARGPARSE
	STATUS_REQUIRED_ARG_MISSING     Status = 0x04 | CODE_FAIL | MODULE_ARGPARSE
	STATUS_NO_MORE_DATA             Status = 0x00 | CODE_FAIL | MODULE_ZYDIS
	STATUS_DECODING_ERROR           Status = 0x01 | CODE_FAIL | MODULE_ZYDIS
	STATUS_INSTRUCTION_TOO_LONG     Status = 0x02 | CODE_FAIL | MODULE_ZYDIS
	STATUS_BAD_REGISTER             Status = 0x03 | CODE_FAIL | MODULE_ZYDIS
	STATUS_ILLEGAL_LOCK             Status = 0x04 | CODE_FAIL | MODULE_ZYDIS
	STATUS_ILLEGAL_LEGACY_PFX       Status = 0x05 | CODE_FAIL | MODULE_ZYDIS
	STATUS_ILLEGAL_REX              Status = 0x06 | CODE_FAIL | MODULE_ZYDIS
	STATUS_INVALID_MAP              Status = 0x07 | CODE_FAIL | MODULE_ZYDIS
	STATUS_MALFORMED_EVEX           Status = 0x08 | CODE_FAIL | MODULE_ZYDIS
	STATUS_MALFORMED_MVEX           Status = 0x09 | CODE_FAIL | MODULE_ZYDIS
	STATUS_INVALID_MASK             Status = 0x0A | CODE_FAIL | MODULE_ZYDIS
	STATUS_SKIP_TOKEN               Status = 0x0B | CODE_OKAY | MODULE_ZYDIS
	STATUS_IMPOSSIBLE_INSTRUCTION   Status = 0x0C | CODE_FAIL | MODULE_ZYDIS
)
