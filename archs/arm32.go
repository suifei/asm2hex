package archs

import (
	"fmt"

	"github.com/suifei/asm2hex/bindings/keystone"
	"github.com/suifei/asm2hex/bindings/capstone"
)

func Arm32Disasm(encoding []byte, offset uint64, bigEndian bool) (code string, stat_count uint64, ok bool, err error) {
	//panic
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
			ok = false
			return
		}
	}()

	engine, err := capstone.New(capstone.ARCH_ARM, capstone.MODE_ARM)
	if err == nil {
		engine.Option(capstone.OPT_TYPE_SYNTAX, capstone.OPT_SYNTAX_ATT)
		defer engine.Close()

		if bigEndian {
			encoding = bigEndian32Bytes(encoding)
		}
		insns, err := engine.Dis(encoding, offset, 0)

		if err == nil {
			for _, insn := range insns {
				code += fmt.Sprintf("%08X:\t%-6s\t%-20s\n", insn.Addr(), insn.Mnemonic(), insn.OpStr())
			}
			stat_count = uint64(len(insns))
			ok = true
			return code, stat_count, ok, err
		}
		return code, stat_count, ok, err
	}
	return code, stat_count, ok, err
}

func Arm32(code string, offset uint64, bigEndian bool) (encoding []byte, stat_count uint64, ok bool, err error) {
	//panic
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
			ok = false
			return
		}
	}()

	var ks *keystone.Keystone
	ks, _ = keystone.New(keystone.ARCH_ARM, keystone.MODE_ARM)
	// ks.Option(keystone.OPT_SYNTAX, keystone.OPT_SYNTAX_GAS)
	encoding, stat_count, ok = ks.Assemble(code, offset)
	err = ks.LastError()
	ks.Close()

	if ok {
		if bigEndian {
			encoding = bigEndian32Bytes(encoding)
		}
	}

	return encoding, stat_count, ok, err
}
