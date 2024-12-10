package main

import "github.com/nitehawk/advent-of-code/aoclib"

func trailStartB(trailmap [][]byte, x, y int) int {
	endmap := traverse(trailmap, x, y)
	score := 0
	for _, v := range endmap {
		score += v
	}

	return score
}

func puzzleb(inF string) int {
	trailmap := aoclib.ReadInputMatrix(inF)

	count := 0
	for y := range trailmap {
		for x := range trailmap[y] {
			if trailmap[y][x] == start {
				count += trailStartB(trailmap, x, y)
			}
		}
	}
	return count
}
