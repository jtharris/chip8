package operations

import (
	"chip8/system"
)


type ReturnParser struct {}
func(p ReturnParser) Matches(opcode system.OpCode) bool {
	return opcode == 0x00EE
}

func(p ReturnParser) CreateOp(opcode system.OpCode) Operation {
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
