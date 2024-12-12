package main

import (
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func ingProperties(ingredients []string) map[string]map[string]int {
	goodies := make(map[string]map[string]int)
	re := regexp.MustCompile(`(.+): capacity ([0-9\-]+), durability ([0-9\-]+), flavor ([0-9\-]+), texture ([0-9\-]+), calories ([0-9\-]+)`)

	for _, line := range ingredients {
		matches := re.FindAllStringSubmatch(line, -1)

		cap, _ := strconv.Atoi(matches[0][2])
		dur, _ := strconv.Atoi(matches[0][3])
		flav, _ := strconv.Atoi(matches[0][4])
		tex, _ := strconv.Atoi(matches[0][5])
		cal, _ := strconv.Atoi(matches[0][6])

		ing := map[string]int{"capacity": cap, "durability": dur, "flavor": flav, "texture": tex, "calories": cal}

		goodies[matches[0][1]] = make(map[string]int)
		goodies[matches[0][1]] = ing
	}
	return goodies
}

func cook(ings []string, mix map[string]int, recipe chan map[string]int) {
	cookie := make(map[string]int)
	tgtmix := 100
	curmix := 0

	for k, v := range mix {
		cookie[k] = v
		curmix += v
	}

	if len(ings) == len(cookie) {
		recipe <- cookie
		return
	}

	for i := 1; i <= tgtmix-curmix; i++ {
		cookie[ings[len(mix)]] = i
		cook(ings, cookie, recipe)
	}
}

func rateCookie(props map[string]map[string]int, cookie map[string]int) int {

	score := 1
	pvals := make(map[string]int)

	for k, v := range cookie {
		for pk, pv := range props[k] {
			pvals[pk] += v * pv
		}
	}

	for k, v := range pvals {
		if k == "calories" {
			continue
		}
		if v < 0 {
			return 0
		}
		score *= v
	}
	return score
}

func workerTastyCookies(props map[string]map[string]int, recipe chan map[string]int, bestscore chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	best := -5000
	for {
		select {
		case cookie := <-recipe:
			score := rateCookie(props, cookie)
			if score > best {
				best = score
			}
		case <-time.After(time.Second * 1):
			bestscore <- best
			return
		}

	}
}

func puzzlea(inF string) int {
	ingredients := aoclib.ReadStringSlice(inF)
	props := ingProperties(ingredients)

	ingNames := make([]string, 0)
	for k := range props {
		ingNames = append(ingNames, k)
	}

	// Setup channels
	recipe := make(chan map[string]int, 10000)
	bestscore := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(1)
	// Start happy guest worker
	go workerTastyCookies(props, recipe, bestscore, &wg)

	cook(ingNames, map[string]int{}, recipe)

	best := <-bestscore
	wg.Wait()
	return best
}
