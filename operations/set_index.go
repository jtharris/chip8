package operations

import (
	"fmt"
	"chip8/system"
)

type setIndexParser struct {}
func(p setIndexParser) matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0xA
}

func(p setIndexParser) createOp(opcode system.OpCode) Operation {
	return SetIndexOp{
		value: uint16(opcode) & 0x0FFF,
	}
}

type SetIndexOp struct {
	value uint16
}
func(o SetIndexOp) String() string {
	return fmt.Sprintf("I = %X", o.value)
}

func(o SetIndexOp) Execute(vm *system.VirtualMachine) {
	vm.IndexRegister = o.value
}
