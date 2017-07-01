package operations

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"chip8/system"
)

func TestBitwiseAndParser_Matches(t *testing.T) {
	parser := BitwiseAndParser{}

	assert.True(t, parser.Matches(0x8032))
}

func TestBitwiseAndParser_DoesNotMatchFirst(t *testing.T) {
	parser := BitwiseAndParser{}

	assert.False(t, parser.Matches(0xA032))
}

func TestBitwiseAndParser_DoesNotMatchLast(t *testing.T) {
	parser := BitwiseAndParser{}

	assert.False(t, parser.Matches(0x8031))
}

func TestBitwiseAndParser_CreateOp(t *testing.T) {
	parser := BitwiseAndParser{}
	expected := BitwiseAndOp{
		register1: 0xA,
		register2: 0x4,
	}

	assert.Equal(t, expected, parser.CreateOp(0x8A42))
}

func TestBitwiseAndOp_String(t *testing.T) {
	op := BitwiseAndOp{
		register1: 0xA,
		register2: 0x4,
	}

	assert.Equal(t, "VA = VA & V4", op.String())
}

func TestBitwiseAndOp_Execute(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Registers[0x3] = 0xBF   // 1011 1111
	vm.Registers[0x6] = 0x76   // 0111 0110
	expected := byte(0x36)     // 0011 0110

	op := BitwiseAndOp{
		register1: 0x3,
		register2: 0x6,
	}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, expected, vm.Registers[0x3])
}
