package mu0

const (
	shiftLeftAssemblyOpcode = "LSL"
	shiftLeftOpcode         = 0b1001
)

func init() {
	assemblyOpcodeToMachineOpcode[shiftLeftAssemblyOpcode] = shiftLeftOpcode
	opcodeWithoutArgumentParsers[shiftLeftOpcode] = createShiftLeftInstruction
}

func createShiftLeftInstruction() Instruction {
	return shiftLeftInstructionSingleton
}

// Bitwise shifts the accumulator left by one bit
type shiftLeftInstruction struct {
}

var shiftLeftInstructionSingleton = new(shiftLeftInstruction)

func (instruction *shiftLeftInstruction) Assemble() (result uint16) {
	return assembleInstructionNoArgument(shiftLeftOpcode)
}

func (instruction *shiftLeftInstruction) Disassemble() string {
	return shiftLeftAssemblyOpcode
}
