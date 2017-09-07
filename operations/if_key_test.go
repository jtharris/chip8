package operations

import (
	"github.com/jtharris/chip8/system"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIfKeyParser_Matches(t *testing.T) {
	parser := ifKeyParser{}

	assert.True(t, parser.matches(0xED9E))
}

func TestIfKeyParser_DoesNotMatchFirst(t *testing.T) {
	parser := ifKeyParser{}

	assert.False(t, parser.matches(0xFD9E))
}

func TestIfKeyParser_DoesNotMatchLast(t *testing.T) {
	parser := ifKeyParser{}

	assert.False(t, parser.matches(0xED7E))
}

func TestIfKeyParser_CreateOp(t *testing.T) {
	parser := ifKeyParser{}
	expected := IfKeyOp{register: 0xA}

	assert.Equal(t, expected, parser.createOp(0xEA9E))
}

func TestIfKeyOp_String(t *testing.T) {
	op := IfKeyOp{register: 0x3}

	assert.Equal(t, "If key == V3", op.String())
}

func TestIfKeyOp_ExecuteNotPressed(t *testing.T) {
	vm := system.VirtualMachine{}
	vm.Keyboard[0x3] = false
	vm.Registers[0x0] = 0x3
	vm.ProgramCounter = 0x07

	op := IfKeyOp{register: 0x0}

	op.Execute(&vm)

	assert.Equal(t, uint16(0x7), vm.ProgramCounter)
}

func TestIfKeyOp_ExecuteTrue(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Keyboard[0x3] = true
	vm.Registers[0x0] = 0x3
	vm.ProgramCounter = 0x07

	op := IfKeyOp{register: 0x0}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, uint16(0x9), vm.ProgramCounter)
	assert.False(t, vm.Keyboard[0x3])
}
