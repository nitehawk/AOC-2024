package main

import (
	"fmt"
	"os"

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

func neighbors(garden *gardeninfo, x int, y int, neigh map[byte]int) map[byte]int {
	// Update neighbors -- we're brute forcing a bit here, but hopefully we'll get there
	// This gets closer, but ends up picking up contained regions as well

	if x == 0 || y == 0 || x == len(garden.layout[y])-1 || y == len(garden.layout)-1 {
		neigh['.']++
	}
	if x > 0 {
		neigh[garden.layout[y][x-1]]++
	}
	if y > 0 {
		neigh[garden.layout[y-1][x]]++
	}
	if x < len(garden.layout[y])-1 {
		neigh[garden.layout[y][x+1]]++
	}
	if y < len(garden.layout)-1 {
		neigh[garden.layout[y+1][x]]++
	}
	// Diagonals
	if x > 0 && y > 0 {
		neigh[garden.layout[y-1][x-1]]++
	}
	if x < len(garden.layout[y])-1 && y > 0 {
		neigh[garden.layout[y-1][x+1]]++
	}
	if x > 0 && y < len(garden.layout)-1 {
		neigh[garden.layout[y+1][x-1]]++
	}
	if x < len(garden.layout[y])-1 && y < len(garden.layout)-1 {
		neigh[garden.layout[y+1][x+1]]++
	}
	return neigh
}

func fenceSides(garden *gardeninfo, reg regioninfo) regioninfo {
	// Walk the region and count sides
	sides := 0
	crop := byte(reg.crop[0])
	cx := reg.sx
	cy := reg.sy
	minx, maxx := cx, cx
	miny, maxy := cy, cy
	neigh := make(map[byte]int)
	// Starting point for a region will be the topmost left corner
	dir := E

	for {
		if sides >= 4 && cx == reg.sx && cy == reg.sy {
			break
		}
		if cx < minx {
			minx = cx
		} else if cx > maxx {
			maxx = cx
		}
		if cy < miny {
			miny = cy
		} else if cy > maxy {
			maxy = cy
		}

		neigh = neighbors(garden, cx, cy, neigh)

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
	return regioninfo{
		crop:      reg.crop,
		sx:        reg.sx,
		sy:        reg.sy,
		area:      reg.area,
		perimeter: reg.perimeter,
		minx:      minx,
		maxx:      maxx,
		miny:      miny,
		maxy:      maxy,
		sides:     sides,
		neighbors: neigh,
	}
}

func findContainedRegions(garden *gardeninfo, o regioninfo) []regioninfo {
	contained := make([]regioninfo, 0)
	for _, i := range garden.regions {
		if i.perimeter*2 >= o.area {
			continue
		}
		if o.minx <= i.minx && o.maxx >= i.maxx && o.miny <= i.miny && o.maxy >= i.maxy {
			contained = append(contained, i)
		}
	}
	if len(contained) == 0 {
		return contained
	}

	// Check contained regions of contained regions
	for cn, c := range contained {
		subregions := findContainedRegions(garden, c)
		if len(subregions) == 0 {
			continue
		}
		fmt.Println("c", contained, "s", subregions)
		for _, sr := range subregions {
			delete(contained[cn].neighbors, sr.crop[0])
		}
	}

	filtered := make([]regioninfo, 0)
	for _, c := range contained {
		good := false
		for k := range c.neighbors {
			if k == c.crop[0] || k == o.crop[0] {
				good = true
				continue
			} else {
				good = false
				break
			}
		}
		if good {
			filtered = append(filtered, c)
		}
	}

	return filtered
}

func adjustContainedRegions(garden *gardeninfo) []regioninfo {
	out, err := os.Create("contained.txt")
	if err != nil {
		fmt.Println("Error opening output file")
	}
	defer out.Close()

	adjusted := make([]regioninfo, 0)

	for _, o := range garden.regions {
		contained := findContainedRegions(garden, o)
		s := fmt.Sprintf("R: %s s: %d,%d a: %d s: %d - Contains:\n", o.crop, o.sx, o.sy, o.area, o.sides)
		for _, c := range contained {
			o.sides += c.sides
			s += fmt.Sprintf("  %s s: %d,%d  s: %d\n", c.crop, c.sx, c.sy, c.sides)
		}
		adjusted = append(adjusted, o)
		if len(contained) > 0 {
			out.Write([]byte(s))
			//fmt.Print(s)
		}
	}

	return adjusted
}

func puzzleb(inF string) int {
	garden := gardeninfo{visited: make([][]bool, 0), regions: make([]regioninfo, 0)}
	garden.layout = aoclib.ReadInputMatrix(inF)

	// initialize visited matrix
	for y := range garden.layout {
		garden.visited = append(garden.visited, make([]bool, len(garden.layout[y])))
	}

	// Find regions
	garden.regions = fencePlan(garden)
	// Get outer side counts
	for n, v := range garden.regions {
		garden.regions[n] = fenceSides(&garden, v) // New pricing based on sides
	}

	garden.regions = adjustContainedRegions(&garden) // New pricing based on sides
	cost := 0

	for _, v := range garden.regions {
		c := v.area * v.sides
		cost += c
		//fmt.Printf("Region %c%d - s: %d p: %d a: %d c: %d\n", v.crop[0], n, v.sides, v.perimeter, v.area, c)
		//fmt.Printf("  s: %d,%d  tl: %d,%d  br: %d,%d\n", v.sx, v.sy, v.minx, v.miny, v.maxx, v.maxy)
	}

	return cost
}
