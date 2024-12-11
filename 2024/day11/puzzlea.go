package main

import (
	"fmt"
	"strconv"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func blink(stones []int) []int {
	newstones := make([]int, 0)

	for _, s := range stones {
		stonestr := strconv.Itoa(s)
		if s == 0 { // Rule 1
			newstones = append(newstones, 1)
		} else if len(stonestr)%2 == 0 { // Rule 2
			left := stonestr[:len(stonestr)/2]
			right := stonestr[len(stonestr)/2:]
			lval, _ := strconv.Atoi(left)
			rval, _ := strconv.Atoi(right)
			newstones = append(newstones, lval)
			newstones = append(newstones, rval)
		} else { // Rule 3
			newstones = append(newstones, s*2024)
		}
	}
	return newstones
}

func puzzlea(inF string) int {
	puzstr := aoclib.ReadSimpleInput(inF)
	stones := aoclib.LineToArray(puzstr, " ")
	blinktgt := 25

	count := 0
	fivecount = 0
	for _, s := range stones {
		count += blinkFive([]int{s}, blinktgt/5)
	}
	fmt.Printf("Final runs of blinkfive: %d\n", fivecount)
	return count
}
