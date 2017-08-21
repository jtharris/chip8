package operations

import (
	"fmt"
	"chip8/system"
)

type bitwiseXorParser struct {}
func(p bitwiseXorParser) matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0x8 && opcode & 0x000F == 0x3
}

func(p bitwiseXorParser) createOp(opcode system.OpCode) Operation {
	return BitwiseXorOp{
		register1: byte(opcode & 0x0F00 >> 8),
		register2: byte(opcode & 0x00F0 >> 4),
	}
}

type BitwiseXorOp struct {
	register1 byte
	register2 byte
}
func(o BitwiseXorOp) String() string {
	return fmt.Sprintf("V%X = V%X ^ V%X", o.register1, o.register1, o.register2)
}

func(o BitwiseXorOp) Execute(vm *system.VirtualMachine) {
	vm.Registers[o.register1] ^= vm.Registers[o.register2]
}
