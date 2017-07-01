package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"chip8/system"
)

func TestDelayTimerParser_Matches(t *testing.T) {
	parser := DelayTimerParser{}

	assert.True(t, parser.Matches(0xf815))
}

func TestDelayTimerParser_DoesNotMatch(t *testing.T) {
	parser := DelayTimerParser{}

	assert.False(t, parser.Matches(0xf825))
}

func TestDelayTimerParser_CreateOp(t *testing.T) {
	parser := DelayTimerParser{}
	expected := DelayTimerOp{register: 0x9}

	assert.Equal(t, expected, parser.CreateOp(0xf915))
}

func TestDelayTimerOp_String(t *testing.T) {
	op := DelayTimerOp{register: 0xD}

	assert.Equal(t, "delay_timer = VD", op.String())
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