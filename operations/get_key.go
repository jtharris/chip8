package operations

import (
	"chip8/system"
	"fmt"
	"time"
)

type GetKeyParser struct {}
func(p GetKeyParser) Matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0xF && byte(opcode) == 0x0A
}

func(p GetKeyParser) CreateOp(opcode system.OpCode) Operation {
	return GetKeyOp{
		register: byte(opcode >> 8) & 0x0F,
	}
}


type GetKeyOp struct { register byte }
func (o GetKeyOp) String() string {
	return fmt.Sprintf("V%X = get_key()", o.register)
}

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
