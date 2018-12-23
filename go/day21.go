package main

// condition is register 1

func day23() {
	myValue := 1
	// Instruction 0
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
	// Register 2
	var valueThree int

	// Instruction 6
	valueOne = 65536 | comparison // 2^16
	comparison = 3798839
	// Register 5 + Instruction 8
	valueTwo := valueOne & 255 // you only get the first 8 bits!!! but this goes to 0 immediately the first run
	comparison += valueTwo     // 3798839 first run - 1110011111011100110111
	comparison &= 16777215     // 16777215				- 111111111111111111111111  ->  1110011111011100110111 (again) (this should destroy all bits more significant than 2^24 - 1)
	comparison *= 65899        // idk
	comparison &= 16777215     // again, destroys all bits greater than 2^24 - 1
	if 256 > valueOne {
		valueTwo = 1
	} else {
		valueTwo = 0
	}

	if valueTwo == 1 {
		// skip to instruction 27 then 28
		if myValue == comparison { // Ok, looks like you need this program to run a few times until comparison is big enough to get rid of the 17th bit
			// terminate program
			return
		} else {
			// Go back to instruction 5, then 6

		}
	} else {
		// Instruction 19
		valueThree *= 256
		if valueThree > valueOne {
			valueThree = 1
		} else {
			valueThree = 0
		}

		// Skipping an instruction if value 3 is 1
		if valueThree == 1 {
			// Go to instruction 25 then 26
			valueOne = valueTwo
			// Go to instruction 7 then 8
		} else {
			// Go to instruction 24
			// Increment value two by 1
			valueTwo++
			// go back to instruction 17 then 18
			valueThree = valueTwo + 1
			// loop back to instruction 19
		}
	}
}
