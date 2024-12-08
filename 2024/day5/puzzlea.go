package main

import (
	"fmt"
	"math"
	"slices"
)

func CheckUpdate(upd []int, rules map[int][]int) bool {
	for p, pg := range upd {
		for _, prec := range upd[:p] {
			if slices.Contains(rules[pg], prec) {
				return false
			}
		}
	}
	return true
}

func puzzlea() int {
	rules, updates := readInput("input.txt")
	fmt.Printf("Input loaded: %d pages with rules, %d updates\n", len(rules), len(updates))

	midsum := 0
	// For each update...
	for _, upd := range updates {
		if CheckUpdate(upd, rules) {
			mid := (int)(math.Floor((float64)(len(upd)) / 2))
			midsum += upd[mid]
		}
	}

	return midsum
}
