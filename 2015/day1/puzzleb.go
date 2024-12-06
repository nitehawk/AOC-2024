package main

import (
	aoclib "github.com/nitehawk/advent-of-code/aoclib"
)

func puzzleb(inF string) int {
	p := aoclib.ReadSimpleInput(inF)

	fl := 0
	for i, c := range p {
		switch c {
		case '(':
			fl++
		case ')':
			fl--
			if fl < 0 {
				return i + 1
			}
		}
	}
	return fl
}
