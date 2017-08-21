package operations

import (
	"fmt"
	"chip8/system"
)

type ifNotConstantParser struct {}
func(p ifNotConstantParser) matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0x4
}

func(p ifNotConstantParser) createOp(opcode system.OpCode) Operation {
	return IfNotConstantOp{
		register: uint8(opcode & 0x0F00 >> 8),
		value: uint8(opcode & 0x00FF),
	}
}

type IfNotConstantOp struct {
	register uint8
	value uint8
}
func(o IfNotConstantOp) String() string {
	return fmt.Sprintf("If V%X != %X", o.register, o.value)
}

func(o IfNotConstantOp) Execute(vm *system.VirtualMachine) {
	if vm.Registers[o.register] != o.value {
		vm.IncrementPC()
	}
}
