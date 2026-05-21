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

const (
	ATTRIB_HAS_MODRM                InstructionAttributes = 1 << 0
	ATTRIB_HAS_SIB                  InstructionAttributes = 1 << 1
	ATTRIB_HAS_REX                  InstructionAttributes = 1 << 2
	ATTRIB_HAS_XOP                  InstructionAttributes = 1 << 3
	ATTRIB_HAS_VEX                  InstructionAttributes = 1 << 4
	ATTRIB_HAS_EVEX                 InstructionAttributes = 1 << 5
	ATTRIB_HAS_MVEX                 InstructionAttributes = 1 << 6
	ATTRIB_IS_RELATIVE              InstructionAttributes = 1 << 7
	ATTRIB_IS_PRIVILEGED            InstructionAttributes = 1 << 8
	ATTRIB_CPUFLAG_ACCESS           InstructionAttributes = 1 << 9
	ATTRIB_CPU_STATE_CR             InstructionAttributes = 1 << 10
	ATTRIB_CPU_STATE_CW             InstructionAttributes = 1 << 11
	ATTRIB_FPU_STATE_CR             InstructionAttributes = 1 << 12
	ATTRIB_FPU_STATE_CW             InstructionAttributes = 1 << 13
	ATTRIB_XMM_STATE_CR             InstructionAttributes = 1 << 14
	ATTRIB_XMM_STATE_CW             InstructionAttributes = 1 << 15
	ATTRIB_ACCEPTS_LOCK             InstructionAttributes = 1 << 16
	ATTRIB_ACCEPTS_REP              InstructionAttributes = 1 << 17
	ATTRIB_ACCEPTS_REPE             InstructionAttributes = 1 << 18
	ATTRIB_ACCEPTS_REPZ             InstructionAttributes = ATTRIB_ACCEPTS_REPE
	ATTRIB_ACCEPTS_REPNE            InstructionAttributes = 1 << 19
	ATTRIB_ACCEPTS_REPNZ            InstructionAttributes = ATTRIB_ACCEPTS_REPNE
	ATTRIB_ACCEPTS_BND              InstructionAttributes = 1 << 20
	ATTRIB_ACCEPTS_XACQUIRE         InstructionAttributes = 1 << 21
	ATTRIB_ACCEPTS_XRELEASE         InstructionAttributes = 1 << 22
	ATTRIB_ACCEPTS_HLE_WITHOUT_LOCK InstructionAttributes = 1 << 23
	ATTRIB_ACCEPTS_BRANCH_HINTS     InstructionAttributes = 1 << 24
	ATTRIB_ACCEPTS_NOTRACK          InstructionAttributes = 1 << 25
	ATTRIB_ACCEPTS_SEGMENT          InstructionAttributes = 1 << 26
	ATTRIB_HAS_LOCK                 InstructionAttributes = 1 << 27
	ATTRIB_HAS_REP                  InstructionAttributes = 1 << 28
	ATTRIB_HAS_REPE                 InstructionAttributes = 1 << 29
	ATTRIB_HAS_REPZ                 InstructionAttributes = ATTRIB_HAS_REPE
	ATTRIB_HAS_REPNE                InstructionAttributes = 1 << 30
	ATTRIB_HAS_REPNZ                InstructionAttributes = ATTRIB_HAS_REPNE
	ATTRIB_HAS_BND                  InstructionAttributes = 1 << 31
	ATTRIB_HAS_XACQUIRE             InstructionAttributes = 1 << 32
	ATTRIB_HAS_XRELEASE             InstructionAttributes = 1 << 33
	ATTRIB_HAS_BRANCH_NOT_TAKEN     InstructionAttributes = 1 << 34
	ATTRIB_HAS_BRANCH_TAKEN         InstructionAttributes = 1 << 35
	ATTRIB_HAS_NOTRACK              InstructionAttributes = 1 << 36
	ATTRIB_HAS_SEGMENT_CS           InstructionAttributes = 1 << 37
	ATTRIB_HAS_SEGMENT_SS           InstructionAttributes = 1 << 38
	ATTRIB_HAS_SEGMENT_DS           InstructionAttributes = 1 << 39
	ATTRIB_HAS_SEGMENT_ES           InstructionAttributes = 1 << 40
	ATTRIB_HAS_SEGMENT_FS           InstructionAttributes = 1 << 41
	ATTRIB_HAS_SEGMENT_GS           InstructionAttributes = 1 << 42
	ATTRIB_HAS_SEGMENT              InstructionAttributes = ATTRIB_HAS_SEGMENT_CS |
		ATTRIB_HAS_SEGMENT_SS |
		ATTRIB_HAS_SEGMENT_DS |
		ATTRIB_HAS_SEGMENT_ES |
		ATTRIB_HAS_SEGMENT_FS |
		ATTRIB_HAS_SEGMENT_GS
	ATTRIB_HAS_OPERANDSIZE InstructionAttributes = 1 << 43
	ATTRIB_HAS_ADDRESSSIZE InstructionAttributes = 1 << 44
	ATTRIB_HAS_EVEX_B      InstructionAttributes = 1 << 45
)

const (
	MAX_INSTRUCTION_LENGTH    int = 15
	MAX_OPERAND_COUNT         int = 10
	MAX_OPERAND_COUNT_VISIBLE int = 5
)
