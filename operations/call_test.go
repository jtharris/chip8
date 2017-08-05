package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"chip8/system"
)

func TestCallParser_Matches(t *testing.T) {
	parser := CallParser{}

	assert.True(t, parser.Matches(0x2FF8))
}

func TestCallParser_DoesNotMatch(t *testing.T) {
	parser := CallParser{}

	assert.False(t, parser.Matches(0x1FF8))
}

func TestCallParser_CreateOp(t *testing.T) {
	parser := CallParser{}
	expected := CallOp{address: 0x07F9}

	assert.Equal(t, expected, parser.CreateOp(0x27F9))
}

func TestCallOp_String(t *testing.T) {
	op := CallOp{address: 0x027A}

	assert.Equal(t, "Call subroutine at:  27A", op.String())
}

func TestCallOp_Execute(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.ProgramCounter = 0x0280

	op := CallOp{address: 0x0800}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, uint16(0x07FE), vm.ProgramCounter)
	assert.Equal(t, []uint16{0x0280}, vm.Stack)
}
