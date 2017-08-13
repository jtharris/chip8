package operations

import (
	"chip8/system"
	"fmt"
)

type UnknownOp struct{
	code system.OpCode
}

func (o UnknownOp) String() string {
	return "Unknown Operation"
}

func (o UnknownOp) Execute(machine *system.VirtualMachine) {
	panic(fmt.Sprint("Unknown opcode:  ", o.code))
}
