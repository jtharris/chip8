package operations

import (
	"fmt"
	"chip8/system"
)

// Parser for BitwiseOrOp
type bitwiseOrParser struct {}
func(p bitwiseOrParser) matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0x8 && opcode & 0x000F == 0x1
}

func(p bitwiseOrParser) createOp(opcode system.OpCode) Operation {
	return BitwiseOrOp{
		register1: byte(opcode & 0x0F00 >> 8),
		register2: byte(opcode & 0x00F0 >> 4),
	}
}

// BitwiseOrOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#8xy1
type BitwiseOrOp struct {
	register1 byte
	register2 byte
}

// String returns a text representation of this operation
func(o BitwiseOrOp) String() string {
	return fmt.Sprintf("V%X = V%X | V%X", o.register1, o.register1, o.register2)
}

// Execute this operation on the given virtual machine
func(o BitwiseOrOp) Execute(vm *system.VirtualMachine) {
	vm.Registers[o.register1] |= vm.Registers[o.register2]
}
