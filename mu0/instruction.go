package mu0

import (
	"fmt"
	"regexp"
	"strings"
)

type Instruction interface {
	// Assemble into machine code
	Assemble() uint8
	// Disassemble into assembly
	Disassemble() string
}

var assemblyOpcodeRegex = regexp.MustCompile("^([A-Z]{3}).+")
var assemblyOpcodeParsers = make(map[string]func(string)(Instruction, error))
var machineCodeOpcodeParsers = make(map[uint8]func(uint8)(Instruction, error))

func ParseAssembly(assembly string) ([]Instruction, error) {
	lines := strings.Split(assembly, "\n")
	var instructions []Instruction
	for i, line := range lines {
		line = strings.TrimSpace(line)
		// Ignore blank lines and comments
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}

		matches := assemblyOpcodeRegex.FindStringSubmatch(line)
		if len(matches) < 2 {
			return nil, fmt.Errorf("parse error on line %d", i+1)
		}
		opcode := matches[1]

		opcodeParser, ok := assemblyOpcodeParsers[opcode]
		if !ok {
			return nil, fmt.Errorf("could not find opcode %s on line %d", opcode, i+1)
		}

		instruction, err := opcodeParser(line)
		if err != nil {
			return nil, fmt.Errorf("instruction parse error on line %d: %e", i+1, err)
		}

		instructions = append(instructions, instruction)
	}

	return instructions, nil
}

func ParseMachineCode(machineCode []byte) ([]Instruction, error) {
	var instructions []Instruction
	for i, instruction := range machineCode {
		opcode := (instruction >> 6) & 0x3
		opcodeParser, ok := machineCodeOpcodeParsers[opcode]
		if !ok {
			return nil, fmt.Errorf("error processing instruction %d: did not recognize opcode 0x%x", i, opcode)
		}

		parsedInstruction, err := opcodeParser(instruction)
		if err != nil {
			return nil, fmt.Errorf("instruction parse error for instruction %d: %e", i, err)
		}

		instructions = append(instructions, parsedInstruction)
	}

	return instructions, nil
}
