package operations

import (
	"fmt"
	"chip8/system"
)

type IfNotKeyParser struct {}
func(p IfNotKeyParser) Matches(opcode OpCode) bool {
	return opcode.String()[0] == 'e' && byte(opcode) == 0xA1
}

func(p IfNotKeyParser) CreateOp(opcode OpCode) Operation {
	return IfNotKeyOp{
		register: uint8(opcode & 0x0F00 >> 8),
	}
}

type IfNotKeyOp struct {
	register uint8
}
func(o IfNotKeyOp) String() string {
	return fmt.Sprintf("If key != V%X", o.register)
}

func(o IfNotKeyOp) Execute(machine *system.VirtualMachine) {
	// TODO:  Get this going
}
