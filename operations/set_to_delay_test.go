package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"chip8/system"
)

func TestSetToDelayParser_Matches(t *testing.T) {
	parser := SetToDelayParser{}

	assert.True(t, parser.Matches(0xF407))
}

func TestSetToDelayParser_DoesNotMatch(t *testing.T) {
	parser := SetToDelayParser{}

	assert.False(t, parser.Matches(0xE407))
}

func TestSetToDelayParser_CreateOp(t *testing.T) {
	parser := SetToDelayParser{}
	expected := SetToDelayOp{register: 0x7}

	assert.Equal(t, expected, parser.CreateOp(0xF707))
}

func TestSetToDelayOp_String(t *testing.T) {
	op := SetToDelayOp{register: 0xA}

	assert.Equal(t, "VA = delay_timer", op.String())
}

func TestSetToDelayOp_Execute(t *testing.T) {
	vm := system.VirtualMachine{}
	vm.DelayTimer = 0x8C
	op := SetToDelayOp{register: 0xB}

	op.Execute(&vm)

	assert.Equal(t, byte(0x8C), vm.Registers[0xB])
}