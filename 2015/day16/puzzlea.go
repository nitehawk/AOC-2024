package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func parseAunts(auntstrings []string) map[int]map[string]int {
	aunts := make(map[int]map[string]int)

	re := regexp.MustCompile(`^Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)$`)

	for _, line := range auntstrings {
		matches := re.FindAllStringSubmatch(line, -1)
		an, _ := strconv.Atoi(matches[0][1])
		vone, _ := strconv.Atoi(matches[0][3])
		vtwo, _ := strconv.Atoi(matches[0][5])
		vthree, _ := strconv.Atoi(matches[0][7])
		aunts[an] = make(map[string]int)
		aunts[an][matches[0][2]] = vone
		aunts[an][matches[0][4]] = vtwo
		aunts[an][matches[0][6]] = vthree
	}

	return aunts
}

func puzzlea(inF string) int {
	auntstrings := aoclib.ReadStringSlice(inF)
	aunts := parseAunts(auntstrings)

	search := map[string]int{"children": 3, "cats": 7, "samoyeds": 2, "pomeranians": 3, "akitas": 0, "vizslas": 0, "goldfish": 5, "trees": 3, "cars": 2, "perfumes": 1}

	// Filter the list of aunts based on each search criteria
	for k, v := range search {
		fmt.Println("Looking for aunts with ", k, "=", v, "aunts left: ", len(aunts))
		for an, aunt := range aunts {
			if k == "cats" || k == "trees" {
				if aunt[k] != 0 && aunt[k] <= v {
					delete(aunts, an)
				}
			} else if k == "pomeranians" || k == "goldfish" {
				if aunt[k] != 0 && aunt[k] >= v {
					delete(aunts, an)
				}
			} else {
				if aunt[k] != v && aunt[k] != 0 {
					delete(aunts, an)
				}
			}
		}
	}

	fmt.Println(aunts)
	return 0
}
