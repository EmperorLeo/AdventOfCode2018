package main

import (
	"fmt"
)

func day21() {
	// Copied and pasted most of this code from day 19
	cpu := Processor{}
	// For part 1, change this to read from ../input/day21-part1.txt
	lines := readFile("../input/day21.txt")
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
	// For part 2, assume a cycle if there is going to be a value that makes the program run the longest
	// Register 1 needs to match (from my bad decompiled code that doesnt work!)
	valueMap := map[int]bool{}
	var lastValue int
	var ioLength int
	for cpu.ExecuteInstruction(false) {
		// New IO!!! Wooooooo
		if len(cpu.ioDevice) != ioLength {
			curVal := cpu.ioDevice[len(cpu.ioDevice)-1] // Get the latest IO value
			exists := valueMap[curVal]
			if exists {
				break
			}
			valueMap[curVal] = true
			lastValue = curVal
			ioLength = len(cpu.ioDevice)
		}
	}

	fmt.Printf("Last value that wasn't repeated is %d.\n", lastValue)
}
