package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func lineToArray(line string, sep string) []int {
	strs := strings.Split(line, sep)
	arr := make([]int, len(strs))

	for i, v := range strs {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Read error converting to int:", err)
			continue
		}
		arr[i] = num
	}
	return arr
}

func readInput(inputname string) (map[int][]int, [][]int) {
	file, err := os.Open(inputname)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", inputname, err))
	}

	scan := bufio.NewScanner(file)
	// Load rules
	rules := make(map[int][]int)
	for scan.Scan() {
		line := scan.Text()
		if len(line) == 0 {
			// Rules loaded
			break
		}
		rule := lineToArray(line, "|")
		rules[rule[0]] = append(rules[rule[0]], rule[1])
	}

	// Load Updates
	var updates [][]int
	for scan.Scan() {
		line := scan.Text()
		updates = append(updates, lineToArray(line, ","))
	}

	return rules, updates
}
