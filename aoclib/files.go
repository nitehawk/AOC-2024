package aoclib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Convert a string to a slice of ints
func LineToArray(line string, sep string) []int {
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

// Read entire input as a concatenated single string
func ReadSimpleInput(inputname string) string {
	file, err := os.Open(inputname)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", inputname, err))
	}

	scan := bufio.NewScanner(file)
	puzzle := ""
	for scan.Scan() {
		puzzle += scan.Text()
	}

	return puzzle
}

// Expected to result in a y*x matrix where all rows are the same length
func ReadInputMatrix(inputname string) [][]byte {
	puzzle := [][]byte{}
	file, err := os.Open(inputname)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", inputname, err))
	}

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		puzzle = append(puzzle, []byte(line))
	}

	return puzzle
}

// Read input into a slice of tripple-ints as dimensions
// lxwxh -- Probably should have writtent this a bit more generic
func ReadInputDim(inputname string) [][3]int {
	file, err := os.Open(inputname)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", inputname, err))
	}

	dim := make([][3]int, 0)

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		dim = append(dim, [3]int(LineToArray(line, "x")))
	}

	return dim
}

// First used 2024-day7
// 123: 1 2 ...
func ReadInputMathList(inputname string) [][]int {
	file, err := os.Open(inputname)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", inputname, err))
	}

	puzzle := make([][]int, 0)

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		d := strings.Replace(line, ":", "", 1)
		r := LineToArray(d, " ")
		puzzle = append(puzzle, r)
	}

	return puzzle
}

func ReadStringSlice(inputname string) []string {
	file, err := os.Open(inputname)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", inputname, err))
	}

	out := make([]string, 0)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		out = append(out, line)
	}

	return out
}

// Template read function
func readInputBase(inputname string) int {
	file, err := os.Open(inputname)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", inputname, err))
	}

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		_ = scan.Text()
	}

	return 0
}
