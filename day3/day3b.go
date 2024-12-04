package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func day3b() int {
	sum := 0
	ena := true

	fmt.Println("Compiling search for mul()")
	re := regexp.MustCompile(`(don)\'t\(\)|(do)\(\)|mul\(([0-9]+),([0-9]+)\)`)

	fmt.Println("Opening memory")
	scan := day3readInput("input.txt")

	fmt.Println("Reading memory")
	for scan.Scan() {
		line := scan.Text()
		matches := re.FindAllStringSubmatch(line, -1)
		for i := 0; i < len(matches); i++ {
			cmd := matches[i]
			// Filter do/don't commands
			if cmd[1] == "don" {
				ena = false
				continue
			}
			if cmd[2] == "do" {
				ena = true
				continue
			}
			if ena {
				a, _ := strconv.Atoi(cmd[3])
				b, _ := strconv.Atoi(cmd[4])
				sum += a * b
			}
		}
	}

	return sum
}
