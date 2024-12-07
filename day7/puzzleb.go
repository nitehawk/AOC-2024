package main

import (
	"slices"
	"strconv"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func checkCalB(cal []int) bool {
	t := cal[0]

	// Put the first value on the results list
	res := make([]int, 0)
	res = append(res, cal[1])
	for i := 2; i < len(cal); i++ {
		newres := make([]int, 0)
		for e, v := range res {
			// ||
			a := strconv.Itoa(v)
			b := strconv.Itoa(cal[i])
			rs := a + b
			rv, _ := strconv.Atoi(rs)
			newres = append(newres, rv)
			// *
			newres = append(newres, cal[i]*v)
			// +
			res[e] = res[e] + cal[i]
		}
		res = append(res, newres[:]...)
	}

	return slices.Contains(res, t)
}

func puzzleb(inF string) int {
	puz := aoclib.ReadInputMathList(inF)

	cv := 0
	for _, cal := range puz {
		if checkCalB(cal) {
			cv += cal[0]
		}
	}

	return cv
}
