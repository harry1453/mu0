package mu0

import (
	"fmt"
	"strconv"
	"strings"
)

type Register uint8
const (
	R0 Register = iota
	R1
	R2
	R3
)

func stringToRegister(register string) (Register, error) {
	register = strings.TrimPrefix(register, "R")
	registerNumber, err := strconv.Atoi(register)
	if err != nil {
		return R0, err
	}
	switch registerNumber {
	case 0:
		return R0, nil
	case 1:
		return R1, nil
	case 2:
		return R2, nil
	case 3:
		return R3, nil
	}
	return R0, fmt.Errorf("register %d does not exist", registerNumber)
}

func registerToString(register Register) string {
	switch register {
	case R0:
		return "R0"
	case R1:
		return "R1"
	case R2:
		return "R2"
	case R3:
		return "R3"
	}
	return "R?"
}
