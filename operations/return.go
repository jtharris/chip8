package operations

import (
	"github.com/jtharris/chip8/system"
)

// Parser for ReturnOp
type returnParser struct{}

func (p returnParser) matches(opcode system.OpCode) bool {
	return opcode == 0x00EE
}

func (p returnParser) createOp(opcode system.OpCode) Operation {
	return ReturnOp{}
}

// ReturnOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#00EE
type ReturnOp struct{}

// String returns a text representation of this operation
func (o ReturnOp) String() string {
	return "Return from subroutine"
}

// Execute this operation on the given virtual machine
func (o ReturnOp) Execute(vm *system.VirtualMachine) {
	lastItem := len(vm.Stack) - 1

	vm.ProgramCounter = vm.Stack[lastItem]
	vm.Stack = vm.Stack[:lastItem]
}
