package operations

import (
	"github.com/jtharris/chip8/system"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDelayTimerParser_Matches(t *testing.T) {
	parser := delayTimerParser{}

	assert.True(t, parser.matches(0xf815))
}

func TestDelayTimerParser_DoesNotMatch(t *testing.T) {
	parser := delayTimerParser{}

	assert.False(t, parser.matches(0xf825))
}

func TestDelayTimerParser_CreateOp(t *testing.T) {
	parser := delayTimerParser{}
	expected := DelayTimerOp{register: 0x9}

	assert.Equal(t, expected, parser.createOp(0xf915))
}

func TestDelayTimerOp_String(t *testing.T) {
	op := DelayTimerOp{register: 0xD}

	assert.Equal(t, "DT = VD", op.String())
}

func TestDelayTimerOp_Execute(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Registers[0xD] = 0xA4

	op := DelayTimerOp{register: 0xD}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, vm.DelayTimer, byte(0xA4))
}
