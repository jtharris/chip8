package operations

import (
	"fmt"
	"github.com/jtharris/chip8/system"
	"time"
)

// Parser for GetKeyOp
type getKeyParser struct{}

func (p getKeyParser) matches(opcode system.OpCode) bool {
	return opcode>>12 == 0xF && byte(opcode) == 0x0A
}

func (p getKeyParser) createOp(opcode system.OpCode) Operation {
	return GetKeyOp{
		register: byte(opcode>>8) & 0x0F,
	}
}

// GetKeyOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#Fx0A
type GetKeyOp struct{ register byte }

// String returns a text representation of this operation
func (o GetKeyOp) String() string {
	return fmt.Sprintf("V%X = get_key()", o.register)
}

// Execute this operation on the given virtual machine
func (o GetKeyOp) Execute(vm *system.VirtualMachine) {
	vm.Running = false

	for {
		for key := range vm.Keyboard {
			if vm.Keyboard[key] {
				vm.Registers[o.register] = byte(key)
				vm.Running = true

				return
			}
		}

		// sleep
		time.Sleep(5 * time.Millisecond)
	}
}
