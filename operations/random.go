package operations

import (
	"fmt"
	"chip8/system"
	"math/rand"
)

type RandomParser struct {}
func(p RandomParser) Matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0xC
}

func(p RandomParser) CreateOp(opcode system.OpCode) Operation {
	return RandomOp{
		register: byte(opcode & 0x0F00 >> 8),
		value: byte(opcode),
	}
}

type RandomOp struct {
	register byte
	value byte
}
func(o RandomOp) String() string {
	return fmt.Sprintf("V%X = rand(255) & %X", o.register, o.value)
}

func(o RandomOp) Execute(machine *system.VirtualMachine) {
	machine.Registers[o.register] = byte(rand.Int()) & o.value
}
