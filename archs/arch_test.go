package archs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/suifei/asm2hex/bindings/capstone"
	"github.com/suifei/asm2hex/bindings/keystone"
)

func TestDisassemble(t *testing.T) {
	code := []byte{0x48, 0x83, 0x3d, 0x01, 0x02, 0x03, 0x04, 0x05}
	arch := capstone.ARCH_X86
	mode := capstone.MODE_64
	syntaxValue := capstone.OPT_SYNTAX_INTEL
	str, count, ok, err := Disassemble(arch, mode, code, 0x100, false, syntaxValue)
	require.NoError(t, err)
	require.True(t, ok)
	require.Equal(t, 1, int(count))
	require.Equal(t, "cmpq\t$0x4030201020304,0x105\t;0x100\n", str)
}

func TestAssemble_x86_64_code_big_endian(t *testing.T) {
	code := "mov rax, 1"
	encoding, count, ok, err := Assemble(keystone.ARCH_X86, keystone.MODE_64, code, 0x100, true, keystone.OPT_SYNTAX_ATT)
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), count)
	assert.True(t, ok)
	assert.Equal(t, []byte{0x48, 0xc7, 0xc0, 0x01, 0x00, 0x00, 0x00, 0x00}, encoding)
}

func TestAssemble_x86_64_code_little_endian(t *testing.T) {
	code := "mov rax, 1"
	encoding, count, ok, err := Assemble(keystone.ARCH_X86, keystone.MODE_64, code, 0x100, false, keystone.OPT_SYNTAX_ATT)
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), count)
	assert.True(t, ok)
	assert.Equal(t, []byte{0xc7, 0xc0, 0x01, 0x00, 0x00, 0x00, 0x00, 0x48}, encoding)
}

func TestAssemble_x86_32_code_big_endian(t *testing.T) {
	code := "mov eax, 1"
	encoding, count, ok, err := Assemble(keystone.ARCH_X86, keystone.MODE_32, code, 0x100, true, keystone.OPT_SYNTAX_ATT)
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), count)
	assert.True(t, ok)
	assert.Equal(t, []byte{0x66, 0x3d, 0x01, 0x00, 0x00, 0x00}, encoding)
}

func TestAssemble_x86_32_code_little_endian(t *testing.T) {
	code := "mov eax, 1"
	encoding, count, ok, err := Assemble(keystone.ARCH_X86, keystone.MODE_32, code, 0x100, false, keystone.OPT_SYNTAX_ATT)
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), count)
	assert.True(t, ok)
	assert.Equal(t, []byte{0x01, 0x3d, 0x66, 0x00, 0x00, 0x00}, encoding)
}

func TestAssemble_ARM_code_big_endian(t *testing.T) {
	code := "mov r1, 1"
	encoding, count, ok, err := Assemble(keystone.ARCH_ARM, keystone.MODE_ARM, code, 0x100, true, keystone.OPT_SYNTAX_ATT)
	assert.Nil(t, err)
	assert.Equal(t, uint64(4), count)
	assert.True(t, ok)
	assert.Equal(t, []byte{0x0, 0x0, 0xa0, 0xe3}, encoding)
}

func TestAssemble_ARM_code_little_endian(t *testing.T) {
	code := "mov r1, 1"
	encoding, count, ok, err := Assemble(keystone.ARCH_ARM, keystone.MODE_ARM, code, 0x100, false, keystone.OPT_SYNTAX_ATT)
	assert.Nil(t, err)
	assert.Equal(t, uint64(4), count)
	assert.True(t, ok)
	assert.Equal(t, []byte{0xe3, 0xa0, 0x0, 0x0}, encoding)
}

func TestAssemble_ARM64_code_big_endian(t *testing.T) {
	code := "mov w1, 1"
	encoding, count, ok, err := Assemble(keystone.ARCH_ARM64, keystone.MODE_ARM, code, 0x100, true, keystone.OPT_SYNTAX_ATT)
	assert.Nil(t, err)
	assert.Equal(t, uint64(4), count)
	assert.True(t, ok)
	assert.Equal(t, []byte{0x0, 0x0, 0x20, 0x1f}, encoding)
}

func TestAssemble_ARM64_code_little_endian(t *testing.T) {
	code := "mov w1, 1"
	encoding, count, ok, err := Assemble(keystone.ARCH_ARM64, keystone.MODE_ARM, code, 0x100, false, keystone.OPT_SYNTAX_ATT)
	assert.Nil(t, err)
	assert.Equal(t, uint64(4), count)
	assert.True(t, ok)
	assert.Equal(t, []byte{0x1f, 0x20, 0x0, 0x0}, encoding)
}
