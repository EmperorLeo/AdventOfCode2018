package main

import (
	"fmt"
	"strconv"
)

func day1() {
	// Part 1
	scanner, myFile := readFile("../input/day1.txt")
	defer myFile.Close()
	var frequencies []int
	var freq int
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		freq += val
		frequencies = append(frequencies, val)
	}
	fmt.Println(fmt.Sprintf("Frequency is %d", freq))

	// Part 2
	visitedSet := make(map[int]bool)
	freq = 0
	for {
		for _, frequency := range frequencies {
			freq += frequency
			_, exists := visitedSet[freq]
			visitedSet[freq] = true
			if exists {
				fmt.Println(fmt.Sprintf("Repeated frequency is %d", freq))
				return
			}
		}
	}
}
