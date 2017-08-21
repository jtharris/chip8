package operations

import (
	"fmt"
	"chip8/system"
)

type subtractRegisterParser struct {}
func(p subtractRegisterParser) matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0x8 && opcode & 0x000F == 0x5
}

func(p subtractRegisterParser) createOp(opcode system.OpCode) Operation {
	return SubtractRegisterOp{
		register1: uint8(opcode & 0x0F00 >> 8),
		register2: uint8(opcode & 0x00F0 >> 4),
	}
}

type SubtractRegisterOp struct {
	register1 uint8
	register2 uint8
}
func(o SubtractRegisterOp) String() string {
	return fmt.Sprintf("V%X -= V%X", o.register1, o.register2)
}

func(o SubtractRegisterOp) Execute(vm *system.VirtualMachine) {
	val1 := vm.Registers[o.register1]
	val2 := vm.Registers[o.register2]

	if val1 < val2 {
		vm.Registers[0xF] = 0x0
	} else {
		vm.Registers[0xF] = 0x1
	}

	vm.Registers[o.register1] = val1 - val2
}
