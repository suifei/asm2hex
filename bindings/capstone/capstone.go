package capstone

import (
	"bytes"
	"unsafe"
)

// #cgo darwin CFLAGS: -I/usr/local/include -O2 -g -Wall -Werror
// #cgo darwin LDFLAGS: -L/usr/local/lib -lcapstone
// #cgo windows CFLAGS: -ID:/works/asm2hex/link/win64/include
// #cgo windows LDFLAGS: -LD:/works/asm2hex/link/win64/lib -lcapstone
// #include <capstone/capstone.h>
import "C"

const (
	ARCH_ARM   = C.CS_ARCH_ARM
	ARCH_ARM64 = C.CS_ARCH_ARM64
	ARCH_MIPS  = C.CS_ARCH_MIPS
	ARCH_X86   = C.CS_ARCH_X86
	ARCH_PPC   = C.CS_ARCH_PPC
	ARCH_SPARC = C.CS_ARCH_SPARC
	ARCH_SYSZ  = C.CS_ARCH_SYSZ
	ARCH_XCORE = C.CS_ARCH_XCORE
)

const (
	MODE_LITTLE_ENDIAN = C.CS_MODE_LITTLE_ENDIAN
	MODE_ARM           = C.CS_MODE_ARM
	MODE_16            = C.CS_MODE_16
	MODE_32            = C.CS_MODE_32
	MODE_64            = C.CS_MODE_64
	MODE_THUMB         = C.CS_MODE_THUMB
	MODE_MCLASS        = C.CS_MODE_MCLASS
	MODE_V8            = C.CS_MODE_V8
	MODE_MICRO         = C.CS_MODE_MICRO
	MODE_MIPS3         = C.CS_MODE_MIPS3
	MODE_MIPS32R6      = C.CS_MODE_MIPS32R6
	MODE_V9            = C.CS_MODE_V9
	MODE_BIG_ENDIAN    = C.CS_MODE_BIG_ENDIAN
	MODE_MIPS32        = C.CS_MODE_MIPS32
	MODE_MIPS64        = C.CS_MODE_MIPS64
)

const (
	OPT_TYPE_INVALID        = C.CS_OPT_INVALID
	OPT_TYPE_SYNTAX         = C.CS_OPT_SYNTAX
	OPT_TYPE_DETAIL         = C.CS_OPT_DETAIL
	OPT_TYPE_MODE           = C.CS_OPT_MODE
	OPT_TYPE_MEM            = C.CS_OPT_MEM
	OPT_TYPE_SKIPDATA       = C.CS_OPT_SKIPDATA
	OPT_TYPE_SKIPDATA_SETUP = C.CS_OPT_SKIPDATA_SETUP
)

const (
	OPT_OFF              = C.CS_OPT_OFF
	OPT_ON               = C.CS_OPT_ON
	OPT_SYNTAX_DEFAULT   = C.CS_OPT_SYNTAX_DEFAULT
	OPT_SYNTAX_INTEL     = C.CS_OPT_SYNTAX_INTEL
	OPT_SYNTAX_ATT       = C.CS_OPT_SYNTAX_ATT
	OPT_SYNTAX_NOREGNAME = C.CS_OPT_SYNTAX_NOREGNAME
)

type Engine struct {
	handle C.csh
}

type CsError C.cs_err

func (e CsError) Error() string {
	return C.GoString(C.cs_strerror(C.cs_err(e)))
}

func New(arch, mode int) (*Engine, error) {
	var handle C.csh
	cserr := C.cs_open(C.cs_arch(arch), C.cs_mode(mode), &handle)
	if cserr != C.CS_ERR_OK {
		return nil, CsError(cserr)
	}
	C.cs_option(handle, C.CS_OPT_DETAIL, C.CS_OPT_OFF)
	return &Engine{handle}, nil
}

func (e *Engine) Option(opt_type C.cs_opt_type, value C.size_t) error {
	if cserr := C.cs_option(e.handle, opt_type, value); cserr != C.CS_ERR_OK {
		return CsError(cserr)
	}
	return nil
}

func (e *Engine) Dis(code []byte, addr, count uint64) ([]Ins, error) {
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
			outins.addr = uint64(ins.address)
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
	addr      uint64
	dataSlice []byte
	mnemonic  string
	opstr     string
	data      [16]byte
}

func (i Ins) Addr() uint64 {
	return i.addr
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
