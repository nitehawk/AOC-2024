package main

import (
	"strconv"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func puzzlea(inF string) int {
	puz := aoclib.ReadStringSlice(inF)

	countC := 0
	countD := 0
	for _, str := range puz {
		// Process strings to expand escape sequences
		t, _ := strconv.Unquote(str)
		countC += len(str)
		countD += len(t)

	}
	return countC - countD
}
