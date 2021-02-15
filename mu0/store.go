package mu0

const (
	storeAssemblyOpcode = "STA"
	storeOpcode         = 0b0001
)

func init() {
	assemblyOpcodeToMachineOpcode[storeAssemblyOpcode] = storeOpcode
	opcodeWithArgumentParsers[storeOpcode] = createStoreInstruction
}

func createStoreInstruction(argument uint16) Instruction {
	return &storeInstruction{
		memoryAddress: argument,
	}
}

// Store the accumulator into the memory at memoryAddress
type storeInstruction struct {
	memoryAddress uint16
}

func (instruction *storeInstruction) Assemble() (result uint16) {
	return assembleInstruction(storeOpcode, instruction.memoryAddress)
}

func (instruction *storeInstruction) Disassemble() string {
	return disassembleInstruction(storeAssemblyOpcode, instruction.memoryAddress)
}

func (instruction *storeInstruction) Execute(vm *VirtualMachineState) {
	vm.Memory[instruction.memoryAddress] = vm.Accumulator
	vm.ProgramCounter++
}
