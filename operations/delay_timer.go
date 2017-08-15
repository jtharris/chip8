package operations

import (
	"fmt"
	"chip8/system"
)

type DelayTimerParser struct {}
func(p DelayTimerParser) Matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0xF && byte(opcode) == 0x15
}

func(p DelayTimerParser) CreateOp(opcode system.OpCode) Operation {
	return DelayTimerOp{
		register: byte(opcode & 0x0F00 >> 8),
	}
}

type DelayTimerOp struct {
	register byte
}
func(o DelayTimerOp) String() string {
	return fmt.Sprintf("delay_timer = V%X", o.register)
}

func(o DelayTimerOp) Execute(vm *system.VirtualMachine) {
	vm.DelayTimer = vm.Registers[o.register]
}
