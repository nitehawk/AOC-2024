package main

import (
	"math"
	"slices"
)

func fixUpdate2(upd []int, rules map[int][]int) int {
	for p, pg := range upd {
		for _, prec := range upd[:p] {
			if slices.Contains(rules[pg], prec) {
				// Current page needs to move towards the beginning of the list
				//fmt.Println("Upd: ", upd, " Moving P:", p, " pg:", pg)
				repl := make([]int, 0)
				repl = append(repl, upd[:p-1]...)
				repl = append(repl, pg)
				repl = append(repl, upd[p-1])
				repl = append(repl, upd[p+1:]...)
				return fixUpdate2(repl, rules)
			}
		}
	}

	mid := (int)(math.Floor((float64)(len(upd)) / 2))
	return upd[mid]
}

func puzzleb2() int {
	//rules, updates := readInput("inputtest.txt")
	rules, updates := readInput("input.txt")

	midsum := 0
	// For each update...
	for _, upd := range updates {
		if !CheckUpdate(upd, rules) {
			midsum += fixUpdate2(upd, rules)
		}
	}

	return midsum
}
