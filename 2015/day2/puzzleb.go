package main

import (
	"slices"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func puzzleb(inF string) int {
	p := aoclib.ReadInputDim(inF)

	// Start from 0
	ribbon := 0
	for _, pkg := range p {
		rot := make([]int, 0)
		rot = append(rot, pkg[:]...)
		slices.Sort(rot)
		pr := 2*rot[0] + 2*rot[1] + rot[0]*rot[1]*rot[2]
		//fmt.Println(pkg, wr)
		ribbon += pr
	}
	return ribbon
}
