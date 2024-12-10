package main

import (
	"fmt"

	"github.com/nitehawk/advent-of-code/aoclib"
)

const (
	start = '0'
	end   = '9'
)

// Walk the trailmap, return the number of 'end' points we get to
func traverse(trailmap [][]byte, px, py int) (endmap map[string]int) {
	endmap = make(map[string]int)
	curheight := trailmap[py][px]
	if curheight == end {
		end := fmt.Sprintf("%d,%d", px, py)
		endmap[end]++
		return endmap
	}
	// Up
	if py > 0 && trailmap[py-1][px] == curheight+1 {
		endmap = traverse(trailmap, px, py-1)
	}
	// Down
	if py < len(trailmap)-1 && trailmap[py+1][px] == curheight+1 {
		for k, v := range traverse(trailmap, px, py+1) {
			endmap[k] += v
		}
	}
	// Right
	if px < len(trailmap[py])-1 && trailmap[py][px+1] == curheight+1 {
		for k, v := range traverse(trailmap, px+1, py) {
			endmap[k] += v
		}
	}
	// Left
	if px > 0 && trailmap[py][px-1] == curheight+1 {
		for k, v := range traverse(trailmap, px-1, py) {
			endmap[k] += v
		}
	}
	return endmap
}

func trailStart(trailmap [][]byte, x, y int) int {
	endmap := traverse(trailmap, x, y)
	return len(endmap)
}

func puzzlea(inF string) int {
	trailmap := aoclib.ReadInputMatrix(inF)

	count := 0
	for y := range trailmap {
		for x := range trailmap[y] {
			if trailmap[y][x] == start {
				count += trailStart(trailmap, x, y)
			}
		}
	}
	return count
}
