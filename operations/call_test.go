package operations

import (
	"chip8/system"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCallParser_Matches(t *testing.T) {
	parser := callParser{}

	assert.True(t, parser.matches(0x2FF8))
}

func TestCallParser_DoesNotMatch(t *testing.T) {
	parser := callParser{}

	assert.False(t, parser.matches(0x1FF8))
}

func TestCallParser_CreateOp(t *testing.T) {
	parser := callParser{}
	expected := CallOp{address: 0x07F9}

	assert.Equal(t, expected, parser.createOp(0x27F9))
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

func TestCallOp_ExecuteTwice(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.ProgramCounter = 0x0280

	op := CallOp{address: 0x0800}
	op2 := CallOp{address: 0x0900}

	// When
	op.Execute(&vm)
	vm.IncrementPC()
	op2.Execute(&vm)

	// Then
	assert.Equal(t, uint16(0x08FE), vm.ProgramCounter)
	assert.Equal(t, []uint16{0x0280, 0x0800}, vm.Stack)
}
