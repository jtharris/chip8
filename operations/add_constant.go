package operations

import (
	"fmt"
	"chip8/system"
)

type AddConstantParser struct {}
func(p AddConstantParser) Matches(opcode system.OpCode) bool {
	return opcode.String()[0] == '7'
}

func(p AddConstantParser) CreateOp(opcode system.OpCode) Operation {
	return AddConstantOp{
		register: byte(opcode & 0x0F00 >> 8),
		value: byte(opcode & 0x00FF),
	}
}

type AddConstantOp struct {
	register byte
	value byte
}
func(o AddConstantOp) String() string {
	return fmt.Sprintf("V%X += %X", o.register, o.value)
}

func(o AddConstantOp) Execute(machine *system.VirtualMachine) {
	// Set the overflow register first
	machine.Registers[0xF] = byte((uint16(machine.Registers[o.register]) + uint16(o.value)) >> 8)

	// Then add the value
	machine.Registers[o.register] += o.value
}
