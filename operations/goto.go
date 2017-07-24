package operations

import (
	"chip8/system"
	"fmt"
)

type GotoParser struct {}
func(p GotoParser) Matches(opcode OpCode) bool {
	return opcode.String()[0] == '1'
}

func(p GotoParser) CreateOp(opcode OpCode) Operation {
	return GotoOp{
		address: uint16(opcode) & 0x0FFF,
	}
}

type GotoOp struct {
	address uint16
}
func(o GotoOp) String() string {
	return fmt.Sprintf("Goto: %X", o.address)
}

func(o GotoOp) Execute(machine *system.VirtualMachine) {
	machine.ProgramCounter = o.address - 0x200
}
