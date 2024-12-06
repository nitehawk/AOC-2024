package main

import (
	"bytes"
	"fmt"

	aoclib "github.com/nitehawk/advent-of-code/aoclib"
)

func puzzlea(inF string) int {
	m := aoclib.ReadInputMatrix(inF)

	g := []byte("^")
	o := []byte("#")
	s := []byte("x") // Replace each step we take with this character

	gy, gx := aoclib.FindPos(m, g[0])
	fmt.Println("Obs: ", string(o), " Guard ", gx, ",", gy)

	// Walk the guard
	dir := 1
	out := false
	for !out {
		m[gy][gx] = s[0]
		switch dir {
		case 1: // Up
			// Check to see if stepping up is possible
			if (gy - 1) < 0 {
				// Leaving the map
				out = true
				break
			}
			if m[gy-1][gx] == o[0] {
				dir = 2
				break
			}
			gy -= 1
		case 2: // Right
			if (gx + 1) >= len(m[0]) {
				// Leaving the map
				out = true
				break
			}
			if m[gy][gx+1] == o[0] {
				dir = 3
				break
			}
			gx += 1
		case 3: // Down
			if (gy + 1) >= len(m) {
				// Leaving the map
				out = true
				break
			}
			if m[gy+1][gx] == o[0] {
				dir = 4
				break
			}
			gy += 1
		case 4: // Left
			if (gx - 1) < 0 {
				// Leaving the map
				out = true
				break
			}
			if m[gy][gx-1] == o[0] {
				dir = 1
				break
			}
			gx -= 1
		}
	}

	c := 0
	for _, r := range m {
		c += bytes.Count(r, s)
	}

	return c
}
