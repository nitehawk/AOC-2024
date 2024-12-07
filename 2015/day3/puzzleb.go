package main

import (
	"github.com/nitehawk/advent-of-code/aoclib"
)

func puzzleb(inF string) int {
	puz := aoclib.ReadSimpleInput(inF)
	sx, sy, rx, ry := 0, 0, 0, 0

	del := make(map[string]int, 0)
	del[makeMapCoord(sx, sy)]++
	del[makeMapCoord(rx, ry)]++
	for i, step := range puz {
		if i%2 == 1 { // Odd steps are real santa
			switch step {
			case '^':
				sy++
			case '>':
				sx++
			case 'v':
				sy--
			case '<':
				sx--
			}
			del[makeMapCoord(sx, sy)]++
		} else { // Even steps are Robo-Santa
			switch step {
			case '^':
				ry++
			case '>':
				rx++
			case 'v':
				ry--
			case '<':
				rx--
			}
			del[makeMapCoord(rx, ry)]++
		}

	}

	return len(del)
}
