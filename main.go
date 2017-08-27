package main

import (
	"io/ioutil"
	"chip8/system"
	"chip8/operations"
	"time"
	"fmt"
	"flag"
)

// This is the main entrypoint of the program.  Parse any command line flags, read the game binary, and start the
// emulation process.
func main() {
	printOps := flag.Bool("print", false, "Print program opcodes, rather than run the binary")
	useTerm := flag.Bool("terminal", false, "Use the terminal renderer instead of opengl.")
	flag.Parse()

	vm := read(flag.Args()[0])
	if *printOps {
		printOpcodes(vm)
	} else {
		var display system.Display
		if *useTerm {
			display = &system.TerminalDisplay{}
		} else {
			display = &system.OpenGLDisplay{}
		}

		run(vm, display)
	}
}

// Read a file containing a CHIP8 game binary, and return a virtual machine struct with the game loaded into
// memory.
func read(fileName string) *system.VirtualMachine {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	vm := system.NewVirtualMachine()
	vm.Load(data)

	return &vm
}

// Start the three main loops:  machine, timers, and display
func run(vm *system.VirtualMachine, display system.Display) {
	go startMachine(vm)
	go startTimers(vm)
	display.Start(vm)
}

func startTimers(vm *system.VirtualMachine) {
	ticker := time.NewTicker(time.Microsecond * 16667)	// Running at 60 Hz

	for range ticker.C {
		if vm.Running {
			vm.DecrementTimers()

			// Just use the terminal bell for now... beep if the sound timer is positive
			if vm.SoundTimer > 0 {
				fmt.Print("\a")
			}
		}
	}
}

func startMachine(vm *system.VirtualMachine) {
	ticker := time.NewTicker(time.Millisecond * 3)

	for range ticker.C {
		if vm.Running {
			opcode := vm.CurrentOpcode()
			op := operations.CreateOperation(opcode)
			op.Execute(vm)
			vm.IncrementPC()
		}
	}
}

func printOpcodes(vm *system.VirtualMachine) {
	for mem := vm.ProgramCounter; mem < uint16(len(vm.Memory)); mem += 2 {
		opcode := vm.OpCodeAt(mem)

		if opcode > 0 {
			op := operations.CreateOperation(opcode)
			fmt.Printf("%X - %v:  %v\n", mem, opcode, op)
		}
	}
}
