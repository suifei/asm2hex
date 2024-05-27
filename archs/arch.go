package archs

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/suifei/asm2hex/bindings/capstone"
	"github.com/suifei/asm2hex/bindings/keystone"
)

var WithRiscv = true

type Option struct {
	Const uint64
	Name  string
}

type OptionSlice []Option

var KeystoneArchOptions = OptionSlice{
	{uint64(keystone.ARCH_ARM), "ARM"},
	{uint64(keystone.ARCH_ARM64), "ARM64"},
	{uint64(keystone.ARCH_MIPS), "MIPS"},
	{uint64(keystone.ARCH_X86), "X86"},
	{uint64(keystone.ARCH_PPC), "PPC"},
	{uint64(keystone.ARCH_RISCV), "RISCV"},
	{uint64(keystone.ARCH_SPARC), "SPARC"},
	{uint64(keystone.ARCH_SYSTEMZ), "SYSTEMZ"},
	{uint64(keystone.ARCH_HEXAGON), "HEXAGON"},
}

var KeystoneModeList = OptionSlice{}
var KeystoneModeOptions = map[uint64]OptionSlice{
	uint64(keystone.ARCH_ARM):     {{uint64(keystone.MODE_ARM), "ARM"}, {uint64(keystone.MODE_THUMB), "THUMB"}, {uint64(keystone.MODE_V8), "V8"}},
	uint64(keystone.ARCH_ARM64):   {{uint64(keystone.MODE_LITTLE_ENDIAN), "LITTLE_ENDIAN"}},
	uint64(keystone.ARCH_MIPS):    {{uint64(keystone.MODE_MICRO), "MICRO"}, {uint64(keystone.MODE_MIPS3), "MIPS3"}, {uint64(keystone.MODE_MIPS32R6), "MIPS32R6"}, {uint64(keystone.MODE_MIPS32), "MIPS32"}, {uint64(keystone.MODE_MIPS64), "MIPS64"}},
	uint64(keystone.ARCH_X86):     {{uint64(keystone.MODE_16), "16"}, {uint64(keystone.MODE_32), "32"}, {uint64(keystone.MODE_64), "64"}},
	uint64(keystone.ARCH_RISCV):   {{uint64(keystone.MODE_RISCV32), "RISCV32"}, {uint64(keystone.MODE_RISCV64), "RISCV64"}},
	uint64(keystone.ARCH_PPC):     {{uint64(keystone.MODE_PPC32), "PPC32"}, {uint64(keystone.MODE_PPC64), "PPC64"}, {uint64(keystone.MODE_QPX), "QPX"}},
	uint64(keystone.ARCH_SPARC):   {{uint64(keystone.MODE_SPARC32), "SPARC32"}, {uint64(keystone.MODE_SPARC64), "SPARC64"}, {uint64(keystone.MODE_V9), "V9"}},
	uint64(keystone.ARCH_SYSTEMZ): {{uint64(keystone.MODE_BIG_ENDIAN), "BIG_ENDIAN"}},
	uint64(keystone.ARCH_HEXAGON): {{uint64(keystone.MODE_BIG_ENDIAN), "BIG_ENDIAN"}},
}

// var KeystoneSyntaxList = OptionSlice{}
// var KeystoneSyntaxOptions = map[uint64]OptionSlice{
// 	uint64(keystone.ARCH_ARM):     {{uint64(keystone.OPT_SYNTAX_INTEL), "Intel"}, {uint64(keystone.OPT_SYNTAX_ATT), "ATT"}},
// 	uint64(keystone.ARCH_ARM64):   {{uint64(keystone.OPT_SYNTAX_INTEL), "Intel"}, {uint64(keystone.OPT_SYNTAX_ATT), "ATT"}},
// 	uint64(keystone.ARCH_MIPS):    {{uint64(keystone.OPT_SYNTAX_INTEL), "Intel"}, {uint64(keystone.OPT_SYNTAX_ATT), "ATT"}},
// 	uint64(keystone.ARCH_X86):     {{uint64(keystone.OPT_SYNTAX_INTEL), "Intel"}, {uint64(keystone.OPT_SYNTAX_ATT), "ATT"}, {uint64(keystone.OPT_SYNTAX_NASM), "NASM"}, {uint64(keystone.OPT_SYNTAX_MASM), "MASM"}, {uint64(keystone.OPT_SYNTAX_GAS), "GAS"}, {uint64(keystone.OPT_SYNTAX_RADIX16), "Radix16"}},
// 	uint64(keystone.ARCH_PPC):     {{uint64(keystone.OPT_SYNTAX_INTEL), "Intel"}, {uint64(keystone.OPT_SYNTAX_ATT), "ATT"}},
// 	uint64(keystone.ARCH_SPARC):   {{uint64(keystone.OPT_SYNTAX_INTEL), "Intel"}, {uint64(keystone.OPT_SYNTAX_ATT), "ATT"}},
// 	uint64(keystone.ARCH_SYSTEMZ): {{uint64(keystone.OPT_SYNTAX_INTEL), "Intel"}, {uint64(keystone.OPT_SYNTAX_ATT), "ATT"}},
// 	uint64(keystone.ARCH_HEXAGON): {{uint64(keystone.OPT_SYNTAX_INTEL), "Intel"}},
// }

var CapstoneArchOptions = OptionSlice{
	{uint64(capstone.ARCH_ARM), "ARM"},
	{uint64(capstone.ARCH_ARM64), "ARM64"},
	{uint64(capstone.ARCH_MIPS), "MIPS"},
	{uint64(capstone.ARCH_X86), "X86"},
	{uint64(capstone.ARCH_PPC), "PPC"},
	{uint64(capstone.ARCH_SPARC), "SPARC"},
	{uint64(capstone.ARCH_SYSZ), "SYSZ"},
	{uint64(capstone.ARCH_XCORE), "XCORE"},
	{uint64(capstone.ARCH_M68K), "M68K"},
	{uint64(capstone.ARCH_TMS320C64X), "TMS320C64X"},
	{uint64(capstone.ARCH_M680X), "M680X"},
	{uint64(capstone.ARCH_EVM), "EVM"},
	{uint64(capstone.ARCH_MOS65XX), "MOS65XX"},
	{uint64(capstone.ARCH_WASM), "WASM"},
	{uint64(capstone.ARCH_BPF), "BPF"},
	{uint64(capstone.ARCH_RISCV), "RISCV"},
	{uint64(capstone.ARCH_SH), "SH"},
	{uint64(capstone.ARCH_TRICORE), "TRICORE"},
}

var CapstoneModeList = OptionSlice{}
var CapstoneModeOptions = map[uint64]OptionSlice{
	uint64(capstone.ARCH_ARM):     {{uint64(capstone.MODE_ARM), "ARM"}, {uint64(capstone.MODE_THUMB), "THUMB"}, {uint64(capstone.MODE_MCLASS), "MCLASS"}, {uint64(capstone.MODE_V8), "V8"}},
	uint64(capstone.ARCH_ARM64):   {{uint64(capstone.MODE_LITTLE_ENDIAN), "LITTLE_ENDIAN"}},
	uint64(capstone.ARCH_MIPS):    {{uint64(capstone.MODE_MIPS32), "MIPS32"}, {uint64(capstone.MODE_MIPS64), "MIPS64"}, {uint64(capstone.MODE_MICRO), "MICRO"}, {uint64(capstone.MODE_MIPS3), "MIPS3"}, {uint64(capstone.MODE_MIPS32R6), "MIPS32R6"}, {uint64(capstone.MODE_MIPS2), "MIPS2"}},
	uint64(capstone.ARCH_X86):     {{uint64(capstone.MODE_16), "16"}, {uint64(capstone.MODE_32), "32"}, {uint64(capstone.MODE_64), "64"}},
	uint64(capstone.ARCH_PPC):     {{uint64(capstone.MODE_LITTLE_ENDIAN), "LITTLE_ENDIAN"}, {uint64(capstone.MODE_QPX), "QPX"}, {uint64(capstone.MODE_SPE), "SPE"}, {uint64(capstone.MODE_BOOKE), "BOOKE"}},
	uint64(capstone.ARCH_SPARC):   {{uint64(capstone.MODE_V9), "V9"}},
	uint64(capstone.ARCH_SYSZ):    {{uint64(capstone.MODE_BIG_ENDIAN), "BIG_ENDIAN"}},
	uint64(capstone.ARCH_XCORE):   {{uint64(capstone.MODE_LITTLE_ENDIAN), "LITTLE_ENDIAN"}},
	uint64(capstone.ARCH_M68K):    {{uint64(capstone.MODE_M68K_000), "M68K_000"}, {uint64(capstone.MODE_M68K_010), "M68K_010"}, {uint64(capstone.MODE_M68K_020), "M68K_020"}, {uint64(capstone.MODE_M68K_030), "M68K_030"}, {uint64(capstone.MODE_M68K_040), "M68K_040"}, {uint64(capstone.MODE_M68K_060), "M68K_060"}},
	uint64(capstone.ARCH_M680X):   {{uint64(capstone.MODE_M680X_6301), "M680X_6301"}, {uint64(capstone.MODE_M680X_6309), "M680X_6309"}, {uint64(capstone.MODE_M680X_6800), "M680X_6800"}, {uint64(capstone.MODE_M680X_6801), "M680X_6801"}, {uint64(capstone.MODE_M680X_6805), "M680X_6805"}, {uint64(capstone.MODE_M680X_6808), "M680X_6808"}, {uint64(capstone.MODE_M680X_6809), "M680X_6809"}, {uint64(capstone.MODE_M680X_6811), "M680X_6811"}, {uint64(capstone.MODE_M680X_CPU12), "M680X_CPU12"}, {uint64(capstone.MODE_M680X_HCS08), "M680X_HCS08"}},
	uint64(capstone.ARCH_EVM):     {{uint64(capstone.MODE_BIG_ENDIAN), "BIG_ENDIAN"}},
	uint64(capstone.ARCH_MOS65XX): {{uint64(capstone.MODE_MOS65XX_6502), "MOS65XX_6502"}, {uint64(capstone.MODE_MOS65XX_65C02), "MOS65XX_65C02"}, {uint64(capstone.MODE_MOS65XX_W65C02), "MOS65XX_W65C02"}, {uint64(capstone.MODE_MOS65XX_65816), "MOS65XX_65816"}, {uint64(capstone.MODE_MOS65XX_65816_LONG_M), "MOS65XX_65816_LONG_M"}, {uint64(capstone.MODE_MOS65XX_65816_LONG_X), "MOS65XX_65816_LONG_X"}, {uint64(capstone.MODE_MOS65XX_65816_LONG_MX), "MOS65XX_65816_LONG_MX"}},
	uint64(capstone.ARCH_WASM):    {{uint64(capstone.MODE_LITTLE_ENDIAN), "LITTLE_ENDIAN"}},
	uint64(capstone.ARCH_BPF):     {{uint64(capstone.MODE_BPF_CLASSIC), "BPF_CLASSIC"}, {uint64(capstone.MODE_BPF_EXTENDED), "BPF_EXTENDED"}},
	uint64(capstone.ARCH_RISCV):   {{uint64(capstone.MODE_RISCV32), "RISCV32"}, {uint64(capstone.MODE_RISCV64), "RISCV64"}, {uint64(capstone.MODE_RISCVC), "RISCVC"}},
	uint64(capstone.ARCH_SH):      {{uint64(capstone.MODE_SH2), "SH2"}, {uint64(capstone.MODE_SH2A), "SH2A"}, {uint64(capstone.MODE_SH3), "SH3"}, {uint64(capstone.MODE_SH4), "SH4"}, {uint64(capstone.MODE_SH4A), "SH4A"}, {uint64(capstone.MODE_SHFPU), "SHFPU"}, {uint64(capstone.MODE_SHDSP), "SHDSP"}},
	uint64(capstone.ARCH_TRICORE): {{uint64(capstone.MODE_TRICORE_110), "TRICORE_110"}, {uint64(capstone.MODE_TRICORE_120), "TRICORE_120"}, {uint64(capstone.MODE_TRICORE_130), "TRICORE_130"}, {uint64(capstone.MODE_TRICORE_131), "TRICORE_131"}, {uint64(capstone.MODE_TRICORE_160), "TRICORE_160"}, {uint64(capstone.MODE_TRICORE_161), "TRICORE_161"}, {uint64(capstone.MODE_TRICORE_162), "TRICORE_162"}},
}

// var CapstoneSyntaxList = OptionSlice{}
// var CapstoneSyntaxOptions = map[uint64]OptionSlice{
// 	uint64(capstone.ARCH_ARM):        {{uint64(capstone.OPT_SYNTAX_DEFAULT), "Default"}, {uint64(capstone.OPT_SYNTAX_NOREGNAME), "NOREGNAME"}},
// 	uint64(capstone.ARCH_ARM64):      {{uint64(capstone.OPT_SYNTAX_DEFAULT), "Default"}, {uint64(capstone.OPT_SYNTAX_NOREGNAME), "NOREGNAME"}},
// 	uint64(capstone.ARCH_MIPS):       {{uint64(capstone.OPT_SYNTAX_DEFAULT), "Default"}, {uint64(capstone.OPT_SYNTAX_NOREGNAME), "NOREGNAME"}},
// 	uint64(capstone.ARCH_X86):        {{uint64(capstone.OPT_SYNTAX_INTEL), "Intel"}, {uint64(capstone.OPT_SYNTAX_ATT), "ATT"}, {uint64(capstone.OPT_SYNTAX_MASM), "MASM"}},
// 	uint64(capstone.ARCH_PPC):        {{uint64(capstone.OPT_SYNTAX_DEFAULT), "Default"}, {uint64(capstone.OPT_SYNTAX_NOREGNAME), "NOREGNAME"}},
// 	uint64(capstone.ARCH_SPARC):      {{uint64(capstone.OPT_SYNTAX_DEFAULT), "Default"}, {uint64(capstone.OPT_SYNTAX_NOREGNAME), "NOREGNAME"}},
// 	uint64(capstone.ARCH_SYSZ):       {{uint64(capstone.OPT_SYNTAX_DEFAULT), "Default"}, {uint64(capstone.OPT_SYNTAX_NOREGNAME), "NOREGNAME"}},
// 	uint64(capstone.ARCH_XCORE):      {{uint64(capstone.OPT_SYNTAX_DEFAULT), "Default"}, {uint64(capstone.OPT_SYNTAX_NOREGNAME), "NOREGNAME"}},
// 	uint64(capstone.ARCH_M68K):       {{uint64(capstone.OPT_SYNTAX_DEFAULT), "Default"}, {uint64(capstone.OPT_SYNTAX_NOREGNAME), "NOREGNAME"}, {uint64(capstone.OPT_SYNTAX_MOTOROLA), "Motorola"}},
// 	uint64(capstone.ARCH_TMS320C64X): {{uint64(capstone.OPT_SYNTAX_DEFAULT), "Default"}, {uint64(capstone.OPT_SYNTAX_NOREGNAME), "NOREGNAME"}},
// 	uint64(capstone.ARCH_M680X):      {{uint64(capstone.OPT_SYNTAX_DEFAULT), "Default"}, {uint64(capstone.OPT_SYNTAX_NOREGNAME), "NOREGNAME"}, {uint64(capstone.OPT_SYNTAX_MOTOROLA), "Motorola"}},
// 	uint64(capstone.ARCH_EVM):        {{uint64(capstone.OPT_SYNTAX_DEFAULT), "Default"}},
// 	uint64(capstone.ARCH_MOS65XX):    {{uint64(capstone.OPT_SYNTAX_DEFAULT), "Default"}, {uint64(capstone.OPT_SYNTAX_MOTOROLA), "Motorola"}},
// 	uint64(capstone.ARCH_WASM):       {{uint64(capstone.OPT_SYNTAX_DEFAULT), "Default"}},
// 	uint64(capstone.ARCH_BPF):        {{uint64(capstone.OPT_SYNTAX_DEFAULT), "Default"}},
// 	uint64(capstone.ARCH_RISCV):      {{uint64(capstone.OPT_SYNTAX_DEFAULT), "Default"}},
// 	uint64(capstone.ARCH_SH):         {{uint64(capstone.OPT_SYNTAX_DEFAULT), "Default"}},
// 	uint64(capstone.ARCH_TRICORE):    {{uint64(capstone.OPT_SYNTAX_DEFAULT), "Default"}},
// }

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
func capstoneToBigEndian(data []byte, arch capstone.Architecture, mode capstone.Mode) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("input data is empty")
	}

	instructionLength := getCapstoneInstructionLength(arch, mode)
	if instructionLength == 0 {
		return nil, fmt.Errorf("unsupported architecture or mode")
	}

	if len(data)%instructionLength != 0 {
		return nil, fmt.Errorf("input data size is not a multiple of instruction length")
	}

	result := make([]byte, 0, len(data))

	for i := 0; i < len(data); i += instructionLength {
		instruction := data[i : i+instructionLength]
		bigEndianInstruction, err := littleToBigEndian(instruction)
		if err != nil {
			return nil, fmt.Errorf("failed to convert instruction at offset %d: %v", i, err)
		}
		result = append(result, bigEndianInstruction...)
	}

	return result, nil
}

func keystoneToBigEndian(data []byte, arch keystone.Architecture, mode keystone.Mode) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("input data is empty")
	}

	instructionLength := getKeystoneInstructionLength(arch, mode)
	if instructionLength == 0 {
		return nil, fmt.Errorf("unsupported architecture or mode")
	}

	if len(data)%instructionLength != 0 {
		return nil, fmt.Errorf("input data size is not a multiple of instruction length")
	}

	result := make([]byte, 0, len(data))

	for i := 0; i < len(data); i += instructionLength {
		instruction := data[i : i+instructionLength]
		bigEndianInstruction, err := littleToBigEndian(instruction)
		if err != nil {
			return nil, fmt.Errorf("failed to convert instruction at offset %d: %v", i, err)
		}
		result = append(result, bigEndianInstruction...)
	}

	return result, nil
}

func getCapstoneInstructionLength(arch capstone.Architecture, mode capstone.Mode) int {
	switch arch {
	case capstone.ARCH_ARM:
		if mode&capstone.MODE_THUMB != 0 {
			return 2
		}
		return 4
	case capstone.ARCH_ARM64:
		return 4
	case capstone.ARCH_MIPS:
		if mode&capstone.MODE_MICRO != 0 {
			return 2
		}
		return 4
	case capstone.ARCH_X86:
		if mode&capstone.MODE_16 != 0 {
			return 2
		}
		return 1
	case capstone.ARCH_PPC:
		return 4
	case capstone.ARCH_SPARC:
		return 4
	case capstone.ARCH_SYSZ:
		return 2
	case capstone.ARCH_XCORE:
		return 2
	case capstone.ARCH_M68K:
		return 2
	case capstone.ARCH_TMS320C64X:
		return 4
	case capstone.ARCH_M680X:
		return 1
	case capstone.ARCH_EVM:
		return 1
	case capstone.ARCH_MOS65XX:
		return 1
	case capstone.ARCH_WASM:
		return 1
	case capstone.ARCH_BPF:
		return 8
	case capstone.ARCH_RISCV:
		if mode&capstone.MODE_RISCV32 != 0 {
			return 4
		}
		return 2
	case capstone.ARCH_SH:
		if mode&capstone.MODE_SH2 != 0 || mode&capstone.MODE_SH2A != 0 {
			return 2
		}
		return 2
	case capstone.ARCH_TRICORE:
		return 2
	default:
		return 0
	}
}

func getKeystoneInstructionLength(arch keystone.Architecture, mode keystone.Mode) int {
	switch arch {
	case keystone.ARCH_ARM:
		if mode&keystone.MODE_THUMB != 0 {
			return 2
		}
		return 4
	case keystone.ARCH_ARM64:
		return 4
	case keystone.ARCH_MIPS:
		if mode&keystone.MODE_MICRO != 0 {
			return 2
		}
		if mode&keystone.MODE_MIPS32 != 0 || mode&keystone.MODE_MIPS32R6 != 0 {
			return 4
		}
		if mode&keystone.MODE_MIPS64 != 0 {
			return 4
		}
		return 4
	case keystone.ARCH_X86:
		if mode&keystone.MODE_16 != 0 {
			return 2
		}
		return 1
	case keystone.ARCH_PPC:
		if mode&keystone.MODE_PPC32 != 0 {
			return 4
		}
		if mode&keystone.MODE_PPC64 != 0 {
			return 4
		}
		return 4
	case keystone.ARCH_SPARC:
		if mode&keystone.MODE_SPARC32 != 0 {
			return 4
		}
		if mode&keystone.MODE_SPARC64 != 0 {
			return 4
		}
		return 4
	case keystone.ARCH_RISCV:
		if mode&keystone.MODE_RISCV32 != 0 {
			return 2
		}
		if mode&keystone.MODE_RISCV64 != 0 {
			return 4
		}
		return 2
	case keystone.ARCH_SYSTEMZ:
		return 2
	case keystone.ARCH_HEXAGON:
		return 4
	case keystone.ARCH_EVM:
		return 1
	default:
		return 0
	}
}

func littleToBigEndian(data []byte) ([]byte, error) {
	switch len(data) {
	case 1:
		return data, nil
	case 2:
		var num uint16
		err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &num)
		if err != nil {
			return nil, err
		}
		result := make([]byte, 2)
		binary.BigEndian.PutUint16(result, num)
		return result, nil
	case 4:
		var num uint32
		err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &num)
		if err != nil {
			return nil, err
		}
		result := make([]byte, 4)
		binary.BigEndian.PutUint32(result, num)
		return result, nil
	case 8:
		var num uint64
		err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &num)
		if err != nil {
			return nil, err
		}
		result := make([]byte, 8)
		binary.BigEndian.PutUint64(result, num)
		return result, nil
	default:
		return nil, fmt.Errorf("unsupported instruction size: %d", len(data))
	}
}
