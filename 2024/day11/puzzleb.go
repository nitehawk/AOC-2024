package main

import (
	"fmt"
	"sync"

	"github.com/nitehawk/advent-of-code/aoclib"
)

var fivecount int

func optimizeStones(stones []int) ([]int, map[int]int) {
	stoneCounts := make(map[int]int)

	for _, s := range stones {
		stoneCounts[s]++
	}

	uniqStones := make([]int, 0, len(stoneCounts))
	for k := range stoneCounts {
		uniqStones = append(uniqStones, k)
	}
	return uniqStones, stoneCounts

}

func blinkStep(stones []int, step int, d int) int {
	// Count down the depth of steps
	if d == 0 {
		return len(stones)
	}
	d--

	// Step forward
	for i := 1; i <= step; i++ {
		stones = blink(stones)
	}

	// Optimize stone list
	uniqStones, stoneCounts := optimizeStones(stones)
	ucount := make(map[int]int)
	for _, s := range uniqStones {
		ucount[s] = blinkStep([]int{s}, step, d)
	}

	count := 0
	for s, c := range ucount {
		count += c * stoneCounts[s]
	}

	// Give a rough sense of progress
	fivecount++
	if fivecount%100000 == 0 {
		fmt.Printf(".")
	}
	return count
}

func blinkWrapper(stones []int, step int, depth int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	count := blinkStep(stones, step, depth)
	c <- count
}

func puzzleb(inF string) int {
	puzstr := aoclib.ReadSimpleInput(inF)
	stones := aoclib.LineToArray(puzstr, " ")
	blinktgt := 75
	blinkstep := 25

	c := make(chan int, len(stones))
	fivecount = 0
	count := 0
	var wg sync.WaitGroup
	wg.Add(len(stones))
	for _, s := range stones {
		go blinkWrapper([]int{s}, blinkstep, blinktgt/blinkstep, c, &wg)
	}

	wg.Wait()
	close(c)
	for v := range c {
		count += v
	}
	fmt.Printf("Final runs of blinkfive: %d\n", fivecount)
	return count
}
