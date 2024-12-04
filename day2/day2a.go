package main

import "fmt"

func safeReport(report []int) bool {
	c := len(report)
	dpri := 0
	for i := 0; i < c-1; i++ {
		d := report[i] - report[i+1]
		if (d < -3) || (d > 3) || (d == 0) {
			return false
		}
		if (d < 0 && dpri > 0) || (d > 0 && dpri < 0) {
			return false
		}
		dpri = d
	}
	return true
}

func day2a() (count int) {
	fmt.Println("Reading day 2 reports")
	reports := day2readInput("input.txt")

	fmt.Println("Day 2 reports: ", len(reports))

	count = 0

	// Loop through the reports and check for safety
	for i := 0; i < len(reports); i++ {
		if safeReport(reports[i]) {
			count++
		}
	}
	return count
}
