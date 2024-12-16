package main

import (
	"fmt"
	"time"
)

const (
	wall  = '#'
	air   = '.'
	box   = 'O'
	bot   = '@'
	up    = '^'
	down  = 'v'
	left  = '<'
	right = '>'
)

func main() {
	day := 15
	inputs := []string{"input.txt", "inputtest.txt", "inputtestbig.txt"}
	ri := 0
	starta := time.Now()
	resa := puzzlea(inputs[ri])
	paf := time.Now()
	startb := time.Now()
	resb := puzzleb(inputs[ri])
	pbf := time.Now()
	fmt.Printf("Day %d, puzzle a: %d -- time: %s\n", day, resa, paf.Sub(starta).String())
	fmt.Printf("Day %d, puzzle b: %d -- time: %s\n", day, resb, pbf.Sub(startb).String())
}
