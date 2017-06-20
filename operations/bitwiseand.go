package operations

import (
	"fmt"
	"chip8/system"
)

type BitwiseAndParser struct {}
func(bap BitwiseAndParser) Matches(opcode OpCode) bool {
	// TODO:  There is probably a more efficient way to do this
	opString := opcode.String()
	return opString[0] == '8' && opString[3] == '2'
}

func(bap BitwiseAndParser) CreateOp(opcode OpCode) Operation {
	return BitwiseAndOp{
		register1: uint8(opcode & 0x0F00 >> 8),
		register2: uint8(opcode & 0x00F0 >> 4),
	}
}

type BitwiseAndOp struct {
	register1 uint8
	register2 uint8
}
func(o BitwiseAndOp) String() string {
	return fmt.Sprintf("V%X = V%X & V%X", o.register1, o.register1, o.register2)
}

func(o BitwiseAndOp) Execute(machine *system.VirtualMachine) {
	// TODO:  Get this going
}
