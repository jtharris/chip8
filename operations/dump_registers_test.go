package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"chip8/system"
)

func TestDumpRegistersParser_Matches(t *testing.T) {
	parser := dumpRegistersParser{}

	assert.True(t, parser.matches(0xFB55))
}

func TestDumpRegistersParser_DoesNotMatchStart(t *testing.T) {
	parser := dumpRegistersParser{}

	assert.False(t, parser.matches(0xEB55))
}

func TestDumpRegistersParser_DoesNotMatchEnd(t *testing.T) {
	parser := dumpRegistersParser{}

	assert.False(t, parser.matches(0xFB58))
}

func TestDumpRegistersParser_DoesNotMatchMiddle(t *testing.T) {
	parser := dumpRegistersParser{}

	assert.False(t, parser.matches(0xFB65))
}

func TestDumpRegistersParser_CreateOp(t *testing.T) {
	parser := dumpRegistersParser{}
	expected := DumpRegistersOp{topRegister: 0xB}

	assert.Equal(t, expected, parser.createOp(0xFB55))
}

func TestDumpRegistersOp_String(t *testing.T) {
	op := DumpRegistersOp{topRegister: 0x7}

	assert.Equal(t, "dump_registers(V7, &I)", op.String())
}

func TestDumpRegistersOp_Execute(t *testing.T) {
	// Given
	vm := system.VirtualMachine{}
	vm.Registers[0x0] = 0x0A
	vm.Registers[0x1] = 0x0B
	vm.Registers[0x2] = 0x0C
	vm.Registers[0x3] = 0x0D
	vm.Registers[0x4] = 0x04	// Not used
	vm.IndexRegister = 0x10

	op := DumpRegistersOp{topRegister: 0x3}

	// When
	op.Execute(&vm)

	// Then
	assert.Equal(t, byte(0x0A), vm.Memory[0x10])
	assert.Equal(t, byte(0x0B), vm.Memory[0x11])
	assert.Equal(t, byte(0x0C), vm.Memory[0x12])
	assert.Equal(t, byte(0x0D), vm.Memory[0x13])

	// Also be sure that the next memory location was untouched
	assert.Equal(t, byte(0x00), vm.Memory[0x14])
}