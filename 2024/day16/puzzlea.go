package main

import (
	"fmt"

	"github.com/nitehawk/advent-of-code/aoclib"
)

// Base distance from current point to end. Does not care about walls
func getWeight(x, y, dx, dy int) int {
	wx := dx - x
	wy := dy - y

	if wx < 0 {
		wx = -wx
	}
	if wy < 0 {
		wy = -wy
	}
	if wx > 0 && wy > 0 { // if movement in both directions is required, there is at least one turn required.
		return wx + wy + 1000
	}
	return wx + wy
}

func aStar(maze [][]byte, sx, sy, dx, dy int) {

	w := getWeight(sx, sy, dx, dy)
	fmt.Println("Starting weight: ", w)

}

func puzzlea(inF string) int {
	maze := aoclib.ReadInputMatrix(inF)

	sx, sy := aoclib.FindPos(maze, start)
	ex, ey := aoclib.FindPos(maze, end)

	aStar(maze, sx, sy, ex, ey)
	return 0
}
