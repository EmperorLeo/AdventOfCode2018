package main

import "fmt"

// condition is register 1

func day21decompiled() {
	myValue := 1
	// Instruction 0 -- looks like this first part is useless?
	condition := 123      // 01111011
	for condition != 72 { // 01001000
		// Instruction 1
		condition = condition & 456 //111001000
		// ok, it looks like this just repeats once... for no reason
	}

	// Register 4
	var valueOne int
	// Register 1
	var comparison int

	// // Instruction 6
	// valueOne = 65536 | comparison // 2^16 - Again, valueOne is register 4.... LOL
	// comparison = 3798839          // Insturction 7 - set immediate to register 1

	// // Register 5 + Instruction 8
	// valueTwo := valueOne & 255 // you only get the first 8 bits!!! but this goes to 0 immediately the first run
	// comparison += valueTwo     // 3798839 first run - 1110011111011100110111			Instruction 9
	// comparison &= 16777215     // 16777215				- 111111111111111111111111  ->  1110011111011100110111 (again) (this should destroy all bits more significant than 2^24 - 1)	Instruction 10
	// comparison *= 65899        // idk, but this is instruction 11
	// comparison &= 16777215     // again, destroys all bits greater than 2^24 - 1	Instruction 12

	// if valueOne > 255 { // Instruction 13 - All of the bits in valueOne greater than 255 are zeroed out
	// 	// variable i is in register 2
	// 	newValueOne := 0 // register 5
	// 	for i := 1; i <= valueOne; newValueOne++ {
	// 		i += newValueOne + 1 // Instruction 18
	// 		i *= 256             // LEFT SHIFT 8 DIGITS, 256 = 2^8, this is really valueThree << 8	Instruction 19
	// 	}

	// 	// Go to instruction 25 then 26
	// 	valueOne = newValueOne
	// 	// Go to instruction 7 then 8
	// } else {
	// 	// skip to instruction 27 then 28
	// 	if myValue == comparison { // Ok, looks like you need this program to run a few times until comparison is big enough to get rid of the 17th bit
	// 		// terminate program
	// 		return
	// 	}
	// 	// Go back to instruction 5, then 6 - GOTO: 6.  What variables are kept intact initially? (out of scope) - valueOne, comparison, valueThree
	// }

	// Terminate the program if they are equal!!!!  But to be fair, myValue cannot be zero! (no cheating)
	for comparison == 0 || myValue != comparison {
		// Instruction 6
		valueOne = 65536 | comparison // 2^16 - Again, valueOne is register 4.... LOL
		comparison = 3798839          // Instruction 7 - set immediate to register 1

		for {
			// Register 5 + Instruction 8
			valueTwo := valueOne & 255 // you only get the first 8 bits!!! but this goes to 0 immediately the first run
			comparison += valueTwo     // 3798839 first run - 1110011111011100110111			Instruction 9
			comparison &= 16777215     // 16777215				- 111111111111111111111111  ->  1110011111011100110111 (again) (this should destroy all bits more significant than 2^24 - 1)	Instruction 10
			comparison *= 65899        // idk, but this is instruction 11, first run  ->  11101001001001011010001000101011111101
			comparison &= 16777215     // again, destroys all bits greater than 2^24 - 1	011010001000101011111101	Instruction 12

			if valueOne <= 255 { // Instruction 13 - All of the bits in valueOne greater than 255 are zeroed out
				break
			}

			// variable i is in register 2
			newValueOne := 0 // register 5
			for i := 1; i <= valueOne; newValueOne++ {
				i += newValueOne + 1 // Instruction 18
				i *= 256             // LEFT SHIFT 8 DIGITS, 256 = 2^8, this is really valueThree << 8	Instruction 19
			}

			// Go to instruction 25 then 26
			valueOne = newValueOne // Set valueOne to the number of iterations taken!
			// Go to instruction 7 then 8
		}

		fmt.Printf("The first value of comparison at the end of this loop is %d.\n", comparison)
		break
	}
	// Code ends here
	// So what is the goal here? Make each for/while loop run the fewest amount of times? Maybe only run once?
}
