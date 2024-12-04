package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput(inputname string) int {
	file, err := os.Open(inputname)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", inputname, err))
	}

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		_ = scan.Text()
	}

	return 0
}
