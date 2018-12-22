package main

// Initial value is register 0
// valueOne is register 5
// valueTwo is register 1
// valueThree is register 4
// valueFour is register 2

func elfProg() {
	initialValue := 1
	valueOne := 2
	// 19 is current prog counter
	valueOne = valueOne * valueOne * 19 * 11
	// valueOne = valueOne * 19
	// valueOne = valueOne * 11

	// twenty two is current prog counter
	valueTwo := 6*22 + 13
	valueOne += valueTwo

	// Here, an instruction is skipped??? Prog counter incremented by 1. I'll write the instruction down in code.
	// seti 0 -> 3 // Initial value gets set to the program counter??? While loop?

	// 27 is the current prog counter
	valueTwo = 27
	// 28 is the current prog counter
	valueTwo = 28 * valueTwo
	// 29 is the current prog counter
	valueTwo = 29 + valueTwo
	// 30 is the current prog counter
	valueTwo = 30 * valueTwo
	valueTwo = 14 * valueTwo
	// 32 is the current prog counter
	valueTwo = 32 * valueTwo
	valueOne += valueTwo

	initialValue = 0
	// Here, prog counter gets set back to 0, then incremented to 1
	valueThree := 1
	valueFour := 1
	valueTwo = valueThree * valueFour
	if valueOne == valueTwo {
		valueTwo = 1
	} else {
		valueTwo = 0
	}

	// IF VALUE 2 equals 1, WE ARE GOING TO SKIP THE NEXT INSTRUCTION.
	// addi 3 1 3 // UH WHAT, SKIP THE NEXT INSTRUCTION AGAIN???

	// If value 2 equaled 0, this instruction gets skipped b/c the "addi 3 1 3" instruction was run
	if valueTwo == 1 {
		// Yay increment the initial value
		initialValue += valueFour
	}

	valueFour += valueTwo // Since value 2 is only 0 or 1, val 4 gets incremented by 1 if valueOne == valueTwo from earlier

	// value one is a really big number
	if valueFour > valueOne {
		valueTwo = 1
	} else {
		valueTwo = 0
	}

	// IF VALUE 2 equals 1, WE ARE GOING TO SKIP THE NEXT INSTRUCTION.
	// seti 2 3 // UH WHAT, go back to instruction 2, then 3??? While loop????
	// Ok, here is the first while loop from instruction 3 -> instruction 11

	valueThree++
	// Lol, now value three has to be bigger than value one
	if valueThree > valueOne {
		valueTwo = 1
	} else {
		valueTwo = 0
	}

	// IF VALUE 2 equals 1, WE ARE GOING TO SKIP THE NEXT INSTRUCTION.
	// seti 1 3, go back to instruction 1, then 2.  Does this mean i have to fucking repeat the first while loop?
	// Is this a double for loop???

	// mul3 3 3 3
	// ^^^ I'm asuming this breaks the program. Back at instruction 16.
	// At the beginning, instruction 16 got incremented immediately.  Now register 3 is about to get fucked up with a value of 256.
}
