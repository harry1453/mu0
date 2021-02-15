package mu0

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var assemblyInstructionRegex = regexp.MustCompile("^[\\t ]*([A-Z]{3})[\\t ]*(\\w+)?[\\t ]*(?://.+)?$")
var dataInstructionRegex = regexp.MustCompile("^[\\t ]*(\\w+)?[\\t ]*(?://.+)?$")

var assemblyOpcodeToMachineOpcode = make(map[string]uint8)
var opcodeWithArgumentParsers = make(map[uint8]func(argument uint16) Instruction)
var opcodeWithoutArgumentParsers = make(map[uint8]func() Instruction)

func Assemble(assembly string) ([]uint16, error) {
	lines := strings.Split(assembly, "\n")
	var instructions []uint16
	for i, line := range lines {
		line = strings.TrimSpace(line)
		// Ignore blank lines and comments
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}

		matches := assemblyInstructionRegex.FindStringSubmatch(line)
		if len(matches) < 3 {
			// This line is static data, not an instruction
			matches := dataInstructionRegex.FindStringSubmatch(line)
			if len(matches) < 2 {
				return nil, fmt.Errorf("parse error on line %d", i+1)
			}
			staticData, err := strconv.ParseUint(matches[1], 0, 64)
			if err != nil {
				return nil, err
			}
			if staticData > 1<<16 {
				return nil, fmt.Errorf("static data 0x%04d too big, max %d on line %d", staticData, 1<<16, i+1)
			}
			instructions = append(instructions, uint16(staticData))
			continue
		}

		opcodeString := matches[1]
		opcode, ok := assemblyOpcodeToMachineOpcode[opcodeString]
		if !ok {
			return nil, fmt.Errorf("could not find opcode %s on line %d", opcodeString, i+1)
		}

		var instruction Instruction
		opcodeParser, ok := opcodeWithArgumentParsers[opcode]
		if ok {
			if matches[2] == "" {
				return nil, fmt.Errorf("argument missing on line %d", i+1)
			}
			argument, err := strconv.ParseUint(matches[2], 0, 64)
			if err != nil {
				return nil, err
			}
			if argument > 1<<12 {
				return nil, fmt.Errorf("argument 0x%04d too big, max %d on line %d", argument, 1<<12, i+1)
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

		instructions = append(instructions, instruction.Assemble())
	}

	return instructions, nil
}

func ParseMachineInstruction(instruction uint16) (Instruction, error) {
	opcode, argument := parseInstruction(instruction)
	opcodeParser, ok := opcodeWithArgumentParsers[opcode]
	if ok {
		return opcodeParser(argument), nil
	} else {
		opcodeParser, ok := opcodeWithoutArgumentParsers[opcode]
		if ok {
			return opcodeParser(), nil
		} else {
			return nil, fmt.Errorf("unrecognized opcode 0x%x", opcode)
		}
	}
}

func Disassemble(machineCode []uint16) ([]Instruction, error) {
	var instructions []Instruction
	for i, instruction := range machineCode {
		parsedInstruction, err := ParseMachineInstruction(instruction)
		if err != nil {
			return nil, fmt.Errorf("error processing instruction %d: %v", i, err)
		}
		instructions = append(instructions, parsedInstruction)
	}

	return instructions, nil
}
