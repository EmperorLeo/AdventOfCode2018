package main

import "fmt"

func day19() {
	cpu := Processor{}
	lines := readFile("../input/day19.txt")
	var pc int
	fmt.Sscanf(lines[0], "#ip %d", &pc)
	instructions := make([]Instruction, len(lines)-1)
	for i, line := range lines[1:] {
		var op string
		var a, b, c int
		fmt.Sscanf(line, "%s %d %d %d", &op, &a, &b, &c)
		instructions[i] = Instruction{op, a, b, c}
		fmt.Printf("%s %d %d %d\n", instructions[i].op, instructions[i].a, instructions[i].b, instructions[i].c)
	}
	fmt.Printf("IP = %d\n", pc)
	cpu.Init(6, pc, instructions)
	for cpu.ExecuteInstruction(true) {
	}

	// Part 1
	fmt.Printf("Value in register 0 == %d\n", cpu.registers[0].value)

	// Part 2
	for _, register := range cpu.registers {
		register.value = 0
	}
	cpu.registers[0].value = 1
	for cpu.ExecuteInstruction(true) {
	}
	fmt.Printf("Value in register 0 == %d\n", cpu.registers[0].value)
}
