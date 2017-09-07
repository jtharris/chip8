package main

import (
	"github.com/jtharris/chip8/operations"
	"github.com/jtharris/chip8/system"
	"flag"
	"fmt"
	"io/ioutil"
	"time"
	"os"
)

// This is the main entrypoint of the program.  Parse any command line flags, read the game binary, and start the
// emulation process.
func main() {
	printOps := flag.Bool("print", false, "Print program opcodes, rather than run the binary")
	useTerm := flag.Bool("terminal", false, "Use the terminal renderer instead of opengl.")
	flag.Parse()

	fileName := flag.Arg(0)

	if fileName == "" {
		fmt.Println("A game file is required as the first argument.")
		os.Exit(1)
	}

	vm := read(fileName)
	if *printOps {
		printOpcodes(vm)
	} else {
		var renderer system.Renderer
		if *useTerm {
			renderer = &system.TerminalRenderer{}
		} else {
			renderer = &system.OpenGLRenderer{}
		}

		run(vm, renderer)
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

// Start the three main loops:  machine, timers, and renderer
func run(vm *system.VirtualMachine, renderer system.Renderer) {
	go startMachine(vm)
	go startTimers(vm)
	renderer.Start(vm)
}

func startTimers(vm *system.VirtualMachine) {
	ticker := time.NewTicker(time.Microsecond * 16667) // Running at 60 Hz

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
	// I haven't found any official spec for the clock speed, but this empirically felt pretty good
	ticker := time.NewTicker(time.Millisecond * 3)

	for range ticker.C {
		if vm.Running {
			opcode := vm.CurrentOpCode()
			op := operations.CreateOperation(opcode)
			op.Execute(vm)
			vm.IncrementPC()
		}
	}
}

func printOpcodes(vm *system.VirtualMachine) {
	fmt.Println("MEM - OPCD:  Description")
	fmt.Println("=============================================")
	for mem := vm.ProgramCounter; mem < uint16(len(vm.Memory)); mem += 2 {
		opcode := vm.OpCodeAt(mem)

		if opcode > 0 {
			op := operations.CreateOperation(opcode)
			fmt.Printf("%X - %v:  %v\n", mem, opcode, op)
		}
	}
}
