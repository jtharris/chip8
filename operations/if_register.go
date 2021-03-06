package operations

import (
	"fmt"
	"github.com/jtharris/chip8/system"
)

// Parser for IfRegisterOp
type ifRegisterParser struct{}

func (p ifRegisterParser) matches(opcode system.OpCode) bool {
	return opcode>>12 == 0x5 && opcode&0x000F == 0x0
}

func (p ifRegisterParser) createOp(opcode system.OpCode) Operation {
	return IfRegisterOp{
		register1: uint8(opcode & 0x0F00 >> 8),
		register2: uint8(opcode & 0x00F0 >> 4),
	}
}

// IfRegisterOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#5xy0
type IfRegisterOp struct {
	register1 uint8
	register2 uint8
}

// String returns a text representation of this operation
func (o IfRegisterOp) String() string {
	return fmt.Sprintf("If V%X == V%X", o.register1, o.register2)
}

// Execute this operation on the given virtual machine
func (o IfRegisterOp) Execute(vm *system.VirtualMachine) {
	if vm.Registers[o.register1] == vm.Registers[o.register2] {
		vm.IncrementPC()
	}
}
