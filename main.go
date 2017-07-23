package main

import (
	"fmt"
	"io/ioutil"
	//"os"
	"chip8/operations"
	"chip8/system"
)

func main() {
	// TODO:  All types of error checking
	//read(os.Args[1])
	system.Start()
}

func read(fileName string) {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	for i :=0; i < len(data); i += 2 {
		bytes := data[i:i + 2]
		var opCode operations.OpCode = operations.OpCode((uint16(bytes[0]) << 8) + uint16(bytes[1]))

		op := operations.CreateOperation(opCode)

		fmt.Print("Opcode ")
		fmt.Print(i / 2)
		fmt.Print(":  ")
		fmt.Print(opCode.String())
		fmt.Print(" - ")
		fmt.Println(op.String())
	}
}
