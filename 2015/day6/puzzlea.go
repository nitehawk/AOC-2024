package main

import (
	"regexp"
	"strconv"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func puzzlea(inF string) int {
	inst := aoclib.ReadStringSlice(inF)
	lights := [1000][1000]bool{{false}}

	re := regexp.MustCompile(`(^.+) (\d+),(\d+) through (\d+),(\d+)`)

	for _, cmd := range inst {
		matches := re.FindAllStringSubmatch(cmd, -1)
		bx, _ := strconv.Atoi(matches[0][2])
		by, _ := strconv.Atoi(matches[0][3])
		ex, _ := strconv.Atoi(matches[0][4])
		ey, _ := strconv.Atoi(matches[0][5])
		switch matches[0][1] {
		case "toggle":
			for x := bx; x <= ex; x++ {
				for y := by; y <= ey; y++ {
					lights[x][y] = !lights[x][y]
				}
			}
		case "turn off":
			for x := bx; x <= ex; x++ {
				for y := by; y <= ey; y++ {
					lights[x][y] = false
				}
			}
		case "turn on":
			for x := bx; x <= ex; x++ {
				for y := by; y <= ey; y++ {
					lights[x][y] = true
				}
			}
		}

	}

	lc := 0
	for _, row := range lights {
		for _, col := range row {
			if col {
				lc++
			}
		}
	}
	return lc
}
