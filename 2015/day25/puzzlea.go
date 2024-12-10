package main

import "fmt"

const (
	mult   = 252533
	div    = 33554393
	tgtrow = 3010
	tgtcol = 3019
	//tgtrow = 2
	//tgtcol = 6
)

func nextCode(code int) int {
	return (code * mult) % div
}

func mapIdx(row, col int) string {
	return fmt.Sprintf("%d,%d", row, col)
}

func puzzlea(inF string) int {
	// 1,1 value from sample
	startCode := 20151125

	code := make(map[string]int, 0)
	nc := startCode

	buildRows := tgtrow + tgtcol + 1

	for row := 1; row <= buildRows; row++ {
		prow := row
		for col := 1; col <= row; col++ {
			code[mapIdx(prow, col)] = nc
			nc = nextCode(nc)
			prow--
		}
	}
	return code[mapIdx(tgtrow, tgtcol)]
}
