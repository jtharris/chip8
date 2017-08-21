package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"chip8/system"
)

func TestSoundTimerParser_Matches(t *testing.T) {
	parser := soundTimerParser{}

	assert.True(t, parser.matches(0xf818))
}

func TestSoundTimerParser_DoesNotMatch(t *testing.T) {
	parser := soundTimerParser{}

	assert.False(t, parser.matches(0xf815))
}

func TestSoundTimerParser_CreateOp(t *testing.T) {
	parser := soundTimerParser{}
	expected := SoundTimerOp{register: 0x9}

	assert.Equal(t, expected, parser.createOp(0xf915))
}

func TestSoundTimerOp_String(t *testing.T) {
	op := SoundTimerOp{register: 0xD}

	assert.Equal(t, "sound_timer = VD", op.String())
}

func TestSoundTimerOp_Execute(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Registers[0xD] = 0xA4

	op := SoundTimerOp{register: 0xD}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, vm.SoundTimer, byte(0xA4))
}