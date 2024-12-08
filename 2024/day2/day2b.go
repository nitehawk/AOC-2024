package main

import "fmt"

func remove(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

// Figure out if we can remove one element to make it safe
func Dampener(report []int) bool {
	for i := 0; i < len(report); i++ {
		testRep := remove(report, i)
		if safeReport(testRep) {
			return true
		}
	}

	return false
}

func day2b() (count int) {
	fmt.Println("Reading day 2 reports")
	reports := day2readInput("input.txt")

	fmt.Println("Day 2 reports: ", len(reports))

	count = 0

	// Loop through the reports and check for safety
	for i := 0; i < len(reports); i++ {
		if safeReport(reports[i]) {
			count++
		} else {
			if Dampener(reports[i]) {
				count++
			}
		}
	}
	return count
}
