package mu0

import "testing"

func TestJumps(t *testing.T) {
	const asm = `
LDI 0x123 // check Acc = 0x0123 after LDI
	LDA 0x000 // check Acc = 0x8123 after LDA
STA 0x102 // use STA, can’t check yet
LSR check // Acc = 0x4091 after LSR
ADD 0x002 // check Acc = 0x4091+0x1102 = 0x5193 after ADD
SUB 0x002 // check Acc = 0x4091 after SUB
LDA 0x102 // check Acc = 0x8123 after LDA - this checks previous STA
LSL check // Acc = 0x0246 after LSL
LDI 0x000
JEQ 0x00b
JMP 0x012 // executed if JEQ does not work
JMI 0x012 // executed if JMI does not work
LDI 0x001
JEQ 0x012 // executed if JEQ does not work
LDA 0x013 // load 0x8000
JMI 0x011
JMP 0x012 // executed if JMI does not work
STP 0x11 // Stop here if Jump tests PASSED
STP 0x12 // Stop here if Jump tests FAILED
0x8000 // constant to test JMI
`
	machineCode, err := Assemble(asm)
	if err != nil {
		t.Fatal(err)
	}

	vmState, err := RunProgram(machineCode, true)
	if err != nil {
		t.Fatal(err)
	}

	if vmState.ProgramCounter != 0x11 {
		t.Fail()
	}
}
