package archs

import (
	"fmt"
	"strings"

	"github.com/suifei/asm2hex/bindings/capstone"
	"github.com/suifei/asm2hex/bindings/keystone"
)

func Disassemble(arch capstone.Architecture, mode capstone.Mode, code []byte, offset uint64, bigEndian bool /*syntaxValue capstone.OptionValue, */, addAddress bool) (string, uint64, bool, error) {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

	engine, err := capstone.New(arch, mode)
	if err != nil {
		return "", 0, false, err
	}
	defer engine.Close()

	// if syntaxValue != 0 {
	// 	engine.Option(capstone.OPT_SYNTAX, capstone.OptionValue(syntaxValue))
	// }

	if bigEndian {
		code, err = capstoneToBigEndian(code, arch, mode)
		if err != nil {
			return "", 0, false, err
		}
	}

	insns, err := engine.Disasm(code, offset, 0)
	if err != nil {
		return "", 0, false, err
	}

	var result string
	for _, insn := range insns {
		if addAddress {
			result += fmt.Sprintf("%08X:\t%-6s\t%-20s\n", insn.Address(), insn.Mnemonic(), insn.OpStr())
		} else {
			result += fmt.Sprintf("%-6s\t%-20s\n", insn.Mnemonic(), insn.OpStr())
		}
	}

	return result, uint64(len(insns)), true, nil
}
func Assemble(arch keystone.Architecture, mode keystone.Mode, code string, offset uint64, bigEndian bool /*, syntaxValue keystone.OptionValue*/) ([]byte, uint64, bool, error) {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

	code = strings.TrimSpace(code)
	if code == "" {
		return nil, 0, false, fmt.Errorf("empty code")
	}
	if strings.HasPrefix(code, ";") {
		return nil, 0, false, fmt.Errorf("commented code")
	}
	if idx := strings.Index(code, ";"); idx > 0 {
		code = code[:idx]
	}

	ks, err := keystone.New(keystone.Architecture(arch), keystone.Mode(mode))
	if err != nil {
		return nil, 0, false, err
	}
	defer ks.Close()

	// if syntaxValue != 0 {
	// 	ks.Option(keystone.OPT_SYNTAX, keystone.OptionValue(syntaxValue))
	// }

	encoding, stat_count, ok := ks.Assemble(code, offset)
	if err := ks.LastError(); err != nil {
		return nil, 0, false, err
	}

	if ok && bigEndian {
		encoding, err = keystoneToBigEndian(encoding, arch, mode)
		if err != nil {
			return nil, 0, false, err
		}
	}

	return encoding, stat_count, ok, nil
}
