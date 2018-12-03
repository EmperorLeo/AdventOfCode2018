package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input/day1.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var freq int
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		freq += val
	}
	fmt.Println(fmt.Sprintf("Frequency is %d", freq))
}
