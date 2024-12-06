package main

import (
	"fmt"
	"slices"

	aoclib "github.com/nitehawk/AOC-2024/aoclib"
)

// yeah.. it's y,x, and y increases going down.   it's how we read the input
type Point struct {
	y int
	x int
}

func cmpPoint(p1 Point, p2 Point) int {
	if p1.x == p2.x && p1.y == p2.y {
		return 0
	}
	return -1
}

// Return true if we exit the map, false otherwise
func walk(m [][]byte, start Point, mark bool, sm byte, obs byte) (bool, []Point) {
	gPos := start
	// Walk the guard
	dir := 1
	out := false
	obsHist := make([]Point, 0)
	for !out {
		if mark {
			m[gPos.y][gPos.x] = sm
		}
		ol := len(obsHist)
		if ol >= 8 && slices.CompareFunc(obsHist[ol-8:ol-4], obsHist[ol-4:], cmpPoint) == 0 {
			return true, obsHist
		}
		switch dir {
		case 1: // Up
			// Check to see if stepping up is possible
			if (gPos.y - 1) < 0 {
				// Leaving the map
				out = true
				break
			}
			if m[gPos.y-1][gPos.x] == obs {
				dir = 2
				obsHist = append(obsHist, Point{gPos.y - 1, gPos.x})
				break
			}
			gPos.y -= 1
		case 2: // Right
			if (gPos.x + 1) >= len(m[0]) {
				// Leaving the map
				out = true
				break
			}
			if m[gPos.y][gPos.x+1] == obs {
				dir = 3
				obsHist = append(obsHist, Point{gPos.y, gPos.x - 1})
				break
			}
			gPos.x += 1
		case 3: // Down
			if (gPos.y + 1) >= len(m) {
				// Leaving the map
				out = true
				break
			}
			if m[gPos.y+1][gPos.x] == obs {
				dir = 4
				obsHist = append(obsHist, Point{gPos.y + 1, gPos.x})
				break
			}
			gPos.y += 1
		case 4: // Left
			if (gPos.x - 1) < 0 {
				// Leaving the map
				out = true
				break
			}
			if m[gPos.y][gPos.x-1] == obs {
				dir = 1
				obsHist = append(obsHist, Point{gPos.y, gPos.x - 1})
				break
			}
			gPos.x -= 1
		}
		if len(obsHist) > 5000 {
			fmt.Println("We're lost...")
			return true, obsHist
		}
	}
	return false, obsHist
}

func puzzleb(inF string) int {
	m := aoclib.ReadInputMatrix(inF)

	g := []byte("^")
	o := []byte("#")

	var gPos Point
	gPos.y, gPos.x = aoclib.FindPos(m, g[0])
	fmt.Println("Obs: ", string(o), " Guard ", gPos)

	// Cycle through entire map
	//  for each '.' on the map, replace with '#' and walk the map checking for loop
	loopCnt := 0
	for r := range m {
		for c, b := range m[r] {
			if b == '.' {
				m[r][c] = o[0]
				e, _ := walk(m, gPos, false, 'x', o[0])
				if e {
					//fmt.Println("Found one!")
					loopCnt++
				}
				m[r][c] = '.'
			}
		}
	}
	return loopCnt
}
