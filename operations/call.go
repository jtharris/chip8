package operations

import (
	"chip8/system"
	"fmt"
)

type CallParser struct {}
func(cp CallParser) Matches(opcode OpCode) bool {
	// TODO:  There is probably a more efficient way to do this
	return opcode.String()[0] == '2'
}

func(cp CallParser) CreateOp(opcode OpCode) Operation {
	return CallOp{
		address: int16(opcode) & 0x0FFF,
	}
}

type CallOp struct {
	address int16
}
func(o CallOp) String() string {
	return fmt.Sprint("Call subroute at:  ", o.address)
}

func(o CallOp) Execute(machine *system.VirtualMachine) {
	// TODO:  Get this going
}
