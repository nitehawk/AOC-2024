package main

import (
	"fmt"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func getLeader(ds map[string]deerState) string {
	leader := ""
	for k, v := range ds {
		if leader == "" {
			leader = k
			continue
		}
		if v.distance > ds[leader].distance {
			leader = k
		}
	}
	return leader
}

func scoreLeader(ds map[string]deerState, leaddist int) map[string]deerState {
	newds := make(map[string]deerState, 0)

	for k, v := range ds {
		if v.distance == leaddist {
			newds[k] = deerState{k, v.flying, v.switchtime, v.distance, v.score + 1}
		} else {
			newds[k] = v
		}

	}
	return newds
}

func raceDeerB(herd []deer, t int) int {
	// Setup race
	ds := make(map[string]deerState, 0)

	// Setup starting line
	for _, d := range herd {
		ds[d.name] = deerState{d.name, true, d.flytime, 0, 0}
	}

	lead := getLeader(ds)
	newlead := lead
	// Step through the race
	for i := 0; i < t; i++ {
		for _, d := range herd {
			ds[d.name] = stepDeer(ds[d.name], d)
		}
		newlead = getLeader(ds)
		if lead != newlead {
			fmt.Printf("%s takes the lead at %d - %d - %d\n", newlead, i, ds[newlead].distance, ds[lead].score)
			lead = newlead
		}
		ds = scoreLeader(ds, ds[lead].distance)
	}

	// Get the winner
	zoom := 0
	for _, d := range ds {
		if d.score > zoom {
			zoom = d.score
		}

	}
	return zoom
}
func puzzleb(inF string) int {
	reinspecs := aoclib.ReadStringSlice(inF)

	herd := parseDeer(reinspecs)

	win := raceDeerB(herd, 2503)

	return win
}
