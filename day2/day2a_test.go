package main

import "testing"

func TestSafeReport(t *testing.T) {
	report := []int{1, 2, 3, 4, 5}

	if !safeReport(report) {
		t.Fatalf(`safeReport() = false for 1 2 3 4 5`)
	}
}

func TestUnsafeReport(t *testing.T) {
	report1 := []int{1, 1, 3, 4, 5}
	if safeReport(report1) {
		t.Fatalf(`safeReport() = true for 1 1 3 4 5`)
	}

	report2 := []int{1, 5, 7, 10, 11}
	if safeReport(report2) {
		t.Fatalf(`safeReport() = true for 1 5 7 10 11`)
	}

	report3 := []int{1, 2, 1, 2, 1}
	if safeReport(report3) {
		t.Fatalf(`safeReport() = true for 1 2 1 2 1`)
	}
}
