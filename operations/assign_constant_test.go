package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"chip8/system"
)

func TestAssignConstantParser_Matches(t *testing.T) {
	parser := AssignConstantParser{}

	assert.True(t, parser.Matches(0x60AF))
}

func TestAssignConstantParser_DoesNotMatch(t *testing.T) {
	parser := AssignConstantParser{}

	assert.False(t, parser.Matches(0x70AF))
}

func TestAssignConstantParser_CreateOp(t *testing.T) {
	parser := AssignConstantParser{}

	op := AssignConstantOp{
		register: 0xD,
		value: 0x15,
	}

	assert.Equal(t, op, parser.CreateOp(0x7D15))
}

func TestAssignConstantOp_String(t *testing.T) {
	op := AssignConstantOp{
		register: 0x3,
		value: 0xFE,
	}

	assert.Equal(t, "V3 = FE", op.String())
}

func TestAssignConstantOp_Execute(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Registers[0x3] = 0x13

	op := AssignConstantOp{
		register: 0x3,
		value: 0xFE,
	}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, byte(0xFE), vm.Registers[0x3])
}