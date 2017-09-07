package operations

import (
	"fmt"
	"github.com/jtharris/chip8/system"
)

// Parser for CallOp
type callParser struct{}

func (p callParser) matches(opcode system.OpCode) bool {
	return opcode>>12 == 0x2
}

func (p callParser) createOp(opcode system.OpCode) Operation {
	return CallOp{
		address: uint16(opcode) & 0x0FFF,
	}
}

// CallOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#2nnn
type CallOp struct {
	address uint16
}

// String returns a text representation of this operation
func (o CallOp) String() string {
	return fmt.Sprintf("Call subroutine at:  %X", o.address)
}

// Execute this operation on the given virtual machine
func (o CallOp) Execute(vm *system.VirtualMachine) {
	// Push the current address onto the call stack
	vm.Stack = append(vm.Stack, vm.ProgramCounter)

	// Set the program counter
	vm.ProgramCounter = o.address - 2
}
