package operations

import (
	"fmt"
	"chip8/system"
)

type IfRegisterParser struct {}
func(p IfRegisterParser) Matches(opcode OpCode) bool {
	// TODO:  There is probably a more efficient way to do this
	opString := opcode.String()
	return opString[0] == '5' && opString[3] == '0'
}

func(p IfRegisterParser) CreateOp(opcode OpCode) Operation {
	return IfRegisterOp{
		register1: uint8(opcode & 0x0F00 >> 8),
		register2: uint8(opcode & 0x00F0 >> 4),
	}
}

type IfRegisterOp struct {
	register1 uint8
	register2 uint8
}
func(o IfRegisterOp) String() string {
	return fmt.Sprintf("If V%X == V%X", o.register1, o.register2)
}

func(o IfRegisterOp) Execute(machine *system.VirtualMachine) {
	// TODO:  Get this going
}
