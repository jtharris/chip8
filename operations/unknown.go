package operations

import (
	"chip8/system"
	"fmt"
)

// UnknownOp represents an instruction that is not recognized - a Null Object
type UnknownOp struct{
	code system.OpCode
}

// String returns a text representation of this operation
func (o UnknownOp) String() string {
	return "Unknown Operation"
}

// Execute panics if invoked, as this represents an unknown or unimplemented instruction
func (o UnknownOp) Execute(machine *system.VirtualMachine) {
	panic(fmt.Sprint("Unknown opcode:  ", o.code))
}
