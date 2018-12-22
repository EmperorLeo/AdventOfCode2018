package main

// Initial value is register 0
// valueOne is register 5
// valueTwo is register 1
// valueThree is register 4
// valueFour is register 2

func elfProg() int {
	// value one = 10,551,381
	// valueOne := 2
	// // 19 is current prog counter
	// valueOne = valueOne * valueOne * 19 * 11
	// // valueOne = valueOne * 19
	// // valueOne = valueOne * 11
	// // 836

	// // twenty two is current prog counter
	// valueTwo := 6*22 + 13
	// valueOne += valueTwo
	// // 981

	// // Here, an instruction is skipped??? Prog counter incremented by 1. I'll write the instruction down in code.
	// // seti 0 -> 3 // Initial value gets set to the program counter??? While loop?

	// // 27 is the current prog counter
	// valueTwo = 27
	// // 28 is the current prog counter
	// valueTwo = 28 * valueTwo
	// // 29 is the current prog counter
	// valueTwo = 29 + valueTwo
	// // 30 is the current prog counter
	// valueTwo = 30 * valueTwo
	// valueTwo = 14 * valueTwo
	// // 32 is the current prog counter
	// valueTwo = 32 * valueTwo
	// valueOne += valueTwo

	incrementor := 10551381
	returnValue := 0

	// Ok, it looks like this little snippet just adds all the factors of the "incrementor" to the return value.
	// Maybe I should look at all the values i needed to multiply to get this awful number LOL.
	// Nvm, i'm gonna cheat and use an online calculator
	// 1 + 3 + 71 + 213 + 49537 + 148611 + 3517127 + 10551381 = 14266944
	for i := 1; i <= incrementor; i++ {
		for j := 1; j < incrementor; j++ {
			intermediate := i * j
			if intermediate == incrementor {
				// addr 4 0 0 (increment the return value by variable 'i' which is register 4.)
				returnValue += i
			}
		}
	}

	return returnValue

	// // Here, prog counter gets set back to 0, then incremented to 1
	// valueThree := 1
	// valueFour := 1
	// valueTwo := valueThree * valueFour
	// if incrementor == valueTwo {
	// 	valueTwo = 1
	// } else {
	// 	valueTwo = 0
	// }

	// IF VALUE 2 equals 1, WE ARE GOING TO SKIP THE NEXT INSTRUCTION.
	// addi 3 1 3 // UH WHAT, SKIP THE NEXT INSTRUCTION AGAIN???

	// // If value 2 equaled 0, this instruction gets skipped b/c the "addi 3 1 3" instruction was run
	// if valueTwo == 1 {
	// 	// Yay increment the initial value
	// 	returnValue += valueFour
	// }

	// valueFour++ // Fuck, this was actually addi, not addr.  Just increment one always.

	// // value one is a really big number
	// if valueFour > incrementor {
	// 	valueTwo = 1
	// } else {
	// 	valueTwo = 0
	// }

	// IF VALUE 2 equals 1, WE ARE GOING TO SKIP THE NEXT INSTRUCTION.
	// seti 2 3 // UH WHAT, go back to instruction 2, then 3??? While loop????
	// Ok, here is the first while loop from instruction 3 -> instruction 11

	// valueThree++
	// // Lol, now value three has to be bigger than value one
	// if valueThree > incrementor {
	// 	valueTwo = 1
	// } else {
	// 	valueTwo = 0
	// }

	// IF VALUE 2 equals 1, WE ARE GOING TO SKIP THE NEXT INSTRUCTION.
	// seti 1 3, go back to instruction 1, then 2.  Does this mean i have to fucking repeat the first while loop?
	// Is this a double for loop???

	// mulr 3 3 3
	// ^^^ I'm asuming this breaks the program. Back at instruction 16.
	// At the beginning, instruction 16 got incremented immediately.  Now register 3 is about to get fucked up with a value of 256.
}
