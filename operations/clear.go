package operations

import (
	"chip8/system"
)


type clearParser struct {}
func(p clearParser) matches(opcode system.OpCode) bool {
	return opcode == 0x00E0
}

func(p clearParser) createOp(opcode system.OpCode) Operation {
	return ClearOp{}
}

type ClearOp struct {}
func(o ClearOp) String() string {
	return "Clear Screen"
}

func(o ClearOp) Execute(vm *system.VirtualMachine) {
	for row := 0; row < len(vm.Pixels); row++ {
		vm.Pixels[row] = 0
	}
}
