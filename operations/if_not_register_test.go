package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"chip8/system"
)

func TestIfNotRegistersNotEqualParser_Matches(t *testing.T) {
	parser := IfNotRegisterParser{}

	assert.True(t, parser.Matches(0x97B0))
}

func TestIfNotRegisterParser_DoesNotMatchFirst(t *testing.T) {
	parser := IfNotRegisterParser{}

	assert.False(t, parser.Matches(0x57B0))
}

func TestIfNotRegisterParser_DoesNotMatchLast(t *testing.T) {
	parser := IfNotRegisterParser{}

	assert.False(t, parser.Matches(0x97B5))
}

func TestIfNotRegisterParser_CreateOp(t *testing.T) {
	parser := IfNotRegisterParser{}
	expected := IfNotRegisterOp{register1: 0x7, register2: 0xB}

	assert.Equal(t, expected, parser.CreateOp(0x97B0))
}

func TestIfNotRegisterOp_String(t *testing.T) {
	op := IfNotRegisterOp{register1: 0x5, register2: 0xA}

	assert.Equal(t, "If V5 != VA", op.String())
}

func TestIfNotRegisterOp_ExecuteMatch(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.ProgramCounter = 0x8
	vm.Registers[0x4] = 0x0F
	vm.Registers[0xA] = 0x0F

	op := IfNotRegisterOp{register1: 0x4, register2: 0xA}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, uint16(0x8), vm.ProgramCounter)
}

func TestIfNotRegisterOp_ExecuteNoMatch(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.ProgramCounter = 0x8
	vm.Registers[0x4] = 0xAC
	vm.Registers[0xC] = 0xAD

	op := IfNotRegisterOp{register1: 0x4, register2: 0xC}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, uint16(0xA), vm.ProgramCounter)
}
