package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Part 1
	file, err := os.Open("../input/day1.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
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
