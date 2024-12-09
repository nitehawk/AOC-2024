package main

import (
	"github.com/nitehawk/advent-of-code/aoclib"
)

func puzzlea(inF string) int {
	puz := aoclib.ReadStringSlice(inF)
	circuit := make(map[string]wire)

	// Process the instructions to load the circuit
	for _, cirwire := range puz {
		connectWire(circuit, cirwire)
	}

	// Turn on the circuit
	//fmt.Println(circuit)
	energizeCircuit(circuit)
	//fmt.Println(circuit)

	return int(circuit["a"].value)
}
