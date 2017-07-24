package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"chip8/system"
)

func TestGotoParser_Matches(t *testing.T) {
	parser := GotoParser{}

	assert.True(t, parser.Matches(0x1847))
}

func TestGotoParser_DoesNotMatch(t *testing.T) {
	parser := GotoParser{}

	assert.False(t, parser.Matches(0x2847))
}

func TestGotoParser_CreateOp(t *testing.T) {
	parser := GotoParser{}
	expected := GotoOp{address: 0x08F4}

	assert.Equal(t, expected, parser.CreateOp(0x18F4))
}

func TestGotoOp_String(t *testing.T) {
	op := GotoOp{address: 0x08F4}

	assert.Equal(t, "Goto: 8F4", op.String())
}

func TestGotoOp_Execute(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	op := GotoOp{address: 0x02C9}

	// When
	op.Execute(&vm)

	// Then
	// Be sure to remove the 0x0200 offset
	assert.Equal(t, uint16(0x00C9), vm.ProgramCounter)
}
