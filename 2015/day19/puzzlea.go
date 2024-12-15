package main

import (
	"regexp"
	"sync"
	"time"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func parseInput(input []string) (map[string][]string, []string) {
	chem := make([]string, 0)

	reactions := make(map[string][]string)
	re := regexp.MustCompile(`(.+) => (.+)`)
	for n, line := range input {
		if line == "" {
			chem = append(chem, input[n+1:]...)
			break
		}
		matches := re.FindAllStringSubmatch(line, -1)
		reactions[matches[0][1]] = append(reactions[matches[0][1]], matches[0][2])
	}
	return reactions, chem
}

func workerCountResults(wg *sync.WaitGroup, res chan string, count chan int) {
	defer wg.Done()

	results := make(map[string]int, 0)
	for {
		select {
		case s := <-res:
			results[s]++
		case <-time.After(time.Second * 1):
			//fmt.Println(results)
			count <- len(results)
			return
		}
	}
}

func reactStep(res chan string, reactions map[string][]string, chem string) {

	for i := 0; i < len(chem); i++ {
		for c, r := range reactions {
			if len(c) > len(chem)-i { // If remaining string too short, skip
				continue
			}
			if chem[i:i+len(c)] == c {
				for _, o := range r {
					out := chem[:i] + o + chem[i+len(c):]
					res <- out
				}
			}
		}
	}

}

func reactMeds(res chan string, reactions map[string][]string, releft string, remain string) {
	//fmt.Println(releft, "  processing ", remain)
	// If the remaining string is empty, we're done, so send it to the result channel
	if len(remain) == 0 {
		res <- releft
		return
	}

	// For each reaction, try to match it to the beginning of the remaining string
	for c, r := range reactions {
		//fmt.Println(remain, " testing reaction ", c)
		if len(c) > len(remain) { // If remaining string too short, skip
			continue
		}
		if remain[:len(c)] == c { // If the reaction matches, react
			for _, o := range r {
				reactMeds(res, reactions, releft+o, remain[len(c):])
			}
		}
	}
	reactMeds(res, reactions, releft+remain[0:0], remain[1:])
}

func puzzlea(inF string) int {
	input := aoclib.ReadStringSlice(inF)
	reactions, chem := parseInput(input)

	// setup channels and wait group
	wg := sync.WaitGroup{}
	wg.Add(1)
	res := make(chan string)
	count := make(chan int)

	// Start our result tracker
	go workerCountResults(&wg, res, count)

	// Process reactions
	reactStep(res, reactions, chem[0])

	resCount := <-count
	wg.Wait()
	return resCount
}
