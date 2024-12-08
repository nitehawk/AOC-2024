package main

import (
	"fmt"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func puzzlea(inF string) int {
	puz := aoclib.ReadInputMatrix(inF)

	maxx := len(puz)
	maxy := len(puz[0])

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

	// For each antenna freq, walk antenna locations and find antinodes
	for _, freq := range bc {
		for a, ant := range freq {
			for b := a + 1; b < len(freq); b++ {
				ax := ant[0]
				ay := ant[1]
				bx := freq[b][0]
				by := freq[b][1]
				// Get distance between points
				// d = a - b
				// a' = a + d
				// b' = b - d
				dx := ax - bx
				dy := ay - by
				apx := ax + dx
				bpx := bx - dx
				apy := ay + dy
				bpy := by - dy

				// Two antenna locations per antenna pair
				pa := fmt.Sprintf("%d,%d", apx, apy)
				pb := fmt.Sprintf("%d,%d", bpx, bpy)
				if apx >= 0 && apy >= 0 && apx < maxx && apy < maxy {
					anti[pa]++
				}
				if bpx >= 0 && bpy >= 0 && bpx < maxx && bpy < maxy {
					anti[pb]++
				}
				//fmt.Printf("Freq %c: %s, %s\n", f, pa, pb)
			}
		}

	}

	return len(anti)
}
