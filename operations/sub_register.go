package operations

import (
	"fmt"
	"chip8/system"
)

type SubtractRegisterParser struct {}
func(p SubtractRegisterParser) Matches(opcode system.OpCode) bool {
	opString := opcode.String()
	return opString[0] == '8' && opString[3] == '5'
}

func(p SubtractRegisterParser) CreateOp(opcode system.OpCode) Operation {
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

func(o SubtractRegisterOp) Execute(machine *system.VirtualMachine) {
	val1 := machine.Registers[o.register1]
	val2 := machine.Registers[o.register2]

	if val1 < val2 {
		machine.Registers[0xF] = 0x0
	} else {
		machine.Registers[0xF] = 0x1
	}

	machine.Registers[o.register1] = val1 - val2
}
