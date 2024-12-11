package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func profileSetup() (*os.File, *os.File) {

	cf, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}

	mf, merr := os.Create("heap.pprof")
	if merr != nil {
		panic(merr)
	}

	return cf, mf
}

func main() {
	cf, mf := profileSetup()
	// Setup profiling
	pprof.StartCPUProfile(cf)

	day := 12
	inputs := []string{"input.txt", "inputtest.txt"}
	ri := 0
	starta := time.Now()
	resa := puzzlea(inputs[ri])
	paf := time.Now()
	startb := time.Now()
	resb := puzzleb(inputs[ri])
	pprof.WriteHeapProfile(mf)
	pbf := time.Now()
	fmt.Printf("Day %d, puzzle a: %d -- time: %s\n", day, resa, paf.Sub(starta).String())
	fmt.Printf("Day %d, puzzle a: %d -- time: %s\n", day, resb, pbf.Sub(startb).String())

	pprof.StopCPUProfile()
}
