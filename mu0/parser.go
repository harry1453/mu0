package mu0

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var assemblyInstructionRegex = regexp.MustCompile("^[\\t ]*([A-Z]{3})[\\t ]*(\\w+)[\\t ]*(?://.+)?$")

var assemblyOpcodeToMachineOpcode = make(map[string]uint8)
var opcodeWithArgumentParsers = make(map[uint8]func(argument uint16) Instruction)
var opcodeWithoutArgumentParsers = make(map[uint8]func() Instruction)

func ParseAssembly(assembly string) ([]Instruction, error) {
	lines := strings.Split(assembly, "\n")
	var instructions []Instruction
	for i, line := range lines {
		line = strings.TrimSpace(line)
		// Ignore blank lines and comments
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}

		matches := assemblyInstructionRegex.FindStringSubmatch(line)
		if len(matches) < 2 {
			return nil, fmt.Errorf("parse error on line %d", i+1)
		}

		opcodeString := matches[1]
		opcode, ok := assemblyOpcodeToMachineOpcode[opcodeString]
		if !ok {
			return nil, fmt.Errorf("could not find opcode %s on line %d", opcodeString, i+1)
		}

		var instruction Instruction
		opcodeParser, ok := opcodeWithArgumentParsers[opcode]
		if ok {
			argument, err := strconv.ParseUint(matches[2], 0, 64)
			if err != nil {
				return nil, err
			}
			if argument > 1<<12 {
				return nil, fmt.Errorf("argument %d too big, max %d", argument, 1<<12)
			}
			instruction = opcodeParser(uint16(argument))
		} else {
			opcodeParser, ok := opcodeWithoutArgumentParsers[opcode]
			if ok {
				instruction = opcodeParser()
			} else {
				return nil, fmt.Errorf("could not find opcode parser for opcode %s  (0x%x) on line %d", opcodeString, opcode, i+1)
			}
		}

		instructions = append(instructions, instruction)
	}

	return instructions, nil
}

func ParseMachineCode(machineCode []uint16) ([]Instruction, error) {
	var instructions []Instruction
	for i, instruction := range machineCode {
		opcode, argument := parseInstruction(instruction)
		opcodeParser, ok := opcodeWithArgumentParsers[opcode]
		if !ok {
			return nil, fmt.Errorf("error processing instruction %d: did not recognize opcode 0x%x", i, opcode)
		}

		instructions = append(instructions, opcodeParser(argument))
	}

	return instructions, nil
}
