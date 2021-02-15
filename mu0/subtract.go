package mu0

import (
	"errors"
	"regexp"
)

const subtractAssemblyOpcode = "SUB"
const subtractMachineCodeOpcode = 0b01
var subtractAssemblyRegex = regexp.MustCompile("^SUB (R[0-3]) (R[0-3]) (R[0-3])$")

func init() {
	assemblyOpcodeParsers[subtractAssemblyOpcode] = parseSubtractAssembly
	machineCodeOpcodeParsers[subtractMachineCodeOpcode] = parseSubtractMachineCode
}

func parseSubtractAssembly(assemblyLine string) (Instruction, error) {
	matches := subtractAssemblyRegex.FindStringSubmatch(assemblyLine)
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

	return &subtractInstruction{
		Destination: destinationRegister,
		OperandA:    operandARegister,
		OperandB:    operandBRegister,
	}, nil
}

func parseSubtractMachineCode(instruction uint8) (Instruction, error) {
	opcode := (instruction >> 6) & 0x3
	if opcode != subtractMachineCodeOpcode {
		return nil, errors.New("incorrect opcode for this parser")
	}

	destination := (instruction >> 4) & 0x3
	operandA := (instruction >> 2) & 0x3
	operandB := instruction & 0x3

	return &subtractInstruction{
		Destination: Register(destination),
		OperandA:    Register(operandA),
		OperandB:    Register(operandB),
	}, nil
}

// Subtracts OperandB from OperandA and stores in Destination
type subtractInstruction struct {
	Destination Register
	OperandA Register
	OperandB Register
}

func (instruction *subtractInstruction) Assemble() (result uint8) {
	result |= subtractMachineCodeOpcode & 0x3
	result <<= 2
	result |= uint8(instruction.Destination & 0x3)
	result <<= 2
	result |= uint8(instruction.OperandA & 0x3)
	result <<= 2
	result |= uint8(instruction.OperandB & 0x3)
	return
}

func (instruction *subtractInstruction) Disassemble() string {
	return subtractAssemblyOpcode + " " + registerToString(instruction.Destination) + " " + registerToString(instruction.OperandA) + " " + registerToString(instruction.OperandB)
}
