package mu0

const (
	jumpIfNegativeAssemblyOpcode = "JMI"
	jumpIfNegativeOpcode         = 0b0101
)

func init() {
	assemblyOpcodeToMachineOpcode[jumpIfNegativeAssemblyOpcode] = jumpIfNegativeOpcode
	opcodeWithArgumentParsers[jumpIfNegativeOpcode] = createJumpIfNegativeInstruction
}

func createJumpIfNegativeInstruction(argument uint16) Instruction {
	return &jumpIfNegativeInstruction{
		memoryAddress: argument,
	}
}

// Sets the program counter to the data at memoryAddress if the accumulator is less than 0 (using 2's complement)
type jumpIfNegativeInstruction struct {
	memoryAddress uint16
}

func (instruction *jumpIfNegativeInstruction) Assemble() (result uint16) {
	return assembleInstruction(jumpIfNegativeOpcode, instruction.memoryAddress)
}

func (instruction *jumpIfNegativeInstruction) Disassemble() string {
	return disassembleInstruction(jumpIfNegativeAssemblyOpcode, instruction.memoryAddress)
}
