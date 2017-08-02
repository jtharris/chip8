package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"chip8/system"
)

func TestClearParser_Matches(t *testing.T) {
	parser := ClearParser{}

	assert.True(t, parser.Matches(0x00E0))
}

func TestClearParser_DoesNotMatch(t *testing.T) {
	parser := ClearParser{}

	assert.False(t, parser.Matches(0x10E0))
}

func TestClearParser_CreateOp(t *testing.T) {
	parser := ClearParser{}
	expected := ClearOp{}

	assert.Equal(t, expected, parser.CreateOp(0x00E0))
}

func TestClearOp_String(t *testing.T) {
	op := ClearOp{}

	assert.Equal(t, "Clear Screen", op.String())
}

func TestDrawOp_Execute(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}

	// Put some arbitrary pixel data
	vm.Pixels[1] = uint64(0xF68A092300007D81)
	vm.Pixels[12] = uint64(0x000009180040777E)
	vm.Pixels[31] = uint64(0x368D052304007265)

	// When
	ClearOp{}.Execute(&vm)

	// Then
	for row := 0; row < len(vm.Pixels); row++ {
		assert.Equal(t, uint64(0), vm.Pixels[row])
	}
}
