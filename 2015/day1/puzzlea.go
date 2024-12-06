package main

import (
	"strings"

	aoclib "github.com/nitehawk/advent-of-code/aoclib"
)

func puzzlea(inF string) int {
	p := aoclib.ReadSimpleInput(inF)

	fl := strings.Count(p, "(")
	fl -= strings.Count(p, ")")

	return fl
}
