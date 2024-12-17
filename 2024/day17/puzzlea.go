package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func eADV() {
	num := float64(comp.reg["A"])
	combo := comp.code[comp.ip+1]
	den := float64(1)
	switch combo {
	case 4:
		den = math.Pow(2, float64(comp.reg["A"]))
	case 5:
		den = math.Pow(2, float64(comp.reg["B"]))
	case 6:
		den = math.Pow(2, float64(comp.reg["C"]))
	default:
		den = math.Pow(2, float64(combo))
	}
	res := int(math.Floor(num / den))
	comp.reg["A"] = res
	comp.ip += 2
}
func eBDV() {
	num := float64(comp.reg["A"])
	combo := comp.code[comp.ip+1]
	den := float64(1)
	switch combo {
	case 4:
		den = math.Pow(2, float64(comp.reg["A"]))
	case 5:
		den = math.Pow(2, float64(comp.reg["B"]))
	case 6:
		den = math.Pow(2, float64(comp.reg["C"]))
	default:
		den = math.Pow(2, float64(combo))
	}
	res := int(math.Floor(num / den))
	comp.reg["B"] = res
	comp.ip += 2
}
func eCDV() {
	num := float64(comp.reg["A"])
	combo := comp.code[comp.ip+1]
	den := float64(1)
	switch combo {
	case 4:
		den = math.Pow(2, float64(comp.reg["A"]))
	case 5:
		den = math.Pow(2, float64(comp.reg["B"]))
	case 6:
		den = math.Pow(2, float64(comp.reg["C"]))
	default:
		den = math.Pow(2, float64(combo))
	}
	res := int(math.Floor(num / den))
	comp.reg["C"] = res
	comp.ip += 2
}

func eOUT() {
	combo := comp.code[comp.ip+1]
	switch combo {
	case 4:
		comp.output = append(comp.output, comp.reg["A"]%8)
	case 5:
		comp.output = append(comp.output, comp.reg["B"]%8)
	case 6:
		comp.output = append(comp.output, comp.reg["C"]%8)
	default:
		comp.output = append(comp.output, comp.code[comp.ip+1]%8)
	}
	comp.ip += 2
}

func eBXL() {
	breg := byte(comp.reg["B"])
	lit := byte(comp.code[comp.ip+1])
	res := breg ^ lit
	comp.reg["B"] = int(res)
	comp.ip += 2
}
func eBST() {
	combo := comp.code[comp.ip+1]
	res := 0
	switch combo {
	case 4:
		res = comp.reg["A"] % 8
	case 5:
		res = comp.reg["B"] % 8
	case 6:
		res = comp.reg["C"] % 8
	default:
		res = comp.code[comp.ip+1] % 8
	}
	comp.reg["B"] = res
	comp.ip += 2
}

func eBXC() {
	breg := byte(comp.reg["B"])
	creg := byte(comp.reg["C"])
	res := breg ^ creg
	comp.reg["B"] = int(res)
	comp.ip += 2
}

func eJNZ() {
	if comp.reg["A"] == 0 {
		comp.ip += 2
		return
	}

	comp.ip = comp.code[comp.ip+1]
}

func decode() {
	switch comp.code[comp.ip] {
	case adv:
		eADV()
	case bxl:
		eBXL()
	case bst:
		eBST()
	case jnz:
		eJNZ()
	case bxc:
		eBXC()
	case out:
		eOUT()
	case bdv:
		eBDV()
	case cdv:
		eCDV()
	default:
		fmt.Println("Unknown instruction - ", comp.code[comp.ip])
		comp.ip += 2
	}
}

type compState struct {
	ip     int
	reg    map[string]int
	code   []int
	output []int
}

// Using a global for this to avoid passing pointers about
var comp compState

func puzzlea(inF string) int {
	statestrings := aoclib.ReadStringSlice(inF)
	re := regexp.MustCompile(`Register ([A-C]): ([0-9]+)`)

	// Initialize global state
	comp.reg = make(map[string]int)
	comp.code = make([]int, 0)
	comp.output = make([]int, 0)
	comp.ip = 0

	// Load state
	blankline := false
	for _, line := range statestrings {
		if line == "" {
			blankline = true
			continue
		}
		if blankline {
			comp.code = append(comp.code, aoclib.LineToArray(line[9:], ",")...)
		} else {
			matches := re.FindAllStringSubmatch(line, -1)
			val, _ := strconv.Atoi(matches[0][2])
			comp.reg[matches[0][1]] = val
		}
	}

	// Run the program
	for comp.ip < len(comp.code) {
		decode()
	}

	// Print the output
	for _, o := range comp.output {
		fmt.Printf("%d,", o)
	}

	fmt.Println("")

	return 0
}
