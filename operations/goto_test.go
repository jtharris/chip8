package operations

import (
	"chip8/system"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGotoParser_Matches(t *testing.T) {
	parser := gotoParser{}

	assert.True(t, parser.matches(0x1847))
}

func TestGotoParser_DoesNotMatch(t *testing.T) {
	parser := gotoParser{}

	assert.False(t, parser.matches(0x2847))
}

func TestGotoParser_CreateOp(t *testing.T) {
	parser := gotoParser{}
	expected := GotoOp{address: 0x08F4}

	assert.Equal(t, expected, parser.createOp(0x18F4))
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
	assert.Equal(t, uint16(0x02C7), vm.ProgramCounter)
}
