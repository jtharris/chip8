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
