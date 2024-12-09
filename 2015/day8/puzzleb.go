package main

import (
	"strconv"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func puzzleb(inF string) int {
	puz := aoclib.ReadStringSlice(inF)

	countC := 0
	countD := 0
	for _, str := range puz {
		// Process strings to expand escape sequences
		t := strconv.Quote(str)
		countC += len(str)
		countD += len(t)

	}
	return countD - countC
}
