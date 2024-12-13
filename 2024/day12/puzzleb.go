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

func fenceSidesB(garden *gardeninfo, reg regioninfo) int {
	// Walk the region counting sides.  Find any regions contained within this region and add their sides to the total
	sides := 0
	crop := byte(reg.crop[0])
	cx := reg.sx
	cy := reg.sy
	// Starting point for a region will be the topmost left corner
	dir := E

	for {
		// If we've hit the edge of the garden, add 1 to the side count
		if cx == 0 || cy == 0 || cx == len(garden.layout[0])-1 || cy == len(garden.layout)-1 {
			sides++
		}
		// If we've hit a different crop, add 1 to the side count
		if cy < 0 || cy >= len(garden.layout) || cx < 0 || cx >= len(garden.layout[cy]) || garden.layout[cy][cx] != crop {
			sides++
		}
		// If we've hit a region that is contained within this one, add its
		// side count to the total
		for _, v := range garden.regions {
			if v.sx > cx && v.sy > cy && v.sx+v.area > cx && v.sy+v.perimeter > cy {
				sides += fenceSidesB(garden, v)
			}
		}

		switch dir {
		case N:
			cy--
		case E:
			cx++
		case S:
			cy++
		case W:
			cx--
		}

		// If we've hit the starting point again, then we're done
		if cx == reg.sx && cy == reg.sy {
			break
		}

		// Move to the next direction
		dir = (dir + 1) % 4
	}

	return sides
}

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
		sides := fenceSidesB(&garden, v) // New pricing based on sides
		fmt.Println("Region: ", v.crop, " sides: ", sides, " area: ", v.area)
		cost += sides * v.area
	}
	return cost
}
