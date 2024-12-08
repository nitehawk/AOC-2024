package main

import (
	"regexp"
	"strconv"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func puzzleb(inF string) int {
	inst := aoclib.ReadStringSlice(inF)
	lights := [1000][1000]uint{{0}}

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
					lights[x][y] += 2
				}
			}
		case "turn off":
			for x := bx; x <= ex; x++ {
				for y := by; y <= ey; y++ {
					if lights[x][y] > 0 {
						lights[x][y]--
					}
				}
			}
		case "turn on":
			for x := bx; x <= ex; x++ {
				for y := by; y <= ey; y++ {
					lights[x][y]++
				}
			}
		}

	}

	tb := uint(0)
	for _, row := range lights {
		for _, col := range row {
			tb += col
		}
	}
	return int(tb)
}
