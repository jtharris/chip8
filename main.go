package main

import (
	"io/ioutil"
	"os"
	"chip8/system"
	"chip8/operations"
	"time"
	"runtime"
	"fmt"
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

	//PrintOpcodes(&vm)
	//return

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
	ticker := time.NewTicker(time.Millisecond * 3)

	for range ticker.C {
		opcode := vm.CurrentOpcode()
		op := operations.CreateOperation(opcode)
		op.Execute(vm)
		vm.ProgramCounter += 2
	}
}

func PrintOpcodes(vm *system.VirtualMachine) {
	for mem := vm.ProgramCounter; mem < uint16(len(vm.Memory)); mem += 2 {
		opcode := vm.OpCodeAt(mem)

		if opcode > 0 {
			op := operations.CreateOperation(opcode)
			fmt.Println(mem, " - ", opcode, ":  ", op.String())
		}
	}
}

