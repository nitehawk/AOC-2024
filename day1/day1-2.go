package main

import (
	"sort"
)

func day1b() {
	// Read the input data
	// 2 columns of numbers
	cola, colb := day1readInput("input.txt")

	//	cola := []int{1, 2, 3, 4}
	//	colb := []int{5, 6, 9, 4}

	// Sort the input arrays
	sort.Ints(cola)
	sort.Ints(colb)

	var sim int = 0

	// Loop through the lists
	for i := 0; i < len(cola); i++ {
		a := cola[i]
		count := 0
		for j := 0; j < len(colb); j++ {
			if a == colb[j] {
				count++
			}
			if a < colb[j] {
				break
			}
		}
		sim += a * count
	}

	println(sim)
}
