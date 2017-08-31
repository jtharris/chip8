package operations

import (
	"chip8/system"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitwiseXorParser_Matches(t *testing.T) {
	parser := bitwiseXorParser{}

	assert.True(t, parser.matches(0x8033))
}

func TestBitwiseXorParser_DoesNotMatchFirst(t *testing.T) {
	parser := bitwiseXorParser{}

	assert.False(t, parser.matches(0xA033))
}

func TestBitwiseXorParser_DoesNotMatchLast(t *testing.T) {
	parser := bitwiseXorParser{}

	assert.False(t, parser.matches(0x8031))
}

func TestBitwiseXorParser_CreateOp(t *testing.T) {
	parser := bitwiseXorParser{}
	expected := BitwiseXorOp{
		register1: 0xA,
		register2: 0x4,
	}

	assert.Equal(t, expected, parser.createOp(0x8A42))
}

func TestBitwiseXorOp_String(t *testing.T) {
	op := BitwiseXorOp{
		register1: 0xA,
		register2: 0x4,
	}

	assert.Equal(t, "VA = VA ^ V4", op.String())
}

func TestBitwiseXorOp_Execute(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Registers[0x3] = 0xB3 // 1011 0011
	vm.Registers[0x6] = 0x76 // 0111 0110
	expected := byte(0xC5)   // 1100 0101

	op := BitwiseXorOp{
		register1: 0x3,
		register2: 0x6,
	}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, expected, vm.Registers[0x3])
}
