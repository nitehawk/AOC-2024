package main

import (
	"slices"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func puzzlea(inF string) int {
	p := aoclib.ReadInputDim(inF)

	// Start from 0
	paper := 0
	for _, pkg := range p {
		rot := make([]int, 0)
		rot = append(rot, pkg[:]...)
		slices.Sort(rot)
		wr := 2*rot[0]*rot[1] + 2*rot[1]*rot[2] + 2*rot[2]*rot[0] + rot[0]*rot[1]
		//fmt.Println(pkg, wr)
		paper += wr
	}
	return paper
}
