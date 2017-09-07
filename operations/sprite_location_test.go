package operations

import (
	"github.com/jtharris/chip8/system"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpriteLocationParser_Matches(t *testing.T) {
	parser := spriteLocationParser{}

	assert.True(t, parser.matches(0xF829))
}

func TestSpriteLocationParser_DoesNotMatch(t *testing.T) {
	parser := spriteLocationParser{}

	assert.False(t, parser.matches(0xF819))
}

func TestSpriteLocationParser_CreateOp(t *testing.T) {
	parser := spriteLocationParser{}
	expected := SpriteLocationOp{register: 0xC}

	assert.Equal(t, expected, parser.createOp(0xFC29))
}

func TestSpriteLocationOp_String(t *testing.T) {
	op := SpriteLocationOp{register: 0x9}

	assert.Equal(t, "I = sprite_address(V9)", op.String())
}

func TestSpriteLocationOp_Execute(t *testing.T) {
	// Given
	op := SpriteLocationOp{register: 0x7}
	vm := system.NewVirtualMachine() // using constructor here to populate the font data
	vm.Registers[0x7] = 0x9

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, uint16(0x2D), vm.IndexRegister)
}
