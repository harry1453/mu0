package mu0

import (
	"errors"
	"fmt"
)

const MemorySize = 1 << 12

type VirtualMachineState struct {
	Running        bool
	ProgramCounter uint16
	Accumulator    uint16
	Memory         [MemorySize]uint16
}

func RunProgram(programMachineCode []uint16, debug bool) (VirtualMachineState, error) {
	vm := VirtualMachineState{
		Running: true,
	}

	// Copy program into memory
	for i := 0; i < len(programMachineCode); i++ {
		vm.Memory[i] = programMachineCode[i]
	}

	// Run program
	for vm.Running {
		if vm.ProgramCounter >= MemorySize {
			return vm, errors.New("program counter went past the end of memory")
		}
		instruction, err := ParseMachineInstruction(vm.Memory[vm.ProgramCounter])
		if err != nil {
			return vm, err
		}
		instruction.Execute(&vm)
		if debug {
			fmt.Printf("Executing instruction: %#v, new program counter 0x%04x, new accumulator 0x%04x\n", instruction, vm.ProgramCounter, vm.Accumulator)
		}
	}

	return vm, nil
}
