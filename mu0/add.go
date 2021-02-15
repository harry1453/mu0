package mu0

import (
	"errors"
	"regexp"
)

const addAssemblyOpcode = "ADD"
const addMachineCodeOpcode = 0b00
var addAssemblyRegex = regexp.MustCompile("^ADD (R[0-3]) (R[0-3]) (R[0-3])$")

func init() {
	assemblyOpcodeParsers[addAssemblyOpcode] = parseAddAssembly
	machineCodeOpcodeParsers[addMachineCodeOpcode] = parseAddMachineCode
}

func parseAddAssembly(assemblyLine string) (Instruction, error) {
	matches := addAssemblyRegex.FindStringSubmatch(assemblyLine)
	if len(matches) < 4 {
		return nil, errors.New("parse error")
	}

	destinationRegister, err := stringToRegister(matches[1])
	if err != nil {
		return nil, err
	}

	operandARegister, err := stringToRegister(matches[2])
	if err != nil {
		return nil, err
	}

	operandBRegister, err := stringToRegister(matches[3])
	if err != nil {
		return nil, err
	}

	return &addInstruction{
		Destination: destinationRegister,
		OperandA:    operandARegister,
		OperandB:    operandBRegister,
	}, nil
}

func parseAddMachineCode(instruction uint8) (Instruction, error) {
	opcode := (instruction >> 6) & 0x3
	if opcode != addMachineCodeOpcode {
		return nil, errors.New("incorrect opcode for this parser")
	}

	destination := (instruction >> 4) & 0x3
	operandA := (instruction >> 2) & 0x3
	operandB := instruction & 0x3

	return &addInstruction{
		Destination: Register(destination),
		OperandA:    Register(operandA),
		OperandB:    Register(operandB),
	}, nil
}

// Adds OperandA and OperandB and stores in Destination
type addInstruction struct {
	Destination Register
	OperandA Register
	OperandB Register
}

func (instruction *addInstruction) Assemble() (result uint8) {
	result |= addMachineCodeOpcode & 0x3
	result <<= 2
	result |= uint8(instruction.Destination & 0x3)
	result <<= 2
	result |= uint8(instruction.OperandA & 0x3)
	result <<= 2
	result |= uint8(instruction.OperandB & 0x3)
	return
}

func (instruction *addInstruction) Disassemble() string {
	return addAssemblyOpcode + " " + registerToString(instruction.Destination) + " " + registerToString(instruction.OperandA) + " " + registerToString(instruction.OperandB)
}

