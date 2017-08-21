package operations

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"chip8/system"
)

func TestBinaryCodedDecimalParser_Matches(t *testing.T) {
	parser := binaryCodedDecimalParser{}

	assert.True(t, parser.matches(0xf033))
}

func TestBinaryCodedDecimalParser_DoesNotMatchFirst(t *testing.T) {
	parser := binaryCodedDecimalParser{}

	assert.False(t, parser.matches(0x9033))
}

func TestBinaryCodedDecimalParser_DoesNotMatchLast(t *testing.T) {
	parser := binaryCodedDecimalParser{}

	assert.False(t, parser.matches(0xf034))
}

func TestBinaryCodedDecimalParser_CreateOp(t *testing.T) {
	parser := binaryCodedDecimalParser{}

	op := BinaryCodedDecimalOp{register: 0x5}

	assert.Equal(t, op, parser.createOp(0xf533))
}

func TestBinaryCodedDecimalOp_String(t *testing.T) {
	op := BinaryCodedDecimalOp{register: 0xA}

	assert.Equal(t, "BCD(VA)", op.String())
}

func TestBinaryCodedDecimalOp_Execute(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Registers[0xA] = 0xF3   // Binary equivalent is 243
	vm.IndexRegister = 0x28

	op := BinaryCodedDecimalOp{register: 0xA}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, byte(2), vm.Memory[0x28])
	assert.Equal(t, byte(4), vm.Memory[0x29])
	assert.Equal(t, byte(3), vm.Memory[0x2A])
}
