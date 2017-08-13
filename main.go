package main

import (
	"io/ioutil"
	"chip8/system"
	"chip8/operations"
	"time"
	"fmt"
	"flag"
)

func main() {
	printOps := flag.Bool("print", false, "Print program opcodes, rather than run the binary")
	flag.Parse()

	vm := read(flag.Args()[0])
	if *printOps {
		printOpcodes(vm)
	} else {
		run(vm)
	}
}

func read(fileName string) *system.VirtualMachine {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	vm := system.NewVirtualMachine()
	vm.Load(data)

	return &vm
}

func run(vm *system.VirtualMachine) {
	//display := system.TerminalDisplay{}
	display := system.OpenGLDisplay{}

	go startMachine(vm)
	go startTimers(vm)
	display.Start(vm)
}

func startTimers(vm *system.VirtualMachine) {
	ticker := time.NewTicker(time.Microsecond * 16667)	// Running at 60 Hz

	for range ticker.C {
		vm.DecrementTimers()

		// Just use the terminal bell for now... beep if the sound timer is positive
		if vm.SoundTimer > 0 {
			fmt.Print("\a")
		}
	}
}

func startMachine(vm *system.VirtualMachine) {
	ticker := time.NewTicker(time.Millisecond * 3)

	for range ticker.C {
		opcode := vm.CurrentOpcode()
		op := operations.CreateOperation(opcode)
		op.Execute(vm)
		vm.IncrementPC()
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
