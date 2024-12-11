package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func loadRelationship(list []string) map[string]map[string]int {
	re := regexp.MustCompile(`(.*) would (.*) (\d+) .*to (.*).`)

	m := make(map[string]map[string]int)

	for _, line := range list {
		matches := re.FindAllStringSubmatch(line, -1)
		change, _ := strconv.Atoi(matches[0][3])
		who := matches[0][1]
		neigh := matches[0][4]
		if matches[0][2] == "lose" {
			change *= -1
		}
		if len(m[who]) == 0 {
			m[who] = make(map[string]int)
		}
		m[who][neigh] = change
	}
	return m
}

func seat(guests []string, seated []string, guestCheck chan []string) {
	if len(guests) == len(seated) {
		guestCheck <- seated
	}
	for _, g := range guests {
		if !slices.Contains(seated, g) {
			newseated := append(seated, g)
			seat(guests, newseated, guestCheck)
		}
	}
}

func workerHappyGuests(relmap map[string]map[string]int, guestCheck chan []string, done chan bool, bestscore chan int) {
	best := -5000
	score := 5
	for {
		select {
		case table := <-guestCheck:
			score = 0
			for gn := range table {
				if gn == 0 {
					score += relmap[table[gn]][table[1]]
					score += relmap[table[gn]][table[len(table)-1]]
				} else if gn == len(table)-1 {
					score += relmap[table[gn]][table[0]]
					score += relmap[table[gn]][table[gn-1]]
				} else {
					score += relmap[table[gn]][table[gn-1]]
					score += relmap[table[gn]][table[gn+1]]
				}
			}
			if score > best {
				best = score
				fmt.Println(table, score)
			}
		case <-done:
			bestscore <- best
			return
		}

	}

}

func puzzlea(inF string) int {
	happy := aoclib.ReadStringSlice(inF)

	rel := loadRelationship(happy)

	// Get the guest list
	guests := make([]string, 0)
	for g := range rel {
		guests = append(guests, g)
	}

	// Setup channels
	done := make(chan bool)
	guestCheck := make(chan []string)
	bestscore := make(chan int)

	// Start happy guest worker
	go workerHappyGuests(rel, guestCheck, done, bestscore)

	// Simulate seating guests
	seat(guests, []string{}, guestCheck)
	done <- true

	best := <-bestscore
	return best
}
