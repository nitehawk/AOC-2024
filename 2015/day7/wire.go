package main

import (
	"regexp"
	"strconv"
)

const ( // Gate types
	ginvalid = iota
	gset
	gand
	gor
	gnot
	glshift
	grshift
	gpass
)

type wire struct {
	ina, inb string // input wire names -- Pondering pointers
	rv       uint16 // Right side value for shift ops
	gate     int    // Gate type
	value    uint16
	calc     bool
}

func connectWire(cir map[string]wire, wiredesc string) {
	// Regex the wire description
	re := regexp.MustCompile(`(^\d+) -> (.+$)|(^.+) ([A-Z]+) (.+) -> (.+$)|([A-Z]+) (.+) -> (.+$)|(^.+) -> (.+$)`)
	matches := re.FindAllStringSubmatch(wiredesc, -1)
	// Check groups
	if len(matches[0][1]) > 0 { // Set - groups 1-2
		val, _ := strconv.Atoi(matches[0][1])
		cir[matches[0][2]] = wire{gate: gset, value: uint16(val), rv: uint16(val), calc: true}
	} else if len(matches[0][7]) > 0 { // Not - groups 7-9
		if cir[matches[0][8]].gate == ginvalid {
			cir[matches[0][8]] = wire{gate: ginvalid, calc: false}
		}
		cir[matches[0][9]] = wire{ina: matches[0][8], gate: gnot, calc: false}
	} else if len(matches[0][10]) > 0 { // passthrough wire
		if cir[matches[0][10]].gate == ginvalid {
			cir[matches[0][10]] = wire{gate: ginvalid, calc: false}
		}
		cir[matches[0][11]] = wire{ina: matches[0][10], gate: gpass, calc: false}
	} else { // everything else  - groups 3-6
		ina := matches[0][3]
		inb := matches[0][5]
		g := matches[0][6]
		switch matches[0][4] {
		case "LSHIFT":
			rv, _ := strconv.Atoi(inb)
			cir[g] = wire{gate: glshift, ina: ina, rv: uint16(rv), calc: false}
		case "RSHIFT":
			rv, _ := strconv.Atoi(inb)
			cir[g] = wire{gate: grshift, ina: ina, rv: uint16(rv), calc: false}
		case "AND":
			if cir[ina].gate == ginvalid {
				cir[ina] = wire{gate: ginvalid, calc: false}
			}
			if cir[inb].gate == ginvalid {
				cir[inb] = wire{gate: ginvalid, calc: false}
			}
			cir[g] = wire{gate: gand, ina: ina, inb: inb, calc: false}

		case "OR":
			if cir[ina].gate == ginvalid {
				cir[ina] = wire{gate: ginvalid, calc: false}
			}
			if cir[inb].gate == ginvalid {
				cir[inb] = wire{gate: ginvalid, calc: false}
			}
			cir[g] = wire{gate: gor, ina: ina, inb: inb, calc: false}
		}

	}
}

// Return the value of a signal, processing parent wires if needed
func energizeWire(circuit map[string]wire, sig string) uint16 {
	sigwire := circuit[sig]
	if sigwire.calc {
		return sigwire.value
	}
	switch sigwire.gate {
	case ginvalid:
		return 0
	case gpass:
		va := energizeWire(circuit, sigwire.ina)
		sigwire.value = va
		sigwire.calc = true
		circuit[sig] = sigwire
	case gand:
		va := energizeWire(circuit, sigwire.ina)
		vb := energizeWire(circuit, sigwire.inb)
		sigwire.value = va & vb
		sigwire.calc = true
		circuit[sig] = sigwire
	case gor:
		va := energizeWire(circuit, sigwire.ina)
		vb := energizeWire(circuit, sigwire.inb)
		sigwire.value = va | vb
		sigwire.calc = true
		circuit[sig] = sigwire
	case glshift:
		va := energizeWire(circuit, sigwire.ina)
		sigwire.value = va << sigwire.rv
		sigwire.calc = true
		circuit[sig] = sigwire
	case grshift:
		va := energizeWire(circuit, sigwire.ina)
		sigwire.value = va >> sigwire.rv
		sigwire.calc = true
		circuit[sig] = sigwire
	case gnot:
		va := energizeWire(circuit, sigwire.ina)
		sigwire.value = ^va
		sigwire.calc = true
		circuit[sig] = sigwire
	}
	return sigwire.value
}

func energizeCircuit(circuit map[string]wire) {
	energizeWire(circuit, "a")
}
