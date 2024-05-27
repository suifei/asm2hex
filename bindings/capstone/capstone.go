package capstone

import (
	"bytes"
	"unsafe"
)

// #cgo LDFLAGS: -lcapstone
// #include <capstone/capstone.h>
import "C"

const (
	API_MAJOR = 5
	API_MINOR = 0
)

type Architecture uint
type Mode uint
type OptionType uint
type OptionValue uint
type OperandType uint
type AccessType uint
type GroupType uint

const (
	ARCH_ARM        Architecture = C.CS_ARCH_ARM
	ARCH_ARM64      Architecture = C.CS_ARCH_ARM64
	ARCH_MIPS       Architecture = C.CS_ARCH_MIPS
	ARCH_X86        Architecture = C.CS_ARCH_X86
	ARCH_PPC        Architecture = C.CS_ARCH_PPC
	ARCH_SPARC      Architecture = C.CS_ARCH_SPARC
	ARCH_SYSZ       Architecture = C.CS_ARCH_SYSZ
	ARCH_XCORE      Architecture = C.CS_ARCH_XCORE
	ARCH_M68K       Architecture = C.CS_ARCH_M68K
	ARCH_TMS320C64X Architecture = C.CS_ARCH_TMS320C64X
	ARCH_M680X      Architecture = C.CS_ARCH_M680X
	ARCH_EVM        Architecture = C.CS_ARCH_EVM
	ARCH_MOS65XX    Architecture = C.CS_ARCH_MOS65XX
	ARCH_WASM       Architecture = C.CS_ARCH_WASM
	ARCH_BPF        Architecture = C.CS_ARCH_BPF
	ARCH_RISCV      Architecture = C.CS_ARCH_RISCV
	ARCH_SH         Architecture = C.CS_ARCH_SH
	ARCH_TRICORE    Architecture = C.CS_ARCH_TRICORE
	ARCH_MAX        Architecture = C.CS_ARCH_MAX
	ARCH_ALL        Architecture = C.CS_ARCH_ALL
)
const (
	MODE_LITTLE_ENDIAN         Mode = C.CS_MODE_LITTLE_ENDIAN
	MODE_ARM                   Mode = C.CS_MODE_ARM
	MODE_16                    Mode = C.CS_MODE_16
	MODE_32                    Mode = C.CS_MODE_32
	MODE_64                    Mode = C.CS_MODE_64
	MODE_THUMB                 Mode = C.CS_MODE_THUMB
	MODE_MCLASS                Mode = C.CS_MODE_MCLASS
	MODE_V8                    Mode = C.CS_MODE_V8
	MODE_MICRO                 Mode = C.CS_MODE_MICRO
	MODE_MIPS3                 Mode = C.CS_MODE_MIPS3
	MODE_MIPS32R6              Mode = C.CS_MODE_MIPS32R6
	MODE_MIPS2                 Mode = C.CS_MODE_MIPS2
	MODE_V9                    Mode = C.CS_MODE_V9
	MODE_QPX                   Mode = C.CS_MODE_QPX
	MODE_SPE                   Mode = C.CS_MODE_SPE
	MODE_BOOKE                 Mode = C.CS_MODE_BOOKE
	MODE_PS                    Mode = C.CS_MODE_PS
	MODE_M68K_000              Mode = C.CS_MODE_M68K_000
	MODE_M68K_010              Mode = C.CS_MODE_M68K_010
	MODE_M68K_020              Mode = C.CS_MODE_M68K_020
	MODE_M68K_030              Mode = C.CS_MODE_M68K_030
	MODE_M68K_040              Mode = C.CS_MODE_M68K_040
	MODE_M68K_060              Mode = C.CS_MODE_M68K_060
	MODE_BIG_ENDIAN            Mode = C.CS_MODE_BIG_ENDIAN
	MODE_MIPS32                Mode = C.CS_MODE_MIPS32
	MODE_MIPS64                Mode = C.CS_MODE_MIPS64
	MODE_M680X_6301            Mode = C.CS_MODE_M680X_6301
	MODE_M680X_6309            Mode = C.CS_MODE_M680X_6309
	MODE_M680X_6800            Mode = C.CS_MODE_M680X_6800
	MODE_M680X_6801            Mode = C.CS_MODE_M680X_6801
	MODE_M680X_6805            Mode = C.CS_MODE_M680X_6805
	MODE_M680X_6808            Mode = C.CS_MODE_M680X_6808
	MODE_M680X_6809            Mode = C.CS_MODE_M680X_6809
	MODE_M680X_6811            Mode = C.CS_MODE_M680X_6811
	MODE_M680X_CPU12           Mode = C.CS_MODE_M680X_CPU12
	MODE_M680X_HCS08           Mode = C.CS_MODE_M680X_HCS08
	MODE_BPF_CLASSIC           Mode = C.CS_MODE_BPF_CLASSIC
	MODE_BPF_EXTENDED          Mode = C.CS_MODE_BPF_EXTENDED
	MODE_RISCV32               Mode = C.CS_MODE_RISCV32
	MODE_RISCV64               Mode = C.CS_MODE_RISCV64
	MODE_RISCVC                Mode = C.CS_MODE_RISCVC
	MODE_MOS65XX_6502          Mode = C.CS_MODE_MOS65XX_6502
	MODE_MOS65XX_65C02         Mode = C.CS_MODE_MOS65XX_65C02
	MODE_MOS65XX_W65C02        Mode = C.CS_MODE_MOS65XX_W65C02
	MODE_MOS65XX_65816         Mode = C.CS_MODE_MOS65XX_65816
	MODE_MOS65XX_65816_LONG_M  Mode = C.CS_MODE_MOS65XX_65816_LONG_M
	MODE_MOS65XX_65816_LONG_X  Mode = C.CS_MODE_MOS65XX_65816_LONG_X
	MODE_MOS65XX_65816_LONG_MX Mode = C.CS_MODE_MOS65XX_65816_LONG_MX
	MODE_SH2                   Mode = C.CS_MODE_SH2
	MODE_SH2A                  Mode = C.CS_MODE_SH2A
	MODE_SH3                   Mode = C.CS_MODE_SH3
	MODE_SH4                   Mode = C.CS_MODE_SH4
	MODE_SH4A                  Mode = C.CS_MODE_SH4A
	MODE_SHFPU                 Mode = C.CS_MODE_SHFPU
	MODE_SHDSP                 Mode = C.CS_MODE_SHDSP
	MODE_TRICORE_110           Mode = C.CS_MODE_TRICORE_110
	MODE_TRICORE_120           Mode = C.CS_MODE_TRICORE_120
	MODE_TRICORE_130           Mode = C.CS_MODE_TRICORE_130
	MODE_TRICORE_131           Mode = C.CS_MODE_TRICORE_131
	MODE_TRICORE_160           Mode = C.CS_MODE_TRICORE_160
	MODE_TRICORE_161           Mode = C.CS_MODE_TRICORE_161
	MODE_TRICORE_162           Mode = C.CS_MODE_TRICORE_162
)
const (
	OPT_INVALID          OptionType = C.CS_OPT_INVALID          ///< No option specified
	OPT_SYNTAX           OptionType = C.CS_OPT_SYNTAX           ///< Assembly output syntax
	OPT_DETAIL           OptionType = C.CS_OPT_DETAIL           ///< Break down instruction structure into details
	OPT_MODE             OptionType = C.CS_OPT_MODE             ///< Change engine's mode at run-time
	OPT_MEM              OptionType = C.CS_OPT_MEM              ///< User-defined dynamic memory related functions
	OPT_SKIPDATA         OptionType = C.CS_OPT_SKIPDATA         ///< Skip data when disassembling. Then engine is in SKIPDATA mode.
	OPT_SKIPDATA_SETUP   OptionType = C.CS_OPT_SKIPDATA_SETUP   ///< Setup user-defined function for SKIPDATA option
	OPT_MNEMONIC         OptionType = C.CS_OPT_MNEMONIC         ///< Customize instruction mnemonic
	OPT_UNSIGNED         OptionType = C.CS_OPT_UNSIGNED         ///< print immediate operands in unsigned form
	OPT_NO_BRANCH_OFFSET OptionType = C.CS_OPT_NO_BRANCH_OFFSET ///< ARM, prints branch immediates without offset.
)

const (
	OPT_OFF              OptionValue = C.CS_OPT_OFF              ///< Turn OFF an option - default for CS_OPT_DETAIL, CS_OPT_SKIPDATA, CS_OPT_UNSIGNED.
	OPT_ON               OptionValue = C.CS_OPT_ON               ///< Turn ON an option (CS_OPT_DETAIL, CS_OPT_SKIPDATA).
	OPT_SYNTAX_DEFAULT   OptionValue = C.CS_OPT_SYNTAX_DEFAULT   ///< Default asm syntax (CS_OPT_SYNTAX).
	OPT_SYNTAX_INTEL     OptionValue = C.CS_OPT_SYNTAX_INTEL     ///< X86 Intel asm syntax - default on X86 (CS_OPT_SYNTAX).
	OPT_SYNTAX_ATT       OptionValue = C.CS_OPT_SYNTAX_ATT       ///< X86 ATT asm syntax (CS_OPT_SYNTAX).
	OPT_SYNTAX_NOREGNAME OptionValue = C.CS_OPT_SYNTAX_NOREGNAME ///< Prints register name with only number (CS_OPT_SYNTAX)
	OPT_SYNTAX_MASM      OptionValue = C.CS_OPT_SYNTAX_MASM      ///< X86 Intel Masm syntax (CS_OPT_SYNTAX).
	OPT_SYNTAX_MOTOROLA  OptionValue = C.CS_OPT_SYNTAX_MOTOROLA  ///< MOS65XX use $ as hex prefix
)

const (
	OP_INVALID OperandType = C.CS_OP_INVALID ///< Uninitialized/invalid.
	OP_REG     OperandType = C.CS_OP_REG     ///< Register operand.
	OP_IMM     OperandType = C.CS_OP_IMM     ///< Immediate operand.
	OP_FP      OperandType = C.CS_OP_FP      ///< Floating-Point operand.
	OP_MEM     OperandType = C.CS_OP_MEM     ///< Memory operand.
)

const (
	AC_INVALID AccessType = C.CS_AC_INVALID ///< Uninitialized/invalid.
	AC_READ    AccessType = C.CS_AC_READ    ///< Operand read.
	AC_WRITE   AccessType = C.CS_AC_WRITE   ///< Operand write.
)

const (
	GRP_INVALID         GroupType = C.CS_GRP_INVALID         ///< Uninitialized/invalid group type.
	GRP_JUMP            GroupType = C.CS_GRP_JUMP            ///< Group of all jump instructions (conditional+direct).
	GRP_CALL            GroupType = C.CS_GRP_CALL            ///< Group of all call instructions.
	GRP_RET             GroupType = C.CS_GRP_RET             ///< Group of all return instructions.
	GRP_INT             GroupType = C.CS_GRP_INT             ///< Group of all interrupt instructions.
	GRP_IRET            GroupType = C.CS_GRP_IRET            ///< Group of all interrupt return instructions.
	GRP_PRIVILEGE       GroupType = C.CS_GRP_PRIVILEGE       ///< Group of all privileged instructions.
	GRP_BRANCH_RELATIVE GroupType = C.CS_GRP_BRANCH_RELATIVE ///< Group of all branch instructions with relative offsets.
)

type Engine struct {
	handle C.csh
}

type CsError C.cs_err

func (e CsError) Error() string {
	return C.GoString(C.cs_strerror(C.cs_err(e)))
}

func New(arch Architecture, mode Mode) (*Engine, error) {
	var handle C.csh
	cserr := C.cs_open(C.cs_arch(arch), C.cs_mode(mode), &handle)
	if cserr != C.CS_ERR_OK {
		return nil, CsError(cserr)
	}
	C.cs_option(handle, C.CS_OPT_DETAIL, C.CS_OPT_OFF)
	return &Engine{handle}, nil
}

func (e *Engine) Option(opt_type OptionType, value C.size_t) error {
	if cserr := C.cs_option(e.handle, C.cs_opt_type(opt_type), value); cserr != C.CS_ERR_OK {
		return CsError(cserr)
	}
	return nil
}

func (e *Engine) Disasm(code []byte, addr, count uint64) ([]Ins, error) {
	if len(code) == 0 {
		return nil, nil
	}
	ptr := (*C.uint8_t)(unsafe.Pointer(&code[0]))

	var disptr *C.cs_insn
	num := C.cs_disasm(e.handle, ptr, C.size_t(len(code)), C.uint64_t(addr), C.size_t(count), &disptr)
	if num > 0 {
		dis := (*[1 << 23]C.cs_insn)(unsafe.Pointer(disptr))[:num]
		ret := make([]Ins, num)
		for i, ins := range dis {
			outins := &ret[i]
			// byte array casts of cs_ins fields
			mnemonic := (*[32]byte)(unsafe.Pointer(&ins.mnemonic[0]))
			byteData := (*[16]byte)(unsafe.Pointer(&ins.bytes[0]))
			opstr := (*[160]byte)(unsafe.Pointer(&ins.op_str[0]))

			// populate the return ins fields
			outins.address = uint64(ins.address)
			// this is faster than C.GoBytes()
			outins.dataSlice = outins.data[:ins.size]
			copy(outins.dataSlice, byteData[:])

			// populate the strings
			if off := bytes.IndexByte(mnemonic[:], 0); off > 0 {
				outins.mnemonic = string(mnemonic[:off])
			}
			if off := bytes.IndexByte(opstr[:], 0); off > 0 {
				outins.opstr = string(opstr[:off])
			}
		}
		C.free(unsafe.Pointer(disptr))
		return ret, nil
	}
	return nil, CsError(C.cs_errno(e.handle))
}

func (e *Engine) Close() error {
	return CsError(C.cs_close(&e.handle))
}

// conforms to usercorn/models.Ins interface
type Ins struct {
	address   uint64
	dataSlice []byte
	mnemonic  string
	opstr     string
	data      [16]byte
}

func (i Ins) Address() uint64 {
	return i.address
}

func (i Ins) Bytes() []byte {
	return i.dataSlice
}

func (i Ins) Mnemonic() string {
	return i.mnemonic
}

func (i Ins) OpStr() string {
	return i.opstr
}
