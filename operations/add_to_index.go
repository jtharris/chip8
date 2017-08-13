package operations

import (
	"chip8/system"
	"fmt"
)

type AddToIndexParser struct{}
func (p AddToIndexParser) Matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0xF && byte(opcode) == 0x1E
}

func (p AddToIndexParser) CreateOp(opcode system.OpCode) Operation {
	return AddToIndexOp{
		register: byte(opcode >> 8) & 0x0F,
	}
}

type AddToIndexOp struct {
	register byte
}

func (o AddToIndexOp) String() string {
	return fmt.Sprintf("I += V%X", o.register)
}

func (o AddToIndexOp) Execute(machine *system.VirtualMachine) {
	machine.IndexRegister += uint16(machine.Registers[o.register])
}
