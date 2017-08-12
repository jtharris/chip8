package operations

import (
	"chip8/system"
	"fmt"
)

type CallParser struct {}
func(p CallParser) Matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0x2
}

func(p CallParser) CreateOp(opcode system.OpCode) Operation {
	return CallOp{
		address: uint16(opcode) & 0x0FFF,
	}
}

type CallOp struct {
	address uint16
}
func(o CallOp) String() string {
	return fmt.Sprintf("Call subroutine at:  %X", o.address)
}

func(o CallOp) Execute(machine *system.VirtualMachine) {
	// Push the current address onto the call stack
	machine.Stack = append(machine.Stack, machine.ProgramCounter)

	// Set the program counter
	machine.ProgramCounter = o.address - 2
}
