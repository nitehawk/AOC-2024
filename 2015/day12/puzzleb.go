package main

import "github.com/nitehawk/advent-of-code/aoclib"

func puzzleb(inF string) int {
	puz := aoclib.ReadSimpleInput(inF)
	return jsonWalk(puz, "red")
}
