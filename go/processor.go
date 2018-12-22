package main

import "fmt"

// Opcode defines a function performed on 3 registers
type Opcode func(a int, b int, c int)

// Register just holds a value in a processor
type Register struct {
	value int
}

// Processor - its a processor used for days 16 and 19
type Processor struct {
	registers     map[int]*Register
	opcodes       []Opcode
	instructions  []Instruction
	pc            int
	opcodeMap     map[int]int
	opcodeNameMap map[string]int
}

// Instruction represents a 4 digit instruction for the Processor to process
type Instruction struct {
	op      string
	a, b, c int
}

// Init the processor with registers and opcodes and stuff
func (p *Processor) Init(numRegisters int, pc int, instructions []Instruction) {
	p.registers = make(map[int]*Register, numRegisters)
	p.pc = pc
	for i := 0; i < numRegisters; i++ {
		p.registers[i] = &Register{}
	}
	p.opcodes = make([]Opcode, 16)
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
	p.opcodeNameMap = map[string]int{
		"addr": 0,
		"addi": 1,
		"mulr": 2,
		"muli": 3,
		"banr": 4,
		"bani": 5,
		"borr": 6,
		"bori": 7,
		"setr": 8,
		"seti": 9,
		"gtir": 10,
		"gtri": 11,
		"gtrr": 12,
		"eqir": 13,
		"eqri": 14,
		"eqrr": 15}
	p.instructions = instructions
}

// ExecuteInstruction updates the registers and the program counter after executing instruction
func (p *Processor) ExecuteInstruction(debug bool) bool {
	if debug {
		fmt.Printf("Value in register 3 = %d\n", p.registers[p.pc].value)
		fmt.Printf("PROG START %d %d %d %d %d %d\n", p.registers[0].value, p.registers[1].value, p.registers[2].value, p.registers[3].value, p.registers[4].value, p.registers[5].value)
	}
	if p.registers[p.pc].value > len(p.instructions)-1 {
		return false
	}

	curInstruction := p.instructions[p.registers[p.pc].value]
	opcode := p.opcodeNameMap[curInstruction.op]
	p.opcodes[opcode](curInstruction.a, curInstruction.b, curInstruction.c)
	p.registers[p.pc].value++
	if debug {
		fmt.Printf("PROG END %d %d %d %d %d %d\n", p.registers[0].value, p.registers[1].value, p.registers[2].value, p.registers[3].value, p.registers[4].value, p.registers[5].value)
	}
	return true
}
