package main

import (
	"fmt"

	"github.com/nitehawk/advent-of-code/aoclib"
)

const (
	// North
	N = 0
	// East
	E = 1
	// South
	S = 2
	// West
	W = 3
)

func fenceSides(garden *gardeninfo, reg regioninfo) int {
	// Walk the region and count sides
	sides := 0
	crop := byte(reg.crop[0])
	cx := reg.sx
	cy := reg.sy
	// Starting point for a region will be the topmost left corner
	dir := E

	for {
		if sides >= 4 && cx == reg.sx && cy == reg.sy {
			break
		}
		switch dir {
		case N:
			// Check left first
			if cx > 0 && garden.layout[cy][cx-1] == crop {
				dir = W
				sides++
				cx--
				continue
			}
			if cy > 0 && garden.layout[cy-1][cx] == crop {
				cy--
				continue
			}
			if cx < len(garden.layout)-1 && garden.layout[cy][cx+1] == crop {
				dir = E
				sides++
				//cx++
				continue
			}
			// if we get here, we came to a dead end and need to go back
			dir = S
			sides += 2
			continue
		case E:
			// Check left first
			if cy > 0 && garden.layout[cy-1][cx] == crop {
				dir = N
				sides++
				cy--
				continue
			}
			if cx < len(garden.layout[cy])-1 && garden.layout[cy][cx+1] == crop {
				cx++
				continue
			}
			if cy < len(garden.layout)-1 && garden.layout[cy+1][cx] == crop {
				// Weird special case that an immediate right turn doesn't count properly
				if sides == 0 {
					sides++
				}
				dir = S
				sides++
				//cy++
				continue
			}
			// if we get here, we came to a dead end and need to go back
			dir = W
			sides += 2
			continue
		case S:
			// Check left first
			if cx < len(garden.layout)-1 && garden.layout[cy][cx+1] == crop {
				dir = E
				sides++
				cx++
				continue
			}
			if cy < len(garden.layout)-1 && garden.layout[cy+1][cx] == crop {
				cy++
				continue
			}
			if cx > 0 && garden.layout[cy][cx-1] == crop {
				dir = W
				sides++
				//cx--
				continue
			}
			// if we get here, we came to a dead end and need to go back
			dir = N
			sides += 2
			continue
		case W:
			// Check left first
			if cy < len(garden.layout)-1 && garden.layout[cy+1][cx] == crop {
				dir = S
				sides++
				cy++
				continue
			}
			if cx > 0 && garden.layout[cy][cx-1] == crop {
				cx--
				continue
			}
			if cy > 0 && garden.layout[cy-1][cx] == crop {
				dir = N
				sides++
				//cy--
				continue
			}
			// if we get here, we came to a dead end and need to go back
			dir = E
			sides += 2
			continue
		}
		break
	}

	// If the number of sides is odd, add one
	if sides%2 != 0 {
		sides++
	}

	return sides
}

func puzzleb(inF string) int {
	garden := gardeninfo{visited: make([][]bool, 0), regions: make([]regioninfo, 0)}
	garden.layout = aoclib.ReadInputMatrix(inF)

	// initialize visited matrix
	for y := range garden.layout {
		garden.visited = append(garden.visited, make([]bool, len(garden.layout[y])))
	}

	garden.regions = fencePlan(garden)
	cost := 0
	for _, v := range garden.regions {
		sides := fenceSides(&garden, v) // New pricing based on sides
		fmt.Println("Region: ", v.crop, " sides: ", sides, " area: ", v.area)
		cost += sides * v.area
	}
	return cost
}
