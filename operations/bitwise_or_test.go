package operations

import (
	"chip8/system"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitwiseOrParser_Matches(t *testing.T) {
	parser := bitwiseOrParser{}

	assert.True(t, parser.matches(0x8031))
}

func TestBitwiseOrParser_DoesNotMatchFirst(t *testing.T) {
	parser := bitwiseOrParser{}

	assert.False(t, parser.matches(0xA031))
}

func TestBitwiseOrParser_DoesNotMatchLast(t *testing.T) {
	parser := bitwiseOrParser{}

	assert.False(t, parser.matches(0x8032))
}

func TestBitwiseOrParser_CreateOp(t *testing.T) {
	parser := bitwiseOrParser{}
	expected := BitwiseOrOp{
		register1: 0xA,
		register2: 0x4,
	}

	assert.Equal(t, expected, parser.createOp(0x8A42))
}

func TestBitwiseOrOp_String(t *testing.T) {
	op := BitwiseOrOp{
		register1: 0xA,
		register2: 0x4,
	}

	assert.Equal(t, "VA = VA | V4", op.String())
}

func TestBitwiseOrOp_Execute(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Registers[0x3] = 0xB1 // 1011 0001
	vm.Registers[0x6] = 0x76 // 0111 0110
	expected := byte(0xF7)   // 1111 0111

	op := BitwiseOrOp{
		register1: 0x3,
		register2: 0x6,
	}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, expected, vm.Registers[0x3])
}
