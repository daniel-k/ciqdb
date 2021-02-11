package ciqdb

import (
	"fmt"
	"strings"
)

var opcodeSize = map[string]int{
	"NOP":       1, // 0
	"INCSP":     2, // 1
	"POPV":      1, // 2
	"ADD":       1, // 3
	"SUB":       1, // 4
	"MUL":       1, // 5
	"DIV":       1, // 6
	"AND":       1, // 7
	"OR":        1, // 8
	"MOD":       1, // 9
	"SHL":       2, // 10
	"SHR":       2, // 11
	"XOR":       1, // 12
	"GETV":      1, // 13
	"PUTV":      1, // 14
	"INVOKE":    2, // 15
	"AGETV":     1, // 16
	"APUTV":     1, // 17
	"LGETV":     2, // 18
	"LPUTV":     2, // 19
	"NEWA":      1, // 20
	"NEWC":      1, // 21
	"RETURN":    1, // 22
	"RET":       1, // 23
	"NEWS":      5, // 24
	"GOTO":      3, // 25
	"EQ":        1, // 26
	"LT":        1, // 27
	"LTE":       1, // 28
	"GT":        1, // 29
	"GTE":       1, // 30
	"NE":        1, // 31
	"ISNULL":    1, // 32
	"ISA":       1, // 33
	"CANHAZPLZ": 1, // 34
	"JSR":       3, // 35
	"TS":        1, // 36
	"IPUSH":     5, // 37
	"FPUSH":     5, // 38
	"SPUSH":     5, // 39
	"BT":        3, // 40
	"BF":        3, // 41
	"FRPUSH":    1, // 42
	"BPUSH":     2, // 43
	"NPUSH":     1, // 44
	"INV":       1, // 45
	"DUP":       2, // 46
	"NEWD":      1, // 47
	"GETM":      1, // 48
	"LPUSH":     9, // 49
	"DPUSH":     9, // 50
	"THROW":     1, // 51
	"CPUSH":     5, // 52
	"ARGC":      2, // 53
	"NEWBA":     1, // 54
}

var opcodeToString = map[int]string{
	0:  "NOP",
	1:  "INCSP",
	2:  "POPV",
	3:  "ADD",
	4:  "SUB",
	5:  "MUL",
	6:  "DIV",
	7:  "AND",
	8:  "OR",
	9:  "MOD",
	10: "SHL",
	11: "SHR",
	12: "XOR",
	13: "GETV",
	14: "PUTV",
	15: "INVOKE",
	16: "AGETV",
	17: "APUTV",
	18: "LGETV",
	19: "LPUTV",
	20: "NEWA",
	21: "NEWC",
	22: "RETURN",
	23: "RET",
	24: "NEWS",
	25: "GOTO",
	26: "EQ",
	27: "LT",
	28: "LTE",
	29: "GT",
	30: "GTE",
	31: "NE",
	32: "ISNULL",
	33: "ISA",
	34: "CANHAZPLZ",
	35: "JSR",
	36: "TS",
	37: "IPUSH",
	38: "FPUSH",
	39: "SPUSH",
	40: "BT",
	41: "BF",
	42: "FRPUSH",
	43: "BPUSH",
	44: "NPUSH",
	45: "INV",
	46: "DUP",
	47: "NEWD",
	48: "GETM",
	49: "LPUSH",
	50: "DPUSH",
	51: "THROW",
	52: "CPUSH",
	53: "ARGC",
	54: "NEWBA",
}

type Instruction struct {
	offset int
	opcode string
	arg    []byte
}

func (i *Instruction) String() string {
	if len(i.arg) > 0 {
		// TODO: read proper integer, maybe even resolve symbols
		return fmt.Sprintf("%x\t%v 0x%x", i.offset, i.opcode, i.arg)
	}

	return fmt.Sprintf("%x\t%v", i.offset, i.opcode)
}

type Code struct {
	PRGSection
	instructions []Instruction
}

func (d *Code) String() string {
	var instructions []string

	for _, ins := range d.instructions {
		instructions = append(instructions, ins.String())
	}

	return d.PRGSection.String() + "\n" +
		"    code: \n  " + strings.Join(instructions, "\n  ")
}

func parseCode(p *PRG, t SecType, length int, data []byte) *Code {
	code := Code{
		PRGSection: PRGSection{
			Type:   t,
			length: length,
		},
	}

	offset := 0
	for offset < length {
		opcode := opcodeToString[int(data[offset])]
		insLength := opcodeSize[opcode]

		code.instructions = append(code.instructions, Instruction{
			offset: offset,
			opcode: opcode,
			arg:    data[offset+1 : offset+insLength],
		})

		offset += insLength
	}

	return &code
}
