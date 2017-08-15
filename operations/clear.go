package operations

import (
	"chip8/system"
)


type ClearParser struct {}
func(p ClearParser) Matches(opcode system.OpCode) bool {
	return opcode == 0x00E0
}

func(p ClearParser) CreateOp(opcode system.OpCode) Operation {
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
