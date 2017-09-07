package operations

import (
	"fmt"
	"github.com/jtharris/chip8/system"
)

// Parser for IfConstantOp
type ifConstantParser struct{}

func (p ifConstantParser) matches(opcode system.OpCode) bool {
	return opcode>>12 == 0x3
}

func (p ifConstantParser) createOp(opcode system.OpCode) Operation {
	return IfConstantOp{
		register: uint8(opcode & 0x0F00 >> 8),
		value:    uint8(opcode & 0x00FF),
	}
}

// IfConstantOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#3xkk
type IfConstantOp struct {
	register uint8
	value    uint8
}

// String returns a text representation of this operation
func (o IfConstantOp) String() string {
	return fmt.Sprintf("If V%X == %X", o.register, o.value)
}

// Execute this operation on the given virtual machine
func (o IfConstantOp) Execute(vm *system.VirtualMachine) {
	if vm.Registers[o.register] == o.value {
		vm.IncrementPC()
	}
}
