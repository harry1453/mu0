package mu0

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const loadAssemblyOpcode = "LDI"
const loadMachineCodeOpcode = 0b10
var loadAssemblyRegex = regexp.MustCompile("^LDI (R[0-3]) (.+)$")

func init() {
	assemblyOpcodeParsers[loadAssemblyOpcode] = parseLoadAssembly
	machineCodeOpcodeParsers[loadMachineCodeOpcode] = parseLoadMachineCode
}

func parseLoadAssembly(assemblyLine string) (Instruction, error) {
	matches := loadAssemblyRegex.FindStringSubmatch(assemblyLine)
	if len(matches) < 3 {
		return nil, errors.New("parse error")
	}

	destinationRegister, err := stringToRegister(matches[1])
	if err != nil {
		return nil, err
	}

	constant, err := strconv.ParseUint(matches[2], 0, 64)
	if err != nil {
		return nil, err
	}
	if constant >= 8 {
		return nil, fmt.Errorf("constant %d too big, max 7", constant)
	}

	return &loadInstruction{
		Destination: destinationRegister,
		Constant: uint8(constant),
	}, nil
}

func parseLoadMachineCode(instruction uint8) (Instruction, error) {
	opcode := (instruction >> 6) & 0x3
	if opcode != addMachineCodeOpcode {
		return nil, errors.New("incorrect opcode for this parser")
	}

	destination := (instruction >> 4) & 0x3
	constant := instruction & 0xF

	return &loadInstruction{
		Destination: Register(destination),
		Constant: constant,
	}, nil
}

// Stores Constant in Destination
type loadInstruction struct {
	Destination Register
	// Only 4 LSB used
	Constant uint8
}

func (instruction *loadInstruction) Assemble() (result uint8) {
	result |= loadMachineCodeOpcode & 0x3
	result <<= 2
	result |= uint8(instruction.Destination & 0x3)
	result <<= 4
	result |= instruction.Constant & 0xF
	return
}

func (instruction *loadInstruction) Disassemble() string {
	return loadAssemblyOpcode + " " + registerToString(instruction.Destination) + " " + strconv.Itoa(int(instruction.Constant))
}
