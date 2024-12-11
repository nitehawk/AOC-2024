package main

import "github.com/nitehawk/advent-of-code/aoclib"

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
