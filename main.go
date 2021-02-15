package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"mu0-assembler/mu0"
)

func main() {
	fileData, err := ioutil.ReadFile("assembly.mu0")
	if err != nil {
		panic(err)
	}

	instructions, err := mu0.ParseAssembly(string(fileData))
	if err != nil {
		panic(err)
	}

	fmt.Println("Parsed assembly:")
	var machineCode []byte
	for _, instruction := range instructions {
		fmt.Println(instruction.Disassemble())
		machineCode = append(machineCode, instruction.Assemble())
	}

	fmt.Println()
	fmt.Println("Machine code:")
	fmt.Println(hex.EncodeToString(machineCode))
}
