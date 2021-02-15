package mu0

const (
	jumpIfEqualAssemblyOpcode = "JEQ"
	jumpIfEqualOpcode         = 0b0110
)

func init() {
	assemblyOpcodeToMachineOpcode[jumpIfEqualAssemblyOpcode] = jumpIfEqualOpcode
	opcodeWithArgumentParsers[jumpIfEqualOpcode] = createJumpIfEqualInstruction
}

func createJumpIfEqualInstruction(argument uint16) Instruction {
	return &jumpIfEqualInstruction{
		memoryAddress: argument,
	}
}

// Sets the program counter to the data at memoryAddress if the accumulator is 0
type jumpIfEqualInstruction struct {
	memoryAddress uint16
}

func (instruction *jumpIfEqualInstruction) Assemble() (result uint16) {
	return assembleInstruction(jumpIfEqualOpcode, instruction.memoryAddress)
}

func (instruction *jumpIfEqualInstruction) Disassemble() string {
	return disassembleInstruction(jumpIfEqualAssemblyOpcode, instruction.memoryAddress)
}

func (instruction *jumpIfEqualInstruction) Execute(vm *VirtualMachineState) {
	if vm.Accumulator == 0 {
		vm.ProgramCounter = instruction.memoryAddress
	} else {
		vm.ProgramCounter++
	}
}
