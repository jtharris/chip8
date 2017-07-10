package operations

import (
	"chip8/system"
	"fmt"
)

type CallParser struct {}
func(p CallParser) Matches(opcode OpCode) bool {
	// TODO:  There is probably a more efficient way to do this
	return opcode.String()[0] == '2'
}

func(p CallParser) CreateOp(opcode OpCode) Operation {
	return CallOp{
		address: uint16(opcode) & 0x0FFF,
	}
}

type CallOp struct {
	address uint16
}
func(o CallOp) String() string {
	return fmt.Sprint("Call subroute at:  ", o.address)
}

func(o CallOp) Execute(machine *system.VirtualMachine) {
	// Push the current address onto the call stack
	machine.Stack = append(machine.Stack, machine.ProgramCounter)

	// Set the program counter
	machine.ProgramCounter = o.address
}
