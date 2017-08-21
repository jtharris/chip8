package operations

import (
	"chip8/system"
)


type returnParser struct {}
func(p returnParser) matches(opcode system.OpCode) bool {
	return opcode == 0x00EE
}

func(p returnParser) createOp(opcode system.OpCode) Operation {
	return ReturnOp{}
}

type ReturnOp struct {}
func(o ReturnOp) String() string {
	return "Return from subroutine"
}

func(o ReturnOp) Execute(vm *system.VirtualMachine) {
	lastItem := len(vm.Stack) - 1

	vm.ProgramCounter = vm.Stack[lastItem]
	vm.Stack = vm.Stack[:lastItem]
}
