package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// Regex to pull mul operations:   mul\([0-9]+,[0-9]+\)
func day3a() int {
	sum := 0
	fmt.Println("Compiling search for mul()")
	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)

	fmt.Println("Opening memory")
	scan := day3readInput("input.txt")

	fmt.Println("Reading memory")
	for scan.Scan() {
		line := scan.Text()
		matches := re.FindAllStringSubmatch(line, -1)
		for i := 0; i < len(matches); i++ {
			a, _ := strconv.Atoi(matches[i][1])
			b, _ := strconv.Atoi(matches[i][2])
			sum += a * b
		}
	}

	return sum
}
