package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func uniqueSlice(s []string) []string {
	m := make(map[string]bool)
	for _, e := range s {
		m[e] = true
	}
	r := make([]string, 0, len(m))
	for e := range m {
		r = append(r, e)
	}
	return r
}

// If you value your sanity.... don't look at this
func route(r int, dlist []int) [][]int {
	// End of the list
	if len(dlist) == 2 {
		if dlist[0] == r {
			return [][]int{{dlist[0], dlist[1]}}
		} else {
			return [][]int{{dlist[1], dlist[0]}}
		}
	}

	// More than 2 destinations left in list (including current r)
	// remove r from list
	dr := make([]int, 0)
	ridx := slices.Index(dlist, r)
	dr = append(dr, dlist[:ridx]...)
	dr = append(dr, dlist[ridx+1:]...)

	preroutes := make([][]int, 0)
	for _, nr := range dr {
		routes := route(nr, dr)
		for _, rt := range routes {
			preroutes = append(preroutes, append([]int{r}, rt...))

		}
	}
	return preroutes
}

func puzzlea(inF string) int {
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

	return slices.Min(distances)
}
