package main

import (
	"io/ioutil"
	"os"
	"chip8/system"
	"chip8/operations"
	"fmt"
	"time"
)

func main() {
	// TODO:  All types of error checking
	read(os.Args[1])
}

func read(fileName string) {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	vm := system.NewVirtualMachine()
	vm.Load(data)

	display := system.TerminalDisplay{}
	display.Initialize()

	go DisplayLoop(&vm, &display)
	Run(&vm, &display)
	//PrintProgram(&vm)
}

func Run(vm *system.VirtualMachine, display system.Display) {
	// TODO:  Very naive implementation to get going
	ticker := time.NewTicker(time.Microsecond * 16667)	// Running at 60 Hz

	for range ticker.C {
		opcode := vm.CurrentOpcode()
		op := operations.CreateOperation(opcode)

		display.UpdateKeys(vm)
		op.Execute(vm)

		// TODO:  Move to a Cycle method?
		// Decrement timers
		if vm.DelayTimer > 0 {
			vm.DelayTimer--
		}

		if vm.SoundTimer > 0 {
			vm.SoundTimer--
		}

		vm.ProgramCounter += 2
	}
}

func DisplayLoop(vm *system.VirtualMachine, display system.Display) {
	ticker := time.NewTicker(time.Millisecond * 500)  // Very slow refresh rate... 4Hz

	for range ticker.C {
		display.Render(vm)
	}
}

func PrintProgram(vm *system.VirtualMachine) {
	for mem := vm.ProgramCounter; mem < uint16(len(vm.Memory)); mem +=2 {
		opcode := vm.OpCodeAt(mem)

		if opcode == 0x0 {
			return
		}

		fmt.Printf("%X:  ", mem)
		fmt.Println(operations.CreateOperation(opcode).String())
	}
}
