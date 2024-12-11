package main

import (
	"sync"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func puzzleb(inF string) int {
	happy := aoclib.ReadStringSlice(inF)

	rel := loadRelationship(happy)
	rel["me"] = make(map[string]int)

	// Get the guest list
	guests := make([]string, 0)
	for g := range rel {
		guests = append(guests, g)
	}

	// Setup channels
	guestCheck := make(chan []string)
	bestscore := make(chan int)

	wgb := sync.WaitGroup{}
	wgb.Add(1)
	// Start happy guest worker
	go workerHappyGuests(rel, guestCheck, bestscore, &wgb)

	// Simulate seating guests
	seat(guests, []string{"Alice"}, guestCheck)

	best := <-bestscore
	wgb.Wait()
	return best
}
