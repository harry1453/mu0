package mu0

const (
	addAssemblyOpcode = "ADD"
	addOpcode         = 0b0010
)

func init() {
	assemblyOpcodeToMachineOpcode[addAssemblyOpcode] = addOpcode
	opcodeWithArgumentParsers[addOpcode] = createAddInstruction
}

func createAddInstruction(argument uint16) Instruction {
	return &addInstruction{
		memoryAddress: argument,
	}
}

// Adds the data at memoryAddress to the accumulator
type addInstruction struct {
	memoryAddress uint16
}

func (instruction *addInstruction) Assemble() (result uint16) {
	return assembleInstruction(addOpcode, instruction.memoryAddress)
}

func (instruction *addInstruction) Disassemble() string {
	return disassembleInstruction(addAssemblyOpcode, instruction.memoryAddress)
}
