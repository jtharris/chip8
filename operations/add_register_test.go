package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"chip8/system"
)

func TestAddRegiserParser_Matches(t *testing.T) {
	parser := AddRegisterParser{}

	assert.True(t, parser.Matches(0x8454))
}

func TestAddRegiserParser_DoesNotMatchFirst(t *testing.T) {
	parser := AddRegisterParser{}

	assert.False(t, parser.Matches(0x7454))
}

func TestAddRegiserParser_DoesNotMatchLast(t *testing.T) {
	parser := AddRegisterParser{}

	assert.False(t, parser.Matches(0x845B))
}

func TestAddRegisterParser_CreateOp(t *testing.T) {
	parser := AddRegisterParser{}
	op := AddRegisterOp{
		register1: 0x4,
		register2: 0x5,
	}

	assert.Equal(t, op, parser.CreateOp(0x8454))
}

func TestAddRegisterOp_String(t *testing.T) {
	op := AddRegisterOp{
		register1: 0x4,
		register2: 0xC,
	}

	assert.Equal(t, "V4 += VC", op.String())
}

func TestAddRegisterOp_Execute(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Registers[0x2] = 0x34
	vm.Registers[0x3] = 0x80

	op := AddRegisterOp{
		register1: 0x2,
		register2: 0x3,
	}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, byte(0xB4), vm.Registers[0x2])

	// Also verify that V3 remains the same
	assert.Equal(t, byte(0x80), vm.Registers[0x3])

	// And that the carry register is not set
	assert.Equal(t, byte(0x0), vm.Registers[0xF])
}

func TestAddRegisterOp_ExecuteOverflow(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Registers[0x1] = 0xFF
	vm.Registers[0xA] = 0x01

	op := AddRegisterOp{
		register1: 0x1,
		register2: 0xA,
	}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, byte(0x00), vm.Registers[0x1])

	// Also verify that VA remains the same
	assert.Equal(t, byte(0x01), vm.Registers[0xA])

	// And that the carry register is set
	assert.Equal(t, byte(0x01), vm.Registers[0xF])
}

func TestAddRegisterOp_ExecuteUnsetOverflow(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Registers[0x1] = 0x04
	vm.Registers[0x2] = 0x04
	vm.Registers[0xF] = 0x01

	op := AddRegisterOp{
		register1: 0x1,
		register2: 0x2,
	}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, byte(0x08), vm.Registers[0x1])

	// Also verify that V2 remains the same
	assert.Equal(t, byte(0x04), vm.Registers[0x2])

	// And that the carry register is not set
	assert.Equal(t, byte(0x00), vm.Registers[0xF])
}
