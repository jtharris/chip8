package operations

import (
	"chip8/system"
	"fmt"
)

// Parser for AssignConstantOp
type assignConstantParser struct{}

func (p assignConstantParser) matches(opcode system.OpCode) bool {
	return opcode>>12 == 0x6
}

func (p assignConstantParser) createOp(opcode system.OpCode) Operation {
	return AssignConstantOp{
		register: byte(opcode & 0x0F00 >> 8),
		value:    byte(opcode & 0x00FF),
	}
}

// AssignConstantOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#6xkk
type AssignConstantOp struct {
	register byte
	value    byte
}

// String returns a text representation of this operation
func (o AssignConstantOp) String() string {
	return fmt.Sprintf("V%X = %X", o.register, o.value)
}

// Execute this operation on the given virtual machine
func (o AssignConstantOp) Execute(vm *system.VirtualMachine) {
	vm.Registers[o.register] = o.value
}
