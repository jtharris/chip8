package operations

import (
	"chip8/system"
	"fmt"
)

type GotoParser struct {}
func(p GotoParser) Matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0x1
}

func(p GotoParser) CreateOp(opcode system.OpCode) Operation {
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
	machine.ProgramCounter = o.address - 2
}
