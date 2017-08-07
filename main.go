package main

import (
	"io/ioutil"
	"os"
	"chip8/system"
	"chip8/operations"
	"time"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

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
		vm.DecrementTimers()
	}
}

func Run(vm *system.VirtualMachine) {
	ticker := time.NewTicker(time.Millisecond * 5)

	for range ticker.C {
		opcode := vm.CurrentOpcode()
		op := operations.CreateOperation(opcode)
		op.Execute(vm)
		vm.ProgramCounter += 2
	}
}
