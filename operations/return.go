package operations

import (
	"chip8/system"
)


type ReturnParser struct {}
func(p ReturnParser) Matches(opcode OpCode) bool {
	return opcode == 0x00EE
}

func(p ReturnParser) CreateOp(opcode OpCode) Operation {
	return ReturnOp{}
}

type ReturnOp struct {}
func(o ReturnOp) String() string {
	return "Return from subroutine"
}

func(o ReturnOp) Execute(machine *system.VirtualMachine) {
	lastItem := len(machine.Stack) - 1

	machine.ProgramCounter = machine.Stack[lastItem]
	machine.Stack = machine.Stack[:lastItem]
}
