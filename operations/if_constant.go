package operations

import (
	"fmt"
	"chip8/system"
)

type ifConstantParser struct {}
func(p ifConstantParser) matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0x3
}

func(p ifConstantParser) createOp(opcode system.OpCode) Operation {
	return IfConstantOp{
		register: uint8(opcode & 0x0F00 >> 8),
		value: uint8(opcode & 0x00FF),
	}
}

type IfConstantOp struct {
	register uint8
	value uint8
}
func(o IfConstantOp) String() string {
	return fmt.Sprintf("If V%X == %X", o.register, o.value)
}

func(o IfConstantOp) Execute(vm *system.VirtualMachine) {
	if vm.Registers[o.register] == o.value {
		vm.IncrementPC()
	}
}
