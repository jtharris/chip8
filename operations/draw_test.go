package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"chip8/system"
)

func TestDrawParser_Matches(t *testing.T) {
	parser := DrawParser{}

	assert.True(t, parser.Matches(0xD076))
}

func TestDrawParser_DoesNotMatch(t *testing.T) {
	parser := DrawParser{}

	assert.False(t, parser.Matches(0xE076))
}

func TestDrawParser_CreateOp(t *testing.T) {
	parser := DrawParser{}
	expected := DrawOp{xRegister: 0x8, yRegister: 0xD, height: 0x4}

	assert.Equal(t, expected, parser.CreateOp(0xD8D4))
}

func TestDrawOp_String(t *testing.T) {
	op := DrawOp{xRegister: 0xC, yRegister: 0x3, height: 0x8}

	assert.Equal(t, "Draw Screen (VC, V3) Height: 8", op.String())
}

func TestDrawOp_ExecuteBlank(t *testing.T) {
	// Given
	op := DrawOp{xRegister: 0x0, yRegister: 0x1, height: 0x5}
	vm := system.NewVirtualMachine()	// Create vm with font data in memory
	vm.Registers[0x0] = 0x8
	vm.Registers[0x1] = 0x5
	vm.IndexRegister = 0x1E	// Render the '6' character starting at memory slot 30 (6 * 5) because each char is 5 bytes

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, uint64(0x00F0000000000000), vm.Pixels[5])
	assert.Equal(t, uint64(0x0080000000000000), vm.Pixels[6])
	assert.Equal(t, uint64(0x00F0000000000000), vm.Pixels[7])
	assert.Equal(t, uint64(0x0090000000000000), vm.Pixels[8])
	assert.Equal(t, uint64(0x00F0000000000000), vm.Pixels[9])

	// Since there were no pixels flipped from 'on' to 'off', make sure VF remains 0
	assert.Equal(t, byte(0), vm.Registers[0xF])
}

func TestDrawOp_ExecuteFlipped(t *testing.T) {
	// Given
	op := DrawOp{xRegister: 0x0, yRegister: 0x1, height: 0x5}
	vm := system.NewVirtualMachine()	// Create vm with font data in memory
	vm.Registers[0x0] = 0x38
	vm.Registers[0x1] = 0x2
	vm.IndexRegister = 0x00	// Render the '0' character starting at memory slot 0 (0 * 5)

	// Set Pixels to 'on' for all of the pixels' last byte
	for row := 2; row < 7; row++ {
		vm.Pixels[row] = uint64(0xFF)
	}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, uint64(0x000000000000000F), vm.Pixels[2])
	assert.Equal(t, uint64(0x000000000000006F), vm.Pixels[3])
	assert.Equal(t, uint64(0x000000000000006F), vm.Pixels[4])
	assert.Equal(t, uint64(0x000000000000006F), vm.Pixels[5])
	assert.Equal(t, uint64(0x000000000000000F), vm.Pixels[6])

	// Since there were pixels flipped from 'on' to 'off', make sure VF is set
	assert.Equal(t, byte(1), vm.Registers[0xF])
}
