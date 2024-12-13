package main

import (
	"sync"
	"time"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func workerOptionCounterB(wg *sync.WaitGroup, nog int, set chan []int, count chan int) {
	defer wg.Done()
	options := 0
	best := 1000

	for {
		select {
		case s := <-set:
			if len(s) < best {
				best = len(s)
				options = 1
			} else if len(s) == best {
				options++
			}
		case <-time.After(time.Second * 1):
			count <- options
			return
		}
	}
}
func puzzleb(inF string) int {
	containers := aoclib.ReadIntSlice(inF)
	nog := 150

	// Setup channels
	set := make(chan []int, 50)
	count := make(chan int)

	// Setup sync wait group
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Start counting worker
	go workerOptionCounterB(&wg, nog, set, count)

	// Find options to store nog
	fillContainers(containers, nog, set)
	// Get count
	opts := <-count

	// Sync up
	wg.Wait()

	return opts
}
