package operations

import (
	"fmt"
	"chip8/system"
)

type AddRegisterParser struct {}
func(p AddRegisterParser) Matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0x8 && opcode & 0x000F == 0x4
}

func(p AddRegisterParser) CreateOp(opcode system.OpCode) Operation {
	return AddRegisterOp{
		register1: byte(opcode & 0x0F00 >> 8),
		register2: byte(opcode & 0x00F0 >> 4),
	}
}

type AddRegisterOp struct {
	register1 byte
	register2 byte
}
func(o AddRegisterOp) String() string {
	return fmt.Sprintf("V%X += V%X", o.register1, o.register2)
}

func(o AddRegisterOp) Execute(machine *system.VirtualMachine) {
	// Set the overflow register first
	machine.Registers[0xF] = byte((uint16(machine.Registers[o.register1]) + uint16(machine.Registers[o.register2])) >> 8)

	// Then add the value
	machine.Registers[o.register1] += machine.Registers[o.register2]
}
