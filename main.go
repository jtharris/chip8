package main

import (
	"fmt"
	"io/ioutil"
	"os"
	//"chip8/system"
	"chip8/operations"
)

func main() {
	// TODO:  All types of error checking
	read(os.Args[1])
	//system.Start()
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

		fmt.Print("Memory: ")
		fmt.Printf("%X", 512 + i)
		fmt.Print(":  ")
		fmt.Print(opCode.String())
		fmt.Print(" - ")
		fmt.Println(op.String())
	}
}
