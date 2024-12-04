package main

import (
	"sort"
)

func day1a() {
	// Read the input data
	// 2 columns of numbers
	cola, colb := day1readInput("input.txt")

	//	cola := []int{1, 2, 3, 4}
	//	colb := []int{5, 6, 9, 4}

	// Sort the input arrays
	sort.Ints(cola)
	sort.Ints(colb)

	var dist int = 0

	// Loop through the lists
	for i := 0; i < len(cola); i++ {
		a := cola[i]
		b := colb[i]
		d := a - b
		if d < 0 {
			d = -d
		}
		dist += d
	}

	println(dist)
}
