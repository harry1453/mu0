package mu0

const (
	loadAssemblyOpcode = "LDA"
	loadOpcode         = 0b0000
)

func init() {
	assemblyOpcodeToMachineOpcode[loadAssemblyOpcode] = loadOpcode
	opcodeWithArgumentParsers[loadOpcode] = createLoadInstruction
}

func createLoadInstruction(argument uint16) Instruction {
	return &loadInstruction{
		memoryAddress: argument,
	}
}

// Loads the data at memoryAddress into the accumulator
type loadInstruction struct {
	memoryAddress uint16
}

func (instruction *loadInstruction) Assemble() (result uint16) {
	return assembleInstruction(loadOpcode, instruction.memoryAddress)
}

func (instruction *loadInstruction) Disassemble() string {
	return disassembleInstruction(loadAssemblyOpcode, instruction.memoryAddress)
}

func (instruction *loadInstruction) Execute(vm *VirtualMachineState) {
	vm.Accumulator = vm.Memory[instruction.memoryAddress]
	vm.IncrementProgramCounter()
}
