package system

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewVirtualMachineHasFonts(t *testing.T) {
	vm := NewVirtualMachine()

	// Spot check beginning and end
	assert.Equal(t, byte(0xF0), vm.Memory[0x0])
	assert.Equal(t, byte(0x80), vm.Memory[0x4F])
}

func TestVirtualMachine_LoadPC(t *testing.T) {
	vm := NewVirtualMachine()
	vm.Load([]byte{})

	assert.Equal(t, uint16(0x200), vm.ProgramCounter)
}

func TestVirtualMachine_LoadDataOpCode(t *testing.T) {
	vm := NewVirtualMachine()
	vm.Load([]byte{0x0F, 0xA5, 0x04, 0x7D})

	assert.Equal(t, OpCode(0x0FA5), vm.OpCodeAt(0x200))
	assert.Equal(t, OpCode(0x047D), vm.OpCodeAt(0x202))
}

func TestVirtualMachine_CurrentOpcode(t *testing.T) {
	vm := NewVirtualMachine()
	vm.Load([]byte{0x0, 0x0, 0x0F, 0xA5, 0x04, 0x7D})
	vm.ProgramCounter = 0x204

	assert.Equal(t, OpCode(0x047D), vm.CurrentOpcode())
}

func TestVirtualMachine_IncrementPC(t *testing.T) {
	// Given
	vm := NewVirtualMachine()
	vm.Load([]byte{})

	// When
	vm.IncrementPC()

	// Then
	assert.Equal(t, uint16(0x202), vm.ProgramCounter)
}

func TestVirtualMachine_DecrementTimersWhenZero(t *testing.T) {
	vm := NewVirtualMachine()

	// When
	vm.DecrementTimers()

	// Then
	assert.Equal(t, byte(0), vm.DelayTimer)
	assert.Equal(t, byte(0), vm.SoundTimer)
}

func TestVirtualMachine_DecrementTimersWhenNonZero(t *testing.T) {
	vm := NewVirtualMachine()
	vm.DelayTimer = 0x0A
	vm.SoundTimer = 0x03

	// When
	vm.DecrementTimers()

	// Then
	assert.Equal(t, byte(0x09), vm.DelayTimer)
	assert.Equal(t, byte(0x02), vm.SoundTimer)
}

func TestVirtualMachine_PixelSetAt(t *testing.T) {
	vm := NewVirtualMachine()
	vm.Pixels[15] = 0xF

	assert.True(t, vm.PixelSetAt(63, 15))
	assert.True(t, vm.PixelSetAt(60, 15))

	assert.False(t, vm.PixelSetAt(60, 16))
	assert.False(t, vm.PixelSetAt(59, 15))
}
