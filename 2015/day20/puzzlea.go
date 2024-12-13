package main

import "fmt"

func puzzlea(inF string) int {
	tgt := 36000000

	house := 500000
	for {
		// Get the score for the current house by backtracking the elves that would deliver here
		score := 0
		for i := 1; i <= house; i++ {
			if house%i == 0 {
				score += i * 10
			}
		}
		if score >= tgt {
			return house
		}
		house++
		if house%10000 == 0 {
			fmt.Printf("House %d, Score %d\n", house, score)
		}
	}
}
