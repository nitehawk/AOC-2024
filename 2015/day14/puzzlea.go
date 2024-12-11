package main

import (
	"regexp"
	"strconv"

	"github.com/nitehawk/advent-of-code/aoclib"
)

type deer struct {
	name     string
	speed    int
	flytime  int
	resttime int
}

type deerState struct {
	name       string
	flying     bool
	switchtime int
	distance   int
	score      int
}

func parseDeer(specs []string) []deer {
	herd := make([]deer, 0)
	re := regexp.MustCompile(`^(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.$`)
	for _, spec := range specs {
		matches := re.FindAllStringSubmatch(spec, -1)
		speed, _ := strconv.Atoi(matches[0][2])
		flytime, _ := strconv.Atoi(matches[0][3])
		resttime, _ := strconv.Atoi(matches[0][4])
		herd = append(herd, deer{matches[0][1], speed, flytime, resttime})
	}
	return herd
}

func stepDeer(ds deerState, d deer) deerState {
	switch ds.flying {
	case true:
		ds.distance += d.speed
		ds.switchtime--
		if ds.switchtime == 0 {
			ds.flying = false
			ds.switchtime = d.resttime
		}
	case false:
		ds.switchtime--
		if ds.switchtime == 0 {
			ds.flying = true
			ds.switchtime = d.flytime
		}
	}
	return ds
}

// Return the best distance
func raceDeer(herd []deer, t int) int {
	// Setup race
	ds := make(map[string]deerState, 0)

	// Setup starting line
	for _, d := range herd {
		ds[d.name] = deerState{d.name, true, d.flytime, 0, 0}
	}

	// Step through the race
	for i := 0; i < t; i++ {
		for _, d := range herd {
			ds[d.name] = stepDeer(ds[d.name], d)
		}
	}

	// Get the winner
	zoom := 0
	for _, d := range ds {
		if d.distance > zoom {
			zoom = d.distance
		}

	}
	return zoom
}

func puzzlea(inF string) int {
	reinspecs := aoclib.ReadStringSlice(inF)

	herd := parseDeer(reinspecs)

	win := raceDeer(herd, 2503)

	return win
}
