package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"chip8/system"
)

func TestSubtractRegisterParser_Matches(t *testing.T) {
	parser := SubtractRegisterParser{}

	assert.True(t, parser.Matches(0x86C5))
}

func TestSubtractRegisterParser_DoesNotMatchFirst(t *testing.T) {
	parser := SubtractRegisterParser{}

	assert.False(t, parser.Matches(0xA6C5))
}

func TestSubtractRegisterParser_DoesNotMatchLast(t *testing.T) {
	parser := SubtractRegisterParser{}

	assert.False(t, parser.Matches(0x86C6))
}

func TestSubtractRegisterParser_CreateOp(t *testing.T) {
	parser := SubtractRegisterParser{}
	expected := SubtractRegisterOp{register1: 0x6, register2: 0xC}

	assert.Equal(t, expected, parser.CreateOp(0x86C5))
}

func TestSubtractRegisterOp_String(t *testing.T) {
	op := SubtractRegisterOp{register1: 0x4, register2: 0xA}

	assert.Equal(t, "V4 -= VA", op.String())
}

func TestSubtractRegisterOp_ExecuteNoBorrow(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Registers[0x4] = 0x28
	vm.Registers[0xA] = 0x03
	op := SubtractRegisterOp{register1: 0x4, register2: 0xA}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, byte(0x25), vm.Registers[0x4])
	assert.Equal(t, byte(0x03), vm.Registers[0xA])

	// No borrow, so VF should be 1
	assert.Equal(t, byte(0x1), vm.Registers[0xF])
}

func TestSubtractRegisterOp_ExecuteBorrow(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Registers[0x3] = 0x28
	vm.Registers[0x2] = 0x33
	op := SubtractRegisterOp{register1: 0x3, register2: 0x2}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, byte(0xf5), vm.Registers[0x3])
	assert.Equal(t, byte(0x33), vm.Registers[0x2])

	// There was a  borrow, so VF should be 0
	assert.Equal(t, byte(0x0), vm.Registers[0xF])
}
