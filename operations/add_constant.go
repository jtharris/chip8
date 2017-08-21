package operations

import (
	"fmt"
	"chip8/system"
)

type AddConstantParser struct {}
func (p AddConstantParser) Matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0x7
}

func (p AddConstantParser) CreateOp(opcode system.OpCode) Operation {
	return AddConstantOp{
		register: byte(opcode & 0x0F00 >> 8),
		value: byte(opcode),
	}
}

type AddConstantOp struct {
	register byte
	value byte
}
func (o AddConstantOp) String() string {
	return fmt.Sprintf("V%X += %X", o.register, o.value)
}

func (o AddConstantOp) Execute(vm *system.VirtualMachine) {
	vm.Registers[o.register] += o.value
}
