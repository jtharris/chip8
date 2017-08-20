package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"chip8/system"
)

func TestShiftLeftParser_Matches(t *testing.T) {
	parser := ShiftLeftParser{}

	assert.True(t, parser.Matches(0x8D3E))
}

func TestShiftLeftParser_DoesNotMatchFirst(t *testing.T) {
	parser := ShiftLeftParser{}

	assert.False(t, parser.Matches(0x2D3E))
}

func TestShiftLeftParser_DoesNotMatchLast(t *testing.T) {
	parser := ShiftLeftParser{}

	assert.False(t, parser.Matches(0x8D3D))
}

func TestShiftLeftParser_CreateOp(t *testing.T) {
	parser := ShiftLeftParser{}
	expected := ShiftLeftOp{register: 0xD}

	assert.Equal(t, expected, parser.CreateOp(0x8D0E))
}

func TestShiftLeftOp_String(t *testing.T) {
	op := ShiftLeftOp{register: 0xD}

	assert.Equal(t, "VD << 1", op.String())
}

func TestShiftLeft_Execute1Overflow(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Registers[0xD] = 0xAF	// 1010 1111

	op := ShiftLeftOp{register: 0xD}

	// When
	op.Execute(&vm)

	// Then
	expected := byte(0x5E)		// 0101 1110
	assert.Equal(t, expected, vm.Registers[0xD])

	// Also check the overflow register
	assert.Equal(t, byte(0x1), vm.Registers[0xF])
}

func TestShiftLeft_Execute0Overflow(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Registers[0xD] = 0x2E	// 0010 1110

	op := ShiftLeftOp{register: 0xD}

	// When
	op.Execute(&vm)

	// Then
	expected := byte(0x5C)		// 0101 1100
	assert.Equal(t, expected, vm.Registers[0xD])

	// Also check the overflow register
	assert.Equal(t, byte(0x0), vm.Registers[0xF])
}
