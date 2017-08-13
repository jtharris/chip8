package operations

import (
	"fmt"
	"chip8/system"
)

type DumpRegistersParser struct {}
func(p DumpRegistersParser) Matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0xF && byte(opcode) == 0x55
}

func(p DumpRegistersParser) CreateOp(opcode system.OpCode) Operation {
	return DumpRegistersOp{
		topRegister: byte(opcode & 0x0F00 >> 8),
	}
}

type DumpRegistersOp struct {
	topRegister byte
}
func(o DumpRegistersOp) String() string {
	return fmt.Sprintf("dump_registers(V%X, &I)", o.topRegister)
}

func(o DumpRegistersOp) Execute(machine *system.VirtualMachine) {
	for i := byte(0); i <= o.topRegister; i++ {
		machine.Memory[machine.IndexRegister + uint16(i)] = machine.Registers[i]
	}
}
