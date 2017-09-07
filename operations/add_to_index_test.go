package operations

import (
	"github.com/jtharris/chip8/system"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddToIndexParser_Matches(t *testing.T) {
	parser := addToIndexParser{}
	assert.True(t, parser.matches(0xFB1E))
}

func TestAddToIndexParser_DoesNotMatchFirst(t *testing.T) {
	parser := addToIndexParser{}
	assert.False(t, parser.matches(0xCB1E))
}

func TestAddToIndexParser_DoesNotMatchLast(t *testing.T) {
	parser := addToIndexParser{}
	assert.False(t, parser.matches(0xFB1D))
}

func TestAddToIndexParser_CreateOp(t *testing.T) {
	parser := addToIndexParser{}
	expected := AddToIndexOp{register: 0xC}

	assert.Equal(t, expected, parser.createOp(0xFC1E))
}

func TestAddToIndexOp_String(t *testing.T) {
	op := AddToIndexOp{register: 0x7}

	assert.Equal(t, "I += V7", op.String())
}

func TestAddToIndexOp_Execute(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	op := AddToIndexOp{register: 0x4}

	vm.IndexRegister = 0x2A8
	vm.Registers[0x4] = 0xA

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, uint16(0x2B2), vm.IndexRegister)
}
