package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"chip8/system"
)


func TestIfNotConstantParser_Matches(t *testing.T) {
	parser := ifNotConstantParser{}

	assert.True(t, parser.matches(0x47B6))
}

func TestIfNotConstantParser_DoesNotMatch(t *testing.T) {
	parser := ifNotConstantParser{}

	assert.False(t, parser.matches(0x37B6))
}

func TestIfNotConstantParser_CreateOp(t *testing.T) {
	parser := ifNotConstantParser{}
	expected := IfNotConstantOp{register: 0x7, value: 0xB6}

	assert.Equal(t, expected, parser.createOp(0x47B6))
}

func TestIfNotConstantOp_String(t *testing.T) {
	op := IfNotConstantOp{register: 0x5, value: 0xAC}

	assert.Equal(t, "If V5 != AC", op.String())
}

func TestIfNotConstantOp_ExecuteMatch(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.ProgramCounter = 0x9
	vm.Registers[0x4] = 0xA1

	op := IfNotConstantOp{register: 0x4, value: 0xA1}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, uint16(0x9), vm.ProgramCounter)
}

func TestIfNotConstantOp_ExecuteNoMatch(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.ProgramCounter = 0x9
	vm.Registers[0x4] = 0xAC

	op := IfNotConstantOp{register: 0x4, value: 0xA1}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, uint16(0xB), vm.ProgramCounter)
}

