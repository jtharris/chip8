package operations

import (
	"fmt"
	"chip8/system"
)

type ShiftRightParser struct {}
func(p ShiftRightParser) Matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0x8 && opcode & 0x000F == 0x6
}

func(p ShiftRightParser) CreateOp(opcode system.OpCode) Operation {
	return ShiftRightOp{
		register: byte(opcode & 0x0F00 >> 8),
	}
}

type ShiftRightOp struct {
	register byte
}
func(o ShiftRightOp) String() string {
	return fmt.Sprintf("V%X >> 1", o.register)
}

func(o ShiftRightOp) Execute(machine *system.VirtualMachine) {
	val := machine.Registers[o.register]

	machine.Registers[0xF] = val & 0x01
	machine.Registers[o.register] = val >> 1
}
