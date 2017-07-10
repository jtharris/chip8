package operations

import (
	"fmt"
	"chip8/system"
)

type ReverseSubtractRegisterParser struct {}
func(p ReverseSubtractRegisterParser) Matches(opcode OpCode) bool {
	// TODO:  There is probably a more efficient way to do this
	opString := opcode.String()
	return opString[0] == '8' && opString[3] == '7'
}

func(p ReverseSubtractRegisterParser) CreateOp(opcode OpCode) Operation {
	return ReverseSubtractRegisterOp{
		register1: uint8(opcode & 0x0F00 >> 8),
		register2: uint8(opcode & 0x00F0 >> 4),
	}
}

type ReverseSubtractRegisterOp struct {
	register1 uint8
	register2 uint8
}
func(o ReverseSubtractRegisterOp) String() string {
	return fmt.Sprintf("V%X = V%X - V%X", o.register1, o.register2, o.register1)
}

func(o ReverseSubtractRegisterOp) Execute(machine *system.VirtualMachine) {
	val1 := machine.Registers[o.register1]
	val2 := machine.Registers[o.register2]

	if val1 > val2 {
		machine.Registers[0xF] = 0x0
	} else {
		machine.Registers[0xF] = 0x1
	}

	machine.Registers[o.register1] = val2 - val1
}
