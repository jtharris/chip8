package operations

import (
	"github.com/jtharris/chip8/system"
	"fmt"
)

// Parser for AddConstantOp
type addConstantParser struct{}

func (p addConstantParser) matches(opcode system.OpCode) bool {
	return opcode>>12 == 0x7
}

func (p addConstantParser) createOp(opcode system.OpCode) Operation {
	return AddConstantOp{
		register: byte(opcode & 0x0F00 >> 8),
		value:    byte(opcode),
	}
}

// AddConstantOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#7xkk
type AddConstantOp struct {
	register byte
	value    byte
}

// String returns a text representation of this operation
func (o AddConstantOp) String() string {
	return fmt.Sprintf("V%X += %X", o.register, o.value)
}

// Execute this operation on the given virtual machine
func (o AddConstantOp) Execute(vm *system.VirtualMachine) {
	vm.Registers[o.register] += o.value
}
