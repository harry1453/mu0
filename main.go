package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"mu0-assembler/mu0"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Filename to assemble: ")
		filename, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Print("Simulate? (y/N): ")
		simulate, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		simulate = strings.ToLower(strings.TrimSpace(simulate))
		if err := assembleAndSimulate(strings.TrimSpace(filename), simulate == "y"); err != nil {
			fmt.Printf("Error running: %#v\n", err)
		}
	}
}

func assembleAndSimulate(filename string, simulate bool) error {
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	machineCode, err := mu0.Assemble(string(fileData))
	if err != nil {
		return err
	}

	fmt.Println("Machine code:")
	fmt.Println(hex.EncodeToString(mu0.ByteArrayFromMachineCode(machineCode)))

	fmt.Println()
	fmt.Println("Disassembled:")
	for _, instruction := range machineCode {
		parsedInstruction, err := mu0.ParseMachineInstruction(instruction)
		if err != nil {
			fmt.Printf("%04x\n", instruction)
		} else {
			fmt.Println(parsedInstruction.Disassemble())
		}
	}
	fmt.Println()

	if !simulate {
		return nil
	}

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
	fmt.Println()

	return nil
}
