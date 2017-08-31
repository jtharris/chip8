package operations

import (
	"chip8/system"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetToDelayParser_Matches(t *testing.T) {
	parser := setToDelayParser{}

	assert.True(t, parser.matches(0xF407))
}

func TestSetToDelayParser_DoesNotMatch(t *testing.T) {
	parser := setToDelayParser{}

	assert.False(t, parser.matches(0xE407))
}

func TestSetToDelayParser_CreateOp(t *testing.T) {
	parser := setToDelayParser{}
	expected := SetToDelayOp{register: 0x7}

	assert.Equal(t, expected, parser.createOp(0xF707))
}

func TestSetToDelayOp_String(t *testing.T) {
	op := SetToDelayOp{register: 0xA}

	assert.Equal(t, "VA = DT", op.String())
}

func TestSetToDelayOp_Execute(t *testing.T) {
	vm := system.VirtualMachine{}
	vm.DelayTimer = 0x8C
	op := SetToDelayOp{register: 0xB}

	op.Execute(&vm)

	assert.Equal(t, byte(0x8C), vm.Registers[0xB])
}
