package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func puzzleb(inF string) int {
	puz := aoclib.ReadStringSlice(inF)

	distPairs := make(map[string]int)
	destList := make([]string, 0)
	re := regexp.MustCompile(`(.+) to (.+) = (\d+)`)

	for _, line := range puz {
		matches := re.FindAllStringSubmatch(line, -1)
		dist, _ := strconv.Atoi(matches[0][3])
		dira := fmt.Sprintf("%s,%s", matches[0][1], matches[0][2])
		dirb := fmt.Sprintf("%s,%s", matches[0][2], matches[0][1])

		distPairs[dira] = dist
		distPairs[dirb] = dist

		destList = append(destList, matches[0][1], matches[0][2])
	}

	destList = uniqueSlice(destList)

	// Do terrible crimes to create a complete list of routes...
	dl := make([]int, 0)
	routes := make([][]int, 0)
	for i := 0; i < len(destList); i++ {
		dl = append(dl, i)
	}
	for i := 0; i < len(destList); i++ {
		routes = append(routes, route(i, dl)...)
	}

	// At this point, we have a complete list of routes...
	distances := make([]int, 0)
	for _, rt := range routes {
		diR := 0
		for s := 0; s < len(rt)-1; s++ {
			diR += distPairs[fmt.Sprintf("%s,%s", destList[rt[s]], destList[rt[s+1]])]
		}
		distances = append(distances, diR)
	}

	return slices.Max(distances)
}
