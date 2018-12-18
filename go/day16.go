package main

import (
	"fmt"
)

type opcode func(a int, b int, c int)

type register struct {
	value int
}

type processor struct {
	registers map[int]*register
	opcodes   []opcode
	opcodeMap map[int]int
}

func day16() {
	cpu := processor{}
	cpu.init()
	var behaveLikeThreeOrMore int
	lines := readFile("../input/day16.txt")
	impossibleOpcodeSet := make(map[int](map[int]bool), 16)
	for i := 0; i < 16; i++ {
		impossibleOpcodeSet[i] = make(map[int]bool, 16)
	}
	var startOfProgram int
	for i := 0; i < len(lines); i += 4 {
		if lines[i] == "" {
			startOfProgram = i + 2
			break
		}

		var b1, b2, b3, b4, i1, i2, i3, i4, a1, a2, a3, a4 int
		fmt.Sscanf(lines[i], "Before: [%d, %d, %d, %d]", &b1, &b2, &b3, &b4)
		fmt.Sscanf(lines[i+1], "%d %d %d %d", &i1, &i2, &i3, &i4)
		fmt.Sscanf(lines[i+2], "After: [%d, %d, %d, %d]", &a1, &a2, &a3, &a4)
		if cpu.computeViable([]int{b1, b2, b3, b4}, []int{i1, i2, i3, i4}, []int{a1, a2, a3, a4}, impossibleOpcodeSet) >= 3 {
			behaveLikeThreeOrMore++
		}
	}

	// Part 1
	fmt.Printf("%d samples behave like 3 or more opcodes.", behaveLikeThreeOrMore)
	fmt.Println()

	// Part 2
	// initialize registers to 0
	cpu.resolveOpcodeMapping(impossibleOpcodeSet)
	for _, r := range cpu.registers {
		r.value = 0
	}
	for i := startOfProgram; i < len(lines); i++ {
		var i1, i2, i3, i4 int
		fmt.Sscanf(lines[i], "%d %d %d %d", &i1, &i2, &i3, &i4)
		cpu.opcodes[cpu.opcodeMap[i1]](i2, i3, i4)
	}
	fmt.Print("Values in registers are:\n")
	fmt.Printf("0: %d\n", cpu.registers[0].value)
	fmt.Printf("1: %d\n", cpu.registers[1].value)
	fmt.Printf("2: %d\n", cpu.registers[2].value)
	fmt.Printf("3: %d\n", cpu.registers[3].value)
}

func (p *processor) init() {
	p.registers = make(map[int]*register, 4)
	p.registers[0] = &register{}
	p.registers[1] = &register{}
	p.registers[2] = &register{}
	p.registers[3] = &register{}
	p.opcodes = make([]opcode, 16)
	p.opcodes[0] = func(a int, b int, c int) { // addr
		p.registers[c].value = p.registers[a].value + p.registers[b].value
	}
	p.opcodes[1] = func(a int, b int, c int) { //addi
		p.registers[c].value = p.registers[a].value + b
	}
	p.opcodes[2] = func(a int, b int, c int) { // mulr
		p.registers[c].value = p.registers[a].value * p.registers[b].value
	}
	p.opcodes[3] = func(a int, b int, c int) { // muli
		p.registers[c].value = p.registers[a].value * b
	}
	p.opcodes[4] = func(a int, b int, c int) { // banr
		p.registers[c].value = p.registers[a].value & p.registers[b].value
	}
	p.opcodes[5] = func(a int, b int, c int) { // bani
		p.registers[c].value = p.registers[a].value & b
	}
	p.opcodes[6] = func(a int, b int, c int) { // borr
		p.registers[c].value = p.registers[a].value | p.registers[b].value
	}
	p.opcodes[7] = func(a int, b int, c int) { // bori
		p.registers[c].value = p.registers[a].value | b
	}
	p.opcodes[8] = func(a int, b int, c int) { // setr
		p.registers[c].value = p.registers[a].value
	}
	p.opcodes[9] = func(a int, b int, c int) { // seti
		p.registers[c].value = a
	}
	p.opcodes[10] = func(a int, b int, c int) { // gtir
		if a > p.registers[b].value {
			p.registers[c].value = 1
		} else {
			p.registers[c].value = 0
		}
	}
	p.opcodes[11] = func(a int, b int, c int) { // gtri
		if p.registers[a].value > b {
			p.registers[c].value = 1
		} else {
			p.registers[c].value = 0
		}
	}
	p.opcodes[12] = func(a int, b int, c int) { // gtrr
		if p.registers[a].value > p.registers[b].value {
			p.registers[c].value = 1
		} else {
			p.registers[c].value = 0
		}
	}
	p.opcodes[13] = func(a int, b int, c int) { // eqir
		if a == p.registers[b].value {
			p.registers[c].value = 1
		} else {
			p.registers[c].value = 0
		}
	}
	p.opcodes[14] = func(a int, b int, c int) { // eqri
		if p.registers[a].value == b {
			p.registers[c].value = 1
		} else {
			p.registers[c].value = 0
		}
	}
	p.opcodes[15] = func(a int, b int, c int) { // eqrr
		if p.registers[a].value == p.registers[b].value {
			p.registers[c].value = 1
		} else {
			p.registers[c].value = 0
		}
	}
}

func (p *processor) computeViable(before []int, instruction []int, after []int, possibleOpcodeSet map[int](map[int]bool)) int {
	var viableCount int
	for i, op := range p.opcodes {
		p.registers[0].value = before[0]
		p.registers[1].value = before[1]
		p.registers[2].value = before[2]
		p.registers[3].value = before[3]
		op(instruction[1], instruction[2], instruction[3])
		if p.registers[0].value == after[0] && p.registers[1].value == after[1] && p.registers[2].value == after[2] && p.registers[3].value == after[3] {
			viableCount++
		} else {
			// Using true to mean that this opcode pairing is not possible - so i don't have to initalize everything to true!
			possibleOpcodeSet[instruction[0]][i] = true
		}
	}
	return viableCount
}

func (p *processor) resolveOpcodeMapping(impossibleOpcodeSet map[int](map[int]bool)) {
	p.opcodeMap = make(map[int]int, 16)
	unmatchedOpcodes := 16
	for unmatchedOpcodes > 0 {
		for k1, v1 := range impossibleOpcodeSet {
			fmt.Printf("Length of not possible set for opcode %d is %d.\n", k1, len(v1))
			if len(v1) == 15 {
				missingOpcode := getMissingOpcode(v1)
				p.opcodeMap[k1] = missingOpcode
				v1[missingOpcode] = false
				for k2, v2 := range impossibleOpcodeSet {
					if k2 != k1 {
						v2[missingOpcode] = true
					}
				}
				unmatchedOpcodes--
			}
		}
	}

	for k, v := range p.opcodeMap {
		fmt.Printf("Opcode %d has index %d in opcodes array.", k, v)
		fmt.Println()
	}
}

func getMissingOpcode(m map[int]bool) int {
	total := 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 + 11 + 12 + 13 + 14 + 15
	for k, v := range m {
		if v {
			total -= k
		}
	}
	return total
}
