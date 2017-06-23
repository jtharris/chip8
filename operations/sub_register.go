package operations

import (
	"fmt"
	"chip8/system"
)

type SubtractRegisterParser struct {}
func(p SubtractRegisterParser) Matches(opcode OpCode) bool {
	// TODO:  There is probably a more efficient way to do this
	opString := opcode.String()
	return opString[0] == '8' && opString[3] == '5'
}

func(p SubtractRegisterParser) CreateOp(opcode OpCode) Operation {
	return SubtractRegisterOp{
		register1: uint8(opcode & 0x0F00 >> 8),
		register2: uint8(opcode & 0x00F0 >> 4),
	}
}

type SubtractRegisterOp struct {
	register1 uint8
	register2 uint8
}
func(o SubtractRegisterOp) String() string {
	return fmt.Sprintf("V%X -= V%X", o.register1, o.register2)
}

func(o SubtractRegisterOp) Execute(machine *system.VirtualMachine) {
	// TODO:  Get this going
}
