package operations

import (
	"fmt"
	"chip8/system"
)

type SetToDelayParser struct {}
func(cp SetToDelayParser) Matches(opcode OpCode) bool {
	return opcode.String()[0] == 'f' && byte(opcode) == 0x07
}

func(cp SetToDelayParser) CreateOp(opcode OpCode) Operation {
	return SetToDelayOp{
		register: byte(opcode & 0x0F00 >> 8),
	}
}

type SetToDelayOp struct {
	register byte
}
func(o SetToDelayOp) String() string {
	return fmt.Sprintf("Set register %X to delay timer", o.register)
}

func(o SetToDelayOp) Execute(machine *system.VirtualMachine) {
	// TODO:  Get this going
}
