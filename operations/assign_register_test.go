package operations

import (
	"chip8/system"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAssignRegisterParser_Matches(t *testing.T) {
	parser := assignRegisterParser{}

	assert.True(t, parser.matches(0x80A0))
}

func TestAssignRegisterParser_DoesNotMatchFirst(t *testing.T) {
	parser := assignRegisterParser{}

	assert.False(t, parser.matches(0x90A0))
}

func TestAssignRegisterParser_DoesNotMatchLast(t *testing.T) {
	parser := assignRegisterParser{}

	assert.False(t, parser.matches(0x80AF))
}

func TestAssignRegisterParser_CreateOp(t *testing.T) {
	parser := assignRegisterParser{}

	op := AssignRegisterOp{
		register1: 0xD,
		register2: 0x4,
	}

	assert.Equal(t, op, parser.createOp(0x8D40))
}

func TestAssignRegisterOp_String(t *testing.T) {
	op := AssignRegisterOp{
		register1: 0x3,
		register2: 0x5,
	}

	assert.Equal(t, "V3 = V5", op.String())
}

func TestAssignRegisterOp_Execute(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Registers[0x3] = 0x13
	vm.Registers[0x5] = 0x04

	op := AssignRegisterOp{
		register1: 0x3,
		register2: 0x5,
	}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, byte(0x04), vm.Registers[0x3])

	// Also make sure source register is unaltered
	assert.Equal(t, byte(0x04), vm.Registers[0x5])
}
