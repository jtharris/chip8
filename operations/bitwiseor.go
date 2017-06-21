package operations

import (
	"fmt"
	"chip8/system"
)

type BitwiseOrParser struct {}
func(bop BitwiseOrParser) Matches(opcode OpCode) bool {
	// TODO:  There is probably a more efficient way to do this
	opString := opcode.String()
	return opString[0] == '8' && opString[3] == '1'
}

func(bop BitwiseOrParser) CreateOp(opcode OpCode) Operation {
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

func(o BitwiseOrOp) Execute(machine *system.VirtualMachine) {
	// TODO:  Get this going
}
