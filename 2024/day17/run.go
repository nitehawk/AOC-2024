package main

import (
	"fmt"
	"time"
)

const (
	adv = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

func main() {
	day := 17
	inputs := []string{"input.txt", "inputtest.txt"}
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