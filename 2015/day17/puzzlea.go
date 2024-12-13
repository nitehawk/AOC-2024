package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func chooseContainers(containers []int, used []int, nogRem int, set chan []int) {
	for i := 0; i < len(containers); i++ {
		if containers[i] == nogRem {
			nextCont := make([]int, 0)
			nextCont = append(nextCont, used[:]...)
			nextCont = append(nextCont, containers[i])
			set <- nextCont
			continue
		}
		if containers[i] > nogRem {
			continue
		}
		nextCont := make([]int, 0)
		nextCont = append(nextCont, used[:]...)
		nextCont = append(nextCont, containers[i])
		chooseContainers(containers[i+1:], nextCont, nogRem-containers[i], set)
	}
}

func fillContainers(containers []int, nog int, set chan []int) {
	for i := 0; i < len(containers); i++ {
		used := make([]int, 0)
		used = append(used, containers[i])
		if containers[i] == nog {
			set <- used
			continue
		}
		if containers[i] > nog {
			continue
		}
		chooseContainers(containers[i+1:], used, nog-containers[i], set)
	}

}

func workerOptionCounter(wg *sync.WaitGroup, nog int, set chan []int, count chan int) {
	defer wg.Done()
	options := 0

	for {
		select {
		case <-set:
			options++
		case <-time.After(time.Second * 1):
			count <- options
			return
		}
	}
}

func puzzlea(inF string) int {
	containers := aoclib.ReadIntSlice(inF)
	nog := 150

	fmt.Println(containers, " - ", nog)

	// Setup channels
	set := make(chan []int, 50)
	count := make(chan int)

	// Setup sync wait group
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Start counting worker
	go workerOptionCounter(&wg, nog, set, count)

	// Find options to store nog
	fillContainers(containers, nog, set)
	// Get count
	opts := <-count

	// Sync up
	wg.Wait()

	return opts
}
