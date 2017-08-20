package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"chip8/system"
)

func TestShiftRightParser_Matches(t *testing.T) {
	parser := ShiftRightParser{}

	assert.True(t, parser.Matches(0x8D36))
}

func TestShiftRightParser_DoesNotMatchFirst(t *testing.T) {
	parser := ShiftRightParser{}

	assert.False(t, parser.Matches(0x2D36))
}

func TestShiftRightParser_DoesNotMatchLast(t *testing.T) {
	parser := ShiftRightParser{}

	assert.False(t, parser.Matches(0x8D37))
}

func TestShiftRightParser_CreateOp(t *testing.T) {
	parser := ShiftRightParser{}
	expected := ShiftRightOp{register: 0xD}

	assert.Equal(t, expected, parser.CreateOp(0x8D06))
}

func TestShiftRightOp_String(t *testing.T) {
	op := ShiftRightOp{register: 0xD}

	assert.Equal(t, "VD >> 1", op.String())
}

func TestShiftRight_Execute1Overflow(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Registers[0xD] = 0x2F	// 0010 1111

	op := ShiftRightOp{register: 0xD}

	// When
	op.Execute(&vm)

	// Then
	expected := byte(0x17)		// 0001 0111
	assert.Equal(t, expected, vm.Registers[0xD])

	// Also check the overflow register
	assert.Equal(t, byte(0x1), vm.Registers[0xF])
}

func TestShiftRight_Execute0Overflow(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Registers[0xD] = 0x2E	// 0010 1110

	op := ShiftRightOp{register: 0xD}

	// When
	op.Execute(&vm)

	// Then
	expected := byte(0x17)		// 0001 0111
	assert.Equal(t, expected, vm.Registers[0xD])

	// Also check the overflow register
	assert.Equal(t, byte(0x0), vm.Registers[0xF])
}
