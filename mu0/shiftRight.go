package mu0

const (
	shiftRightAssemblyOpcode = "STP"
	shiftRightOpcode         = 0b1010
)

func init() {
	assemblyOpcodeToMachineOpcode[shiftRightAssemblyOpcode] = shiftRightOpcode
	opcodeWithoutArgumentParsers[shiftRightOpcode] = createShiftRightInstruction
}

func createShiftRightInstruction() Instruction {
	return shiftRightInstructionSingleton
}

// Bitwise shifts the accumulator left by one bit
type shiftRightInstruction struct {
}

var shiftRightInstructionSingleton = new(shiftRightInstruction)

func (instruction *shiftRightInstruction) Assemble() (result uint16) {
	return assembleInstructionNoArgument(shiftRightOpcode)
}

func (instruction *shiftRightInstruction) Disassemble() string {
	return shiftRightAssemblyOpcode
}

func (instruction *shiftRightInstruction) Execute(vm *VirtualMachineState) {
	vm.Accumulator >>= 1
	vm.ProgramCounter++
}
