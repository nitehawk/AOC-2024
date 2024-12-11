package main

import (
	"fmt"

	"github.com/nitehawk/advent-of-code/aoclib"
)

var fivecount int

func blinkFive(stones []int, d int) int {
	// Count down the depth of 5's
	if d == 0 {
		return len(stones)
	}
	d--
	for i := 1; i <= 5; i++ {
		stones = blink(stones)
	}

	count := 0
	for _, s := range stones {
		count += blinkFive([]int{s}, d)
	}
	// Give a rough sense of progress
	fivecount++
	if fivecount%50000000 == 0 {
		fmt.Printf(".")
	}
	return count
}

func puzzleb(inF string) int {
	puzstr := aoclib.ReadSimpleInput(inF)
	stones := aoclib.LineToArray(puzstr, " ")
	blinktgt := 75

	count := 0
	fivecount = 0
	for n, s := range stones {
		count += blinkFive([]int{s}, blinktgt/5)
		fmt.Printf("\n%d of %d: count: %d\n", n, len(stones), count)
	}
	fmt.Printf("Final runs of blinkfive: %d\n", fivecount)
	return count
}
