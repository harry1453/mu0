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
	var machineCode []uint16
	for _, instruction := range instructions {
		fmt.Println(instruction.Disassemble())
		machineCode = append(machineCode, instruction.Assemble())
	}

	fmt.Println()
	fmt.Println("Machine code:")
	fmt.Println(hex.EncodeToString(mu0.ByteArrayFromMachineCode(machineCode)))

	fmt.Println()
	fmt.Println("Running program...")
	vm, err := mu0.RunProgram(machineCode, true)
	if err != nil {
		fmt.Printf("Error running program: %v\n", err)
	} else {
		fmt.Println("Success!")
	}

	fmt.Println()
	fmt.Println("Final machine state at stop or error:")
	fmt.Printf("Program Counter: 0x%x\n", vm.ProgramCounter)
	fmt.Printf("Accumulator:     0x%x\n", vm.Accumulator)
	fmt.Printf("Running:         %t\n", vm.Running)
	fmt.Println("Memory Content (up to last non-zero address):")
	fmt.Println("Address | Data")
	lastUsedAddress := 0
	for i := 0; i < len(vm.Memory); i++ {
		if vm.Memory[i] != 0 {
			lastUsedAddress = i
		}
	}
	for i := 0; i < lastUsedAddress+1; i++ {
		fmt.Printf("0x%04x  | 0x%04x\n", i, vm.Memory[i])
	}
}
