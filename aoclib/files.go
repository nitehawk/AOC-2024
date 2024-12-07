package aoclib

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

func ReadInputDim(inputname string) [][3]int {
	file, err := os.Open(inputname)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", inputname, err))
	}

	dim := make([][3]int, 0)

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		dim = append(dim, [3]int(lineToArray(line, "x")))
	}

	return dim
}

func ReadInputBase(inputname string) int {
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
