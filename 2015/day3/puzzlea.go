package main

import (
	"fmt"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func makeMapCoord(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func puzzlea(inF string) int {
	puz := aoclib.ReadSimpleInput(inF)
	x, y := 0, 0

	del := make(map[string]int, 0)
	del[makeMapCoord(x, y)]++
	for _, step := range puz {
		switch step {
		case '^':
			y++
		case '>':
			x++
		case 'v':
			y--
		case '<':
			x--
		}
		del[makeMapCoord(x, y)]++
	}

	return len(del)
}
