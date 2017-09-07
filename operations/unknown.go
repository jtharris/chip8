package operations

import (
	"github.com/jtharris/chip8/system"
	"fmt"
)

// UnknownOp represents an instruction that is not recognized - a Null Object
type UnknownOp struct {
	code system.OpCode
}

// String returns a text representation of this operation
func (o UnknownOp) String() string {
	return fmt.Sprint("Unknown Operation:  ", o.code)
}

// Execute panics if invoked, as this represents an unknown or unimplemented instruction
func (o UnknownOp) Execute(machine *system.VirtualMachine) {
	panic(o.String())
}
