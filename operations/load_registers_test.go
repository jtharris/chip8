package operations

import (
	"github.com/jtharris/chip8/system"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadRegistersParser_Matches(t *testing.T) {
	parser := loadRegistersParser{}

	assert.True(t, parser.matches(0xFB65))
}

func TestLoadRegistersParser_DoesNotMatchStart(t *testing.T) {
	parser := loadRegistersParser{}

	assert.False(t, parser.matches(0xEB65))
}

func TestLoadRegistersParser_DoesNotMatchEnd(t *testing.T) {
	parser := loadRegistersParser{}

	assert.False(t, parser.matches(0xFB63))
}

func TestLoadRegistersParser_DoesNotMatchMiddle(t *testing.T) {
	parser := loadRegistersParser{}

	assert.False(t, parser.matches(0xFB75))
}

func TestLoadRegistersParser_CreateOp(t *testing.T) {
	parser := loadRegistersParser{}
	expected := LoadRegistersOp{topRegister: 0xB}

	assert.Equal(t, expected, parser.createOp(0xFB65))
}

func TestLoadRegistersOp_String(t *testing.T) {
	op := LoadRegistersOp{topRegister: 0x7}

	assert.Equal(t, "load_registers(V7, &I)", op.String())
}

func TestLoadRegistersOp_Execute(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Memory[0x10] = 0x0A
	vm.Memory[0x11] = 0x0B
	vm.Memory[0x12] = 0x0C
	vm.Memory[0x13] = 0x0D
	vm.Memory[0x14] = 0x04 // Not used
	vm.IndexRegister = 0x10

	op := LoadRegistersOp{topRegister: 0x3}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, byte(0x0A), vm.Registers[0x0])
	assert.Equal(t, byte(0x0B), vm.Registers[0x1])
	assert.Equal(t, byte(0x0C), vm.Registers[0x2])
	assert.Equal(t, byte(0x0D), vm.Registers[0x3])

	// Also be sure that the next register was not loaded
	assert.Equal(t, byte(0x00), vm.Registers[0x4])
}
