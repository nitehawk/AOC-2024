package main

import (
	"github.com/nitehawk/advent-of-code/aoclib"
)

// Return additional perimeter for crop at x,y
func mapRegion(garden gardeninfo, x, y int, crop byte) (per int, area int) {
	per = 0
	area = 1
	if x < 0 || y < 0 || x > len(garden.layout[0])-1 || y > len(garden.layout)-1 {
		return 1, 0
	}
	garden.visited[y][x] = true

	if x > 0 && garden.layout[y][x-1] == crop && !garden.visited[y][x-1] {
		psub, parea := mapRegion(garden, x-1, y, crop)
		per += psub
		area += parea
	}
	if x < len(garden.layout[0])-1 && garden.layout[y][x+1] == crop && !garden.visited[y][x+1] {
		psub, parea := mapRegion(garden, x+1, y, crop)
		per += psub
		area += parea
	}
	if y > 0 && garden.layout[y-1][x] == crop && !garden.visited[y-1][x] {
		psub, parea := mapRegion(garden, x, y-1, crop)
		per += psub
		area += parea
	}
	if y < len(garden.layout)-1 && garden.layout[y+1][x] == crop && !garden.visited[y+1][x] {
		psub, parea := mapRegion(garden, x, y+1, crop)
		per += psub
		area += parea
	}
	if x == 0 || garden.layout[y][x-1] != crop {
		per += 1
	}
	if x == len(garden.layout[0])-1 || garden.layout[y][x+1] != crop {
		per += 1
	}
	if y == 0 || garden.layout[y-1][x] != crop {
		per += 1
	}
	if y == len(garden.layout)-1 || garden.layout[y+1][x] != crop {
		per += 1
	}
	return per, area
}

func fencePlan(garden gardeninfo) []regioninfo {
	for y, row := range garden.layout {
		for x, crop := range row {
			if garden.visited[y][x] {
				continue
			}
			per, area := mapRegion(garden, x, y, crop)
			reg := regioninfo{crop: string(crop), sx: x, sy: y, area: area, perimeter: per}
			garden.regions = append(garden.regions, reg)
		}
	}

	return garden.regions
}

func priceFences(regions []regioninfo) int {
	cost := 0
	for _, v := range regions {
		cost += v.area * v.perimeter
	}
	return cost
}

type regioninfo struct {
	crop            string
	sx, sy          int // First node of the region
	area, perimeter int
}
type gardeninfo struct {
	layout  [][]byte
	visited [][]bool
	regions []regioninfo
}

func puzzlea(inF string) int {
	garden := gardeninfo{visited: make([][]bool, 0), regions: make([]regioninfo, 0)}
	garden.layout = aoclib.ReadInputMatrix(inF)

	// initialize visited matrix
	for y := range garden.layout {
		garden.visited = append(garden.visited, make([]bool, len(garden.layout[y])))
	}

	garden.regions = fencePlan(garden)
	cost := priceFences(garden.regions)
	return cost
}
