package mu0

const (
	loadImmediateAssemblyOpcode = "LDI"
	loadImmediateOpcode         = 0b1000
)

func init() {
	assemblyOpcodeToMachineOpcode[loadImmediateAssemblyOpcode] = loadImmediateOpcode
	opcodeWithArgumentParsers[loadImmediateOpcode] = createLoadImmediateInstruction
}

func createLoadImmediateInstruction(argument uint16) Instruction {
	return &loadImmediateInstruction{
		constant: argument,
	}
}

// Loads the constant into the accumulator
type loadImmediateInstruction struct {
	constant uint16
}

func (instruction *loadImmediateInstruction) Assemble() (result uint16) {
	return assembleInstruction(loadImmediateOpcode, instruction.constant)
}

func (instruction *loadImmediateInstruction) Disassemble() string {
	return disassembleInstruction(loadImmediateAssemblyOpcode, instruction.constant)
}
