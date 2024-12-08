package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func lineToArray(line string) (report []int) {
	strs := strings.Split(line, " ")
	report = make([]int, len(strs))

	for i, v := range strs {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Read error converting to int:", err)
			continue
		}
		report[i] = num
	}
	return
}

func day2readInput(inputname string) (reports [][]int) {
	file, err := os.Open(inputname)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", inputname, err))
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		report := lineToArray(line)

		reports = append(reports, report)
	}

	return
}
