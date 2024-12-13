package main

import (
	"regexp"
	"strconv"

	"github.com/nitehawk/advent-of-code/aoclib"
)

// return tokens to win claw machine
func winClawMachine(machine []string) int {
	reBut := regexp.MustCompile(`Button (A|B): X\+([0-9]+), Y\+([0-9]+)`)
	rePrize := regexp.MustCompile(`Prize: X=([0-9]+), Y=([0-9]+)`)

	buttonA := reBut.FindStringSubmatch(machine[0])
	buttonB := reBut.FindStringSubmatch(machine[1])
	prize := rePrize.FindStringSubmatch(machine[2])

	buttonAX, _ := strconv.Atoi(buttonA[2])
	buttonAY, _ := strconv.Atoi(buttonA[3])
	buttonBX, _ := strconv.Atoi(buttonB[2])
	buttonBY, _ := strconv.Atoi(buttonB[3])
	prizeX, _ := strconv.Atoi(prize[1])
	prizeY, _ := strconv.Atoi(prize[2])

	//fmt.Println(buttonAX, buttonAY, buttonBX, buttonBY, prizeX, prizeY)
	// Solve for ka & kb
	// px = ka*ax + kb*bx
	// py = ka*ay + kb*by
	// kb = (pxay - pyax) / (bxay - byax)
	// ka = (px - kb*bx) / ax

	kb := (prizeX*buttonAY - prizeY*buttonAX) / (buttonBX*buttonAY - buttonAX*buttonBY)
	ka := (prizeX - kb*buttonBX) / buttonAX

	// Verify the solution
	if ka*buttonAX+kb*buttonBX != prizeX {
		return 0
	}
	if ka*buttonAY+kb*buttonBY != prizeY {
		return 0
	}

	cost := ka*3 + kb
	return cost
}

func puzzlea(inF string) int {
	clawpuz := aoclib.ReadStringSlice(inF)

	totalTokens := 0
	// each machine is 3 lines long followed by a blank line
	for i := 0; i < len(clawpuz); i += 4 {
		machine := clawpuz[i : i+3]
		win := winClawMachine(machine)
		totalTokens += win
	}

	return totalTokens
}
