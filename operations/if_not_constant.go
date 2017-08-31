package operations

import (
	"chip8/system"
	"fmt"
)

// Parser for IfNotConstantOp
type ifNotConstantParser struct{}

func (p ifNotConstantParser) matches(opcode system.OpCode) bool {
	return opcode>>12 == 0x4
}

func (p ifNotConstantParser) createOp(opcode system.OpCode) Operation {
	return IfNotConstantOp{
		register: uint8(opcode & 0x0F00 >> 8),
		value:    uint8(opcode & 0x00FF),
	}
}

// IfNotConstantOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#4xkk
type IfNotConstantOp struct {
	register uint8
	value    uint8
}

// String returns a text representation of this operation
func (o IfNotConstantOp) String() string {
	return fmt.Sprintf("If V%X != %X", o.register, o.value)
}

// Execute this operation on the given virtual machine
func (o IfNotConstantOp) Execute(vm *system.VirtualMachine) {
	if vm.Registers[o.register] != o.value {
		vm.IncrementPC()
	}
}
