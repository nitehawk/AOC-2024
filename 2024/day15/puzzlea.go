package main

import (
	"fmt"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func initWarehouse(whStrings []string) ([][]byte, int, int) {
	warehouse := make([][]byte, 0)

	// Convert strings to byte matrix
	for _, line := range whStrings {
		warehouse = append(warehouse, []byte(line))
	}

	// Find start position
	for y, line := range warehouse {
		for x, b := range line {
			if b == bot {
				warehouse[y][x] = air
				return warehouse, x, y
			}
		}
	}
	return warehouse, 0, 0
}
func printWarehouse(wh *[][]byte) {
	for _, line := range *wh {
		fmt.Println(string(line))
	}
}

func runRobot(wh *[][]byte, bx, by int, route []string) {
	// Step through the route
	for _, line := range route {
		for _, step := range line {
			switch step {
			case up:
				for y := by - 1; y >= 0; y-- {
					if (*wh)[y][bx] == wall {
						break
					}
					if (*wh)[y][bx] == air {
						if by-y > 1 {
							(*wh)[by-1][bx] = air
							(*wh)[y][bx] = box
							by--
						} else {
							by--
						}
						break
					}
				}
			case right:
				for x := bx + 1; x <= len((*wh)[by]); x++ {
					if (*wh)[by][x] == wall {
						break
					}
					if (*wh)[by][x] == air {
						if x-bx > 1 {
							(*wh)[by][bx+1] = air
							(*wh)[by][x] = box
							bx++
						} else {
							bx++
						}
						break
					}
				}
			case down:
				for y := by + 1; y <= len((*wh)); y++ {
					if (*wh)[y][bx] == wall {
						break
					}
					if (*wh)[y][bx] == air {
						if y-by > 1 {
							(*wh)[by+1][bx] = air
							(*wh)[y][bx] = box
							by++
						} else {
							by++
						}
						break
					}
				}
			case left:
				for x := bx - 1; x >= 0; x-- {
					if (*wh)[by][x] == wall {
						break
					}
					if (*wh)[by][x] == air {
						if bx-x > 1 {
							(*wh)[by][bx-1] = air
							(*wh)[by][x] = box
							bx--
						} else {
							bx--
						}
						break
					}
				}

			}
		}
	}

	(*wh)[by][bx] = bot
}

func sumBoxPositions(wh [][]byte) int {
	sum := 0
	for y, line := range wh {
		for x, b := range line {
			if b == box {
				sum += x + y*100
			}
		}
	}
	return sum
}

func puzzlea(inF string) int {
	puzdata := aoclib.ReadStringSlice(inF)

	warehouseStrings := make([]string, 0)
	route := make([]string, 0)

	// Find the blank line
	for i, line := range puzdata {
		if line == "" {
			warehouseStrings = puzdata[:i]
			route = puzdata[i+1:]
			break
		}
	}

	warehouse, sx, sy := initWarehouse(warehouseStrings)

	runRobot(&warehouse, sx, sy, route)

	printWarehouse(&warehouse)

	return sumBoxPositions(warehouse)
}
