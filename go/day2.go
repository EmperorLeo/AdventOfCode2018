package main

import (
	"fmt"
)

func day2() {
	// Part 1
	lines := readFile("../input/day2.txt")
	var twos, threes int
	for _, line := range lines {
		runeMap := make(map[rune]int)
		for _, char := range line {
			runeMap[char]++
		}
		var isTwo, isThree bool
		for _, numRepeats := range runeMap {
			if numRepeats == 2 {
				isTwo = true
			} else if numRepeats == 3 {
				isThree = true
			}
		}
		if isTwo {
			twos++
		}
		if isThree {
			threes++
		}
	}
	fmt.Println(fmt.Sprintf("Checksum = %d", twos*threes))

	// Part 2
	for _, box1 := range lines {
		for _, box2 := range lines {
			var flaws, flawPos int
			for i := range box1 {
				if box1[i] != box2[i] {
					flaws++
					flawPos = i
				}
			}

			if flaws == 1 {
				fmt.Println(fmt.Sprintf("String is: %s", box1[0:flawPos]+box1[flawPos+1:]))
				return
			}
		}
	}
}
