package operations

import (
	"chip8/system"
	"fmt"
)

// Parser for AddToIndexOp
type addToIndexParser struct{}
func (p addToIndexParser) matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0xF && byte(opcode) == 0x1E
}

func (p addToIndexParser) createOp(opcode system.OpCode) Operation {
	return AddToIndexOp{
		register: byte(opcode >> 8) & 0x0F,
	}
}

// AddToIndexOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#Fx1E
type AddToIndexOp struct {
	register byte
}

// String returns a text representation of this operation
func (o AddToIndexOp) String() string {
	return fmt.Sprintf("I += V%X", o.register)
}

// Execute this operation on the given virtual machine
func (o AddToIndexOp) Execute(vm *system.VirtualMachine) {
	vm.IndexRegister += uint16(vm.Registers[o.register])
}
