package main

import (
	"strings"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func checkName(name string) bool {
	badpairs := []string{"ab", "cd", "pq", "xy"}
	vowels := "aeiou"
	// Check for bad pairs first
	for _, pair := range badpairs {
		if strings.Contains(name, pair) {
			return false
		}
	}
	// Vowels
	vc := 0
	for _, v := range vowels {
		vc += strings.Count(name, string(v))
	}
	if vc < 3 {
		return false
	}
	// Double letters -- walk the string from the second char and look for pair with prev char
	for l := 1; l < len(name); l++ {
		if name[l-1] == name[l] {
			return true
		}
	}

	return false
}

func puzzlea(inF string) int {
	puz := aoclib.ReadStringSlice(inF)

	goodNames := 0
	for _, name := range puz {
		if checkName(name) {
			goodNames++
		}
	}
	return goodNames
}
