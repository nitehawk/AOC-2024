package main

import (
	"github.com/nitehawk/advent-of-code/aoclib"
)

func reactStepB(reactions map[string][]string, chem string) []string {
	results := make([]string, 0)

	for i := 0; i < len(chem); i++ {
		for c, r := range reactions {
			if len(c) > len(chem)-i { // If remaining string too short, skip
				continue
			}
			if chem[i:i+len(c)] == c {
				for _, o := range r {
					out := chem[:i] + o + chem[i+len(c):]
					results = append(results, out)
				}
			}
		}
	}
	return results
}

func Process(reactions map[string][]string, start string, target string) int {
	steps := 1
	// Take an initial step, filter results
	results := reactStepB(reactions, start)

	// Loop until we find the target
	for {
		// Check for complete matches in the results
		for _, r := range results {
			if r == target {
				return steps
			}
		}

		// Otherwise, take another step
		newresults := make([]string, 0)

		for _, r := range results {
			newresults = append(newresults, reactStepB(reactions, r)...)
		}

		results = newresults
		steps++
	}

}

func puzzleb(inF string) int {
	input := aoclib.ReadStringSlice(inF)
	reactions, chem := parseInput(input)

	// Process reactions
	steps := Process(reactions, "e", chem[0])
	return steps
}
