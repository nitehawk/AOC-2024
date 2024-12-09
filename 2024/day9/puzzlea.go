package main

import (
	"strconv"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func loadDisk(in string) []int {
	disk := make([]int, 0)
	nf := -1
	for p, b := range in {
		bi, _ := strconv.Atoi(string(b))
		if p%2 == 0 { // Even puz entries are files
			nf++
			for x := 0; x < bi; x++ {
				disk = append(disk, nf)
			}
		} else { // Odd entries are free blocks
			for x := 0; x < bi; x++ {
				disk = append(disk, -1)
			}
		}

	}
	return disk
}

func checksumDisk(disk []int) int {
	// Get filesystem checksum
	sum := 0
	for p, f := range disk {
		if f != -1 {
			sum += p * f
		}
	}

	return sum
}

func puzzlea(inF string) int {
	puz := aoclib.ReadSimpleInput(inF)

	// Load disk status
	disk := loadDisk(puz)

	// Compact disk (not the seedy kind)
	f := 0
	for e := len(disk) - 1; e > f; e-- {
		// If we're on a block from a file
		if disk[e] != -1 {
			// Find the first free block
			for {
				if f > e {
					break
				}
				if disk[f] == -1 {
					break
				}
				f++
			}
			// If we passed the end we're compacting, break out
			if f > e {
				break
			}
			disk[f] = disk[e]
			disk[e] = -1
		}
	}

	return checksumDisk(disk)
}
