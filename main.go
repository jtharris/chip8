package main

import (
	"io/ioutil"
	"os"
	"chip8/system"
	"chip8/operations"
	"time"
)

//func init() {
//	runtime.LockOSThread()
//}

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

	//display := system.TerminalDisplay{}
	display := system.OpenGLDisplay{}

	go Run(&vm)
	go Timers(&vm)

	display.Start(&vm)
}

func Timers(vm *system.VirtualMachine) {
	ticker := time.NewTicker(time.Microsecond * 16667)	// Running at 60 Hz

	for range ticker.C {
		// TODO:  Move to a Cycle method?
		// Decrement timers
		if vm.DelayTimer > 0 {
			vm.DelayTimer--
		}

		if vm.SoundTimer > 0 {
			vm.SoundTimer--
		}
	}
}

func Run(vm *system.VirtualMachine) {
	ticker := time.NewTicker(time.Millisecond * 10)	// Running at 60 Hz

	for range ticker.C {
		opcode := vm.CurrentOpcode()
		op := operations.CreateOperation(opcode)
		op.Execute(vm)
		vm.ProgramCounter += 2
	}
}
