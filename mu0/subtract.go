package mu0

const (
	subtractAssemblyOpcode = "SUB"
	subtractOpcode         = 0b0011
)

func init() {
	assemblyOpcodeToMachineOpcode[subtractAssemblyOpcode] = subtractOpcode
	opcodeWithArgumentParsers[subtractOpcode] = createSubtractInstruction
}

func createSubtractInstruction(argument uint16) Instruction {
	return &subtractInstruction{
		memoryAddress: argument,
	}
}

// Subtracts the data at memoryAddress into the accumulator
type subtractInstruction struct {
	memoryAddress uint16
}

func (instruction *subtractInstruction) Assemble() (result uint16) {
	return assembleInstruction(subtractOpcode, instruction.memoryAddress)
}

func (instruction *subtractInstruction) Disassemble() string {
	return disassembleInstruction(subtractAssemblyOpcode, instruction.memoryAddress)
}
