package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"chip8/system"
)

func TestReturnParser_Matches(t *testing.T) {
	parser := ReturnParser{}

	assert.True(t, parser.Matches(0x00EE))
}

func TestReturnParser_DoesNotMatch(t *testing.T) {
	parser := ReturnParser{}

	assert.False(t, parser.Matches(0x005E))
}

func TestReturnParser_CreateOp(t *testing.T) {
	parser := ReturnParser{}
	expected := ReturnOp{}

	assert.Equal(t, expected, parser.CreateOp(0x00EE))
}

func TestReturnOp_String(t *testing.T) {
	op := ReturnOp{}

	assert.Equal(t, "Return from subroutine", op.String())
}

func TestReturnOp_Execute(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Stack = []uint16{0x88, 0xAE, 0x2F}
	op := ReturnOp{}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, uint16(0x2F), vm.ProgramCounter)
	assert.Equal(t, []uint16{0x88, 0xAE}, vm.Stack)
}