package mu0

const (
	jumpAssemblyOpcode = "JMP"
	jumpOpcode         = 0b0100
)

func init() {
	assemblyOpcodeToMachineOpcode[jumpAssemblyOpcode] = jumpOpcode
	opcodeWithArgumentParsers[jumpOpcode] = createJumpInstruction
}

func createJumpInstruction(argument uint16) Instruction {
	return &jumpInstruction{
		memoryAddress: argument,
	}
}

// Sets the program counter to the data at memoryAddress
type jumpInstruction struct {
	memoryAddress uint16
}

func (instruction *jumpInstruction) Assemble() (result uint16) {
	return assembleInstruction(jumpOpcode, instruction.memoryAddress)
}

func (instruction *jumpInstruction) Disassemble() string {
	return disassembleInstruction(jumpAssemblyOpcode, instruction.memoryAddress)
}

func (instruction *jumpInstruction) Execute(vm *VirtualMachineState) {
	vm.ProgramCounter = instruction.memoryAddress
}
