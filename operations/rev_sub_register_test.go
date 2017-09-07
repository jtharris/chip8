package operations

import (
	"github.com/jtharris/chip8/system"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseSubtractRegisterParser_Matches(t *testing.T) {
	parser := reverseSubtractRegisterParser{}

	assert.True(t, parser.matches(0x82B7))
}

func TestReverseSubtractRegisterParser_DoesNotMatchFirst(t *testing.T) {
	parser := reverseSubtractRegisterParser{}

	assert.False(t, parser.matches(0x92B7))
}

func TestReverseSubtractRegisterParser_DoesNotMatchLast(t *testing.T) {
	parser := reverseSubtractRegisterParser{}

	assert.False(t, parser.matches(0x82B8))
}

func TestReverseSubtractRegisterParser_CreateOp(t *testing.T) {
	parser := reverseSubtractRegisterParser{}
	expected := ReverseSubtractRegisterOp{register1: 0x2, register2: 0xB}

	assert.Equal(t, expected, parser.createOp(0x82B7))
}

func TestReverseSubtractRegisterOp_String(t *testing.T) {
	op := ReverseSubtractRegisterOp{register1: 0x2, register2: 0xB}

	assert.Equal(t, "V2 = VB - V2", op.String())
}

func TestReverseSubtractRegisterOp_ExecuteNoBorrow(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Registers[0x2] = 0x03
	vm.Registers[0xB] = 0x0A

	op := ReverseSubtractRegisterOp{register1: 0x2, register2: 0xB}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, byte(0x07), vm.Registers[0x2])

	// Second register should be unchanged
	assert.Equal(t, byte(0x0A), vm.Registers[0xB])

	// No Borrow, so set VF to 1
	assert.Equal(t, byte(0x01), vm.Registers[0xF])
}

func TestReverseSubtractRegisterOp_ExecuteBorrow(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Registers[0x3] = 0x13
	vm.Registers[0x1] = 0x04

	op := ReverseSubtractRegisterOp{register1: 0x3, register2: 0x1}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, byte(0xF1), vm.Registers[0x3])

	// Second register should be unchanged
	assert.Equal(t, byte(0x04), vm.Registers[0x1])

	// Borrow, so set VF to 0
	assert.Equal(t, byte(0x0), vm.Registers[0xF])
}
