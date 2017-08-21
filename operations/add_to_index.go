package operations

import (
	"chip8/system"
	"fmt"
)

type addToIndexParser struct{}
func (p addToIndexParser) matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0xF && byte(opcode) == 0x1E
}

func (p addToIndexParser) createOp(opcode system.OpCode) Operation {
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

func (o AddToIndexOp) Execute(vm *system.VirtualMachine) {
	vm.IndexRegister += uint16(vm.Registers[o.register])
}
