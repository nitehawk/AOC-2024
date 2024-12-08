package main

import (
	"bufio"
	"fmt"
	"os"
)

func day3readInput(inputname string) *bufio.Scanner {
	file, err := os.Open(inputname)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", inputname, err))
	}

	return bufio.NewScanner(file)
}
