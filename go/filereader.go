package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(filepath string) []string {
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
