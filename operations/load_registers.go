package operations

import (
	"fmt"
	"chip8/system"
)

type loadRegistersParser struct {}
func(p loadRegistersParser) matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0xF && byte(opcode) == 0x65
}

func(p loadRegistersParser) createOp(opcode system.OpCode) Operation {
	return LoadRegistersOp{
		topRegister: byte(opcode & 0x0F00 >> 8),
	}
}

type LoadRegistersOp struct {
	topRegister byte
}
func(o LoadRegistersOp) String() string {
	return fmt.Sprintf("load_registers(V%X, &I)", o.topRegister)
}

func(o LoadRegistersOp) Execute(vm *system.VirtualMachine) {
	for i := byte(0); i <= o.topRegister; i++ {
		vm.Registers[i] = vm.Memory[vm.IndexRegister + uint16(i)]
	}
}
