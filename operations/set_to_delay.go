package operations

import (
	"fmt"
	"chip8/system"
)

type setToDelayParser struct {}
func(p setToDelayParser) matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0xF && byte(opcode) == 0x07
}

func(p setToDelayParser) createOp(opcode system.OpCode) Operation {
	return SetToDelayOp{
		register: byte(opcode & 0x0F00 >> 8),
	}
}

type SetToDelayOp struct {
	register byte
}
func(o SetToDelayOp) String() string {
	return fmt.Sprintf("V%X = delay_timer", o.register)
}

func(o SetToDelayOp) Execute(vm *system.VirtualMachine) {
	vm.Registers[o.register] = vm.DelayTimer
}
