package operations

import (
	"chip8/system"
	"fmt"
)


type DrawParser struct {}
func(cp DrawParser) Matches(opcode OpCode) bool {
	return opcode.String()[0] == 'd'
}

func(cp DrawParser) CreateOp(opcode OpCode) Operation {
	return DrawOp{
		xRegister: byte(opcode & 0x0F00 >> 8),
		yRegister: byte(opcode & 0x00F0 >> 4),
		height: byte(opcode & 0x000F),
	}
}

type DrawOp struct {
	xRegister byte
	yRegister byte
	height byte
}
func(o DrawOp) String() string {
	return fmt.Sprintf("Draw Screen (V%X, V%X) Height: %X", o.xRegister, o.yRegister, o.height)
}

func(o DrawOp) Execute(machine *system.VirtualMachine) {
	// TODO:  Get this going
}
