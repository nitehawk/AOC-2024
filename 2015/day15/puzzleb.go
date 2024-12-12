package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func rateCookieB(props map[string]map[string]int, cookie map[string]int) (int, map[string]int) {

	score := 1
	pvals := make(map[string]int)

	for k, v := range cookie {
		for pk, pv := range props[k] {
			pvals[pk] += v * pv
		}
	}

	for k, v := range pvals {
		if k == "calories" {
			if v == 500 {
				continue
			} else {
				return 0, pvals
			}
		}
		if v < 0 {
			return 0, pvals
		}
		score *= v
	}
	return score, pvals
}

func workerTastyCookiesB(props map[string]map[string]int, recipe chan map[string]int, bestscore chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	best := -5000
	for {
		select {
		case cookie := <-recipe:
			score, pv := rateCookieB(props, cookie)
			if score > best {
				best = score
				fmt.Println(cookie, score, pv)
			}
		case <-time.After(time.Second * 1):
			bestscore <- best
			return
		}

	}
}
func puzzleb(inF string) int {
	ingredients := aoclib.ReadStringSlice(inF)
	props := ingProperties(ingredients)

	ingNames := make([]string, 0)
	for k := range props {
		ingNames = append(ingNames, k)
	}

	// Setup channels
	recipeB := make(chan map[string]int, 10000)
	bestscoreB := make(chan int)

	wgB := sync.WaitGroup{}
	wgB.Add(1)
	// Start happy guest worker
	go workerTastyCookiesB(props, recipeB, bestscoreB, &wgB)

	cook(ingNames, map[string]int{}, recipeB)

	best := <-bestscoreB
	wgB.Wait()
	return best
}
