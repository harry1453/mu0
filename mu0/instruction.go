package mu0

import "strconv"

type Instruction interface {
	// Assemble into machine code
	Assemble() uint16
	// Disassemble into assembly
	Disassemble() string
	// Execute this instruction on the virtual machine
	Execute(vm *VirtualMachineState)
}

// Converts an opcode and argument to bytecode instruction
func assembleInstruction(opcode uint8, argument uint16) (instruction uint16) {
	instruction |= uint16(opcode & 0xF)
	instruction <<= 12
	instruction |= argument & 0xFFF
	return
}

// Converts an opcode to bytecode instruction
func assembleInstructionNoArgument(opcode uint8) uint16 {
	return assembleInstruction(opcode, 0)
}

// Converts a bytecode instruction to its opcode and argument
func parseInstruction(instruction uint16) (opcode uint8, argument uint16) {
	opcode = uint8((instruction >> 12) & 0xF)
	argument = instruction & 0xFFF
	return
}

// Converts an assembly opcode and argument to assembly instruction
func disassembleInstruction(assemblyOpcode string, argument uint16) string {
	return assemblyOpcode + " 0x" + strconv.FormatUint(uint64(argument), 16)
}
