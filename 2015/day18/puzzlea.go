package main

import (
	"github.com/nitehawk/advent-of-code/aoclib"
)

func countNeighbors(lights [][]bool, x int, y int) int {
	count := 0
	for xp := x - 1; xp <= x+1; xp++ {
		for yp := y - 1; yp <= y+1; yp++ {
			if xp == x && yp == y {
				continue
			}
			if xp < 0 || yp < 0 || xp > len(lights[0])-1 || yp > len(lights)-1 {
				continue
			}
			if lights[yp][xp] {
				count++
			}
		}
	}

	return count
}

func stepLights(lights [][]bool) [][]bool {
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
	return next
}

func countLit(lights [][]bool) int {
	count := 0
	for y := 0; y < len(lights); y++ {
		for x := 0; x < len(lights[0]); x++ {
			if lights[y][x] {
				count++
			}
		}

	}
	return count
}

func puzzlea(inF string) int {
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
		lights = stepLights(lights)
	}
	return countLit(lights)
}
