package main

import (
	"strings"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func checkNameB(name string) bool {
	// Double letters -- walk the string from the second char and look for pair with prev char
	dblPass := false
	for l := 2; l < len(name); l++ {
		if name[l-2] == name[l] {
			dblPass = true
			break
		}
	}

	// CharPairs
	pairPass := false
	for l := 1; l < len(name); l++ {
		pair := name[l-1 : l+1]
		if strings.Count(name, pair) >= 2 {
			pairPass = true
		}
	}

	return dblPass && pairPass
}

func puzzleb(inF string) int {
	puz := aoclib.ReadStringSlice(inF)

	goodNames := 0
	for _, name := range puz {
		if checkNameB(name) {
			goodNames++
		}
	}
	return goodNames
}
