package aoclib

import "slices"

func FindPos(m [][]byte, f byte) (int, int) {
	for y, row := range m {
		x := slices.Index(row, f)
		if x > -1 {
			return y, x
		}
	}
	return -2, -1
}
