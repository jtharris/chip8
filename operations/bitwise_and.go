package operations

import (
	"fmt"
	"chip8/system"
)

type bitwiseAndParser struct {}
func(p bitwiseAndParser) matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0x8 && opcode & 0x000F == 0x2
}

func(p bitwiseAndParser) createOp(opcode system.OpCode) Operation {
	return BitwiseAndOp{
		register1: byte(opcode & 0x0F00 >> 8),
		register2: byte(opcode & 0x00F0 >> 4),
	}
}

type BitwiseAndOp struct {
	register1 byte
	register2 byte
}
func(o BitwiseAndOp) String() string {
	return fmt.Sprintf("V%X = V%X & V%X", o.register1, o.register1, o.register2)
}

func(o BitwiseAndOp) Execute(vm *system.VirtualMachine) {
	vm.Registers[o.register1] &= vm.Registers[o.register2]
}
