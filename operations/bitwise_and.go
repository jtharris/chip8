package operations

import (
	"fmt"
	"chip8/system"
)

type BitwiseAndParser struct {}
func(p BitwiseAndParser) Matches(opcode system.OpCode) bool {
	opString := opcode.String()
	return opString[0] == '8' && opString[3] == '2'
}

func(p BitwiseAndParser) CreateOp(opcode system.OpCode) Operation {
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

func(o BitwiseAndOp) Execute(machine *system.VirtualMachine) {
	machine.Registers[o.register1] &= machine.Registers[o.register2]
}
