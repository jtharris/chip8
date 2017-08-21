package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"chip8/system"
)


func TestAddConstantParser_Matches(t *testing.T) {
	parser := AddConstantParser{}

	assert.True(t, parser.Matches(0x7454))
}

func TestAddConstantParser_DoesNotMatch(t *testing.T) {
	parser := AddConstantParser{}

	assert.False(t, parser.Matches(0x6009))
}

func TestAddConstantParser_CreateOp(t *testing.T) {
	parser := AddConstantParser{}
	op :=parser.CreateOp(0x7A81)

	expected := AddConstantOp{
		register: 0xA,
		value: 0x81,
	}

	assert.Equal(t, expected, op)
}

func TestAddConstantOp_String(t *testing.T) {
	op := AddConstantOp{
		register: 0x3,
		value: 0xFF,
	}

	assert.Equal(t, "V3 += FF", op.String())
}

func TestAddConstantOp_Execute(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Registers[0xA] = 0x34

	op := AddConstantOp{
		register: 0xA,
		value: 0x06,
	}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, byte(0x3A), vm.Registers[0xA])
}

func TestAddConstantOp_ExecuteOverflow(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Registers[0x1] = 0xFE

	op := AddConstantOp{
		register: 0x1,
		value: 0x04,
	}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, byte(0x02), vm.Registers[0x1])

	// This operations doesn't set VF for overflow
	assert.Equal(t, byte(0x00), vm.Registers[0xF])
}

func TestAddConstantOp_ExecuteUnsetOverflow(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Registers[0xF] = 0x1
	vm.Registers[0x1] = 0x0E

	op := AddConstantOp{
		register: 0x1,
		value: 0x4,
	}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, byte(0x12), vm.Registers[0x1])

	// This operations doesn't change VF for overflow
	assert.Equal(t, byte(0x01), vm.Registers[0xF])
}

func TestAddConstantOp_ExecuteOverflowCountdown(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Registers[0x6] = 0x35

	op := AddConstantOp{
		register: 0x6,
		value: 0xFF,
	}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, byte(0x34), vm.Registers[0x6])

	// This operations doesn't set VF for overflow
	assert.Equal(t, byte(0x00), vm.Registers[0xF])
}
