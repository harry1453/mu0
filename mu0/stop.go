package mu0

const (
	stopAssemblyOpcode = "STP"
	stopOpcode         = 0b0111
)

func init() {
	assemblyOpcodeToMachineOpcode[stopAssemblyOpcode] = stopOpcode
	opcodeWithoutArgumentParsers[stopOpcode] = createStopInstruction
}

func createStopInstruction() Instruction {
	return stopInstructionSingleton
}

// Stops the program.
type stopInstruction struct {
}

var stopInstructionSingleton = new(stopInstruction)

func (instruction *stopInstruction) Assemble() (result uint16) {
	return assembleInstructionNoArgument(stopOpcode)
}

func (instruction *stopInstruction) Disassemble() string {
	return stopAssemblyOpcode
}
