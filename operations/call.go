package operations

import (
	"chip8/system"
	"fmt"
)

type callParser struct {}
func(p callParser) matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0x2
}

func(p callParser) createOp(opcode system.OpCode) Operation {
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

func(o CallOp) Execute(vm *system.VirtualMachine) {
	// Push the current address onto the call stack
	vm.Stack = append(vm.Stack, vm.ProgramCounter)

	// Set the program counter
	vm.ProgramCounter = o.address - 2
}
