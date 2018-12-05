package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(filepath string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	return bufio.NewScanner(file), file
}
