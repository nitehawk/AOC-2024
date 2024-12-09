package main

import (
	"strconv"

	"github.com/nitehawk/advent-of-code/aoclib"
)

// Making a second version of this to keep 2 additional versions of info
// A slice of file lengths
// A slice of free space lengths
func loadDiskB(in string) ([]int, [][2]int) {
	disk := make([]int, 0)
	files := make([][2]int, 0)
	nf := -1
	for p, b := range in {
		bi, _ := strconv.Atoi(string(b))
		if p%2 == 0 { // Even puz entries are files
			nf++
			files = append(files, [2]int{len(disk), bi})
			for x := 0; x < bi; x++ {
				disk = append(disk, nf)
			}
		} else { // Odd entries are free blocks
			for x := 0; x < bi; x++ {
				disk = append(disk, -1)
			}
		}

	}
	return disk, files
}

func getFreeSpace(disk []int, l int) (bool, int) {
	sp := 0
	inFree := false
	inL := 0
	for p, id := range disk {
		if id == -1 && inFree {
			inL++
			if inL >= l {
				return true, sp
			}
		} else if id == -1 {
			sp = p
			inL = 1
			inFree = true
			if inL >= l {
				return true, sp
			}
		} else if inFree {
			sp = -1
			inL = 0
			inFree = false
		}
	}
	return false, -1
}

func puzzleb(inF string) int {
	puz := aoclib.ReadSimpleInput(inF)

	// Load disk status
	disk, files := loadDiskB(puz)

	// Compact disk (not the seedy kind)
	// For each file..
	for fn := len(files) - 1; fn > 0; fn-- {
		fi := files[fn]
		// Search from the beginning of disk for a long enough free segment
		found, sp := getFreeSpace(disk, fi[1])
		if found && sp < fi[0] {
			//fmt.Printf("Free space for %d (%d) found at %d\n", fn, fi[1], sp)
			// Move file
			for z := 0; z < fi[1]; z++ {
				disk[sp+z] = fn
				disk[fi[0]+z] = -1
			}
		}
	}
	return checksumDisk(disk)
}
