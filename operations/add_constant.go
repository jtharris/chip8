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
		value: byte(opcode & 0x00FF),
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
	// Set the overflow register first
	vm.Registers[0xF] = byte((uint16(vm.Registers[o.register]) + uint16(o.value)) >> 8)

	// Then add the value
	vm.Registers[o.register] += o.value
}
