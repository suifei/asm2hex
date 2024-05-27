package archs

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/suifei/asm2hex/bindings/capstone"
	"github.com/suifei/asm2hex/bindings/keystone"
)

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
