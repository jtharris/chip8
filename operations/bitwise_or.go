package operations

import (
	"fmt"
	"chip8/system"
)

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

type BitwiseOrOp struct {
	register1 byte
	register2 byte
}
func(o BitwiseOrOp) String() string {
	return fmt.Sprintf("V%X = V%X | V%X", o.register1, o.register1, o.register2)
}

func(o BitwiseOrOp) Execute(vm *system.VirtualMachine) {
	vm.Registers[o.register1] |= vm.Registers[o.register2]
}
