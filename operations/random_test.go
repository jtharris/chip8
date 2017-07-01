package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"chip8/system"
)

func TestRandomParser_Matches(t *testing.T) {
	parser := RandomParser{}

	assert.True(t, parser.Matches(0xC5FE))
}

func TestRandomParser_DoesNotMatch(t *testing.T) {
	parser := RandomParser{}

	assert.False(t, parser.Matches(0xD5FE))
}

func TestRandomParser_CreateOp(t *testing.T) {
	parser := RandomParser{}
	expected := RandomOp{
		register: 0x5,
		value: 0xFE,
	}

	assert.Equal(t, expected, parser.CreateOp(0xC5FE))
}

func TestRandomOp_String(t *testing.T) {
	op := RandomOp{
		register: 0x5,
		value: 0xFE,
	}

	assert.Equal(t, "V5 = rand(255) & FE", op.String())
}

func TestRandomOp_Execute(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}

	op := RandomOp{
		register: 0x5,
		value: 0xFE,
	}

	// When
	op.Execute(&vm)

	// Then
	// Note:  Golang has a deterministic sequence of psuedo-random numbers, so this test assumes that the
	//        first number of the sequence is used
	assert.Equal(t, byte(0x52), vm.Registers[0x5])
}