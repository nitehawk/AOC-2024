package main

import (
	"fmt"
	"strconv"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func lookandsay(in string) string {
	var out string

	pc := in[0]
	c := 1
	for i := 1; i < len(in); i++ {
		if in[i] == pc {
			c++
		} else {
			out += strconv.Itoa(c) + string(pc)
			pc = in[i]
			c = 1
		}
	}
	return out + strconv.Itoa(c) + string(pc)
}

func puzzlea(inF string) int {
	start := aoclib.ReadSimpleInput(inF)

	tgt := 40

	for r := 0; r < tgt; r++ {
		fmt.Printf("%d ", r)
		start = lookandsay(start)
	}

	return len(start)
}
