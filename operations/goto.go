package operations

import (
	"chip8/system"
	"fmt"
)

type GotoParser struct {}
func(p GotoParser) Matches(opcode OpCode) bool {
	// TODO:  There is probably a more efficient way to do this
	return opcode.String()[0] == '1'
}

func(p GotoParser) CreateOp(opcode OpCode) Operation {
	return GotoOp{
		address: int16(opcode) & 0x0FFF,
	}
}

type GotoOp struct {
	address int16
}
func(o GotoOp) String() string {
	return fmt.Sprint("Goto address:  ", o.address)
}

func(o GotoOp) Execute(machine *system.VirtualMachine) {
	// TODO:  Get this going
}
