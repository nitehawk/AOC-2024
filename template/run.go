package main

import (
	"fmt"
	"time"
)

func main() {
	day := 0
	inputs := []string{"input.txt", "inputtest.txt"}
	starta := time.Now()
	resa := puzzlea(inputs[0])
	paf := time.Now()
	startb := time.Now()
	resb := puzzleb(inputs[0])
	pbf := time.Now()
	fmt.Printf("Day %d, puzzle a: %d -- time: %s\n", day, resa, paf.Sub(starta).String())
	fmt.Printf("Day %d, puzzle a: %d -- time: %s\n", day, resb, pbf.Sub(startb).String())
}
