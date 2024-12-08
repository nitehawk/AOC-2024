package main

import (
	"fmt"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func puzzleb(inF string) int {
	puz := aoclib.ReadInputMatrix(inF)

	maxx := len(puz)
	maxy := len(puz[0])

	fmt.Printf("Map size: %d, %d\n", maxx, maxy)

	// Antenna locations
	bc := make(map[byte][][2]int, 0)

	// Walk the map, find antenna locations
	for y, row := range puz {
		for x, col := range row {
			if col != '.' {
				bc[col] = append(bc[col], [2]int{x, y})
			}
		}
	}

	// Store antinode locatgions as a map using sprintf for x,y coords
	anti := make(map[string]int, 0)

	// For each antenna freq, walk antenna location pairs
	for _, freq := range bc {
		for a, ant := range freq {
			for b := a + 1; b < len(freq); b++ {
				// Get points
				ax := ant[0]
				ay := ant[1]
				bx := freq[b][0]
				by := freq[b][1]
				// Get distance between points to figure out slope
				// d = a - b
				// a' = a + d
				// b' = b - d
				dx := ax - bx
				dy := ay - by
				// Adjust slope for obvious special cases
				if dx == dy {
					dx = 1
					dy = 1
				} else if dx == -dy {
					dx = 1
					dy = -1
				} else if dy == 0 {
					dx = 1
				} else if dx == 0 {
					dy = 1
				}
				// Use point a as start point, walk a-d until we exit
				nx := ax
				ny := ay
				for {
					if nx >= 0 && ny >= 0 && nx < maxx && ny < maxy {
						pa := fmt.Sprintf("%d,%d", nx, ny)
						anti[pa]++
					} else {
						break
					}
					nx = nx - dx
					ny = ny - dy
				}

				// Use point a as start point, walk a+d until we exit
				nx = ax + dx
				ny = ay + dy
				for {
					if nx >= 0 && ny >= 0 && nx < maxx && ny < maxy {
						pa := fmt.Sprintf("%d,%d", nx, ny)
						anti[pa]++
					} else {
						break
					}
					nx = nx + dx
					ny = ny + dy
				}
			}
		}

	}

	return len(anti)
}
