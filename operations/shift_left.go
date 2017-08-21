package operations

import (
	"fmt"
	"chip8/system"
)

type shiftLeftParser struct {}
func(p shiftLeftParser) matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0x8 && opcode & 0x000F == 0xE
}

func(p shiftLeftParser) createOp(opcode system.OpCode) Operation {
	return ShiftLeftOp{
		register: byte(opcode & 0x0F00 >> 8),
	}
}

type ShiftLeftOp struct {
	register byte
}
func(o ShiftLeftOp) String() string {
	return fmt.Sprintf("V%X << 1", o.register)
}

func(o ShiftLeftOp) Execute(vm *system.VirtualMachine) {
	val := vm.Registers[o.register]

	vm.Registers[0xF] = val >> 7
	vm.Registers[o.register] = val << 1
}
