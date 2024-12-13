package main

import "github.com/nitehawk/advent-of-code/aoclib"

func stepLightsB(lights [][]bool) [][]bool {
	next := make([][]bool, 0)
	for y := 0; y < len(lights); y++ {
		row := make([]bool, len(lights[0]))
		next = append(next, row)
	}
	for y := 0; y < len(lights); y++ {
		for x := 0; x < len(lights[0]); x++ {
			neighbors := countNeighbors(lights, x, y)
			if lights[y][x] {
				if neighbors == 2 || neighbors == 3 {
					next[y][x] = true
				} else {
					next[y][x] = false
				}
			} else if neighbors == 3 {
				next[y][x] = true
			}
		}
	}

	// Force the corner lights on
	next[0][0] = true
	next[0][len(lights[0])-1] = true
	next[len(lights)-1][0] = true
	next[len(lights)-1][len(lights[0])-1] = true
	return next
}

func puzzleb(inF string) int {
	lightBytes := aoclib.ReadInputMatrix(inF)
	lights := make([][]bool, 0)
	steps := 100

	// Convert lights into bool matrix
	for i := 0; i < len(lightBytes); i++ {
		lightrow := make([]bool, 0)
		for j := 0; j < len(lightBytes[i]); j++ {
			lightrow = append(lightrow, lightBytes[i][j] == '#')
		}
		lights = append(lights, lightrow)
	}

	for i := 1; i <= steps; i++ {
		lights = stepLightsB(lights)
	}
	return countLit(lights)
}
