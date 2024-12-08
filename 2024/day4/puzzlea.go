package main

// Regex to pull mul operations:   mul\([0-9]+,[0-9]+\)
func puzzlea() int {
	count := 0
	xmas := []byte("XMAS")
	puzzle := readInputMatrix("input.txt")

	// Slow search:
	//    Traverse puzzle matrix looking for start byte
	//    For each starting byte match:
	//       search in all 8 directions for complete match
	//   Can we do this without bounds checking?

	for x := 0; x < len(puzzle); x++ {
		for y := 0; y < len(puzzle[x]); y++ {
			if puzzle[x][y] == xmas[0] {
				if x >= 3 && puzzle[x-1][y] == xmas[1] && puzzle[x-2][y] == xmas[2] && puzzle[x-3][y] == xmas[3] {
					count++
				}
				if x < len(puzzle)-3 && puzzle[x+1][y] == xmas[1] && puzzle[x+2][y] == xmas[2] && puzzle[x+3][y] == xmas[3] {
					count++
				}
				if y >= 3 && puzzle[x][y-1] == xmas[1] && puzzle[x][y-2] == xmas[2] && puzzle[x][y-3] == xmas[3] {
					count++
				}
				if y < len(puzzle[x])-3 && puzzle[x][y+1] == xmas[1] && puzzle[x][y+2] == xmas[2] && puzzle[x][y+3] == xmas[3] {
					count++
				}
				if x >= 3 && y >= 3 && puzzle[x-1][y-1] == xmas[1] && puzzle[x-2][y-2] == xmas[2] && puzzle[x-3][y-3] == xmas[3] {
					count++
				}
				if x >= 3 && y < len(puzzle[x])-3 && puzzle[x-1][y+1] == xmas[1] && puzzle[x-2][y+2] == xmas[2] && puzzle[x-3][y+3] == xmas[3] {
					count++
				}
				if x < len(puzzle)-3 && y >= 3 && puzzle[x+1][y-1] == xmas[1] && puzzle[x+2][y-2] == xmas[2] && puzzle[x+3][y-3] == xmas[3] {
					count++
				}
				if x < len(puzzle)-3 && y < len(puzzle[x])-3 && puzzle[x+1][y+1] == xmas[1] && puzzle[x+2][y+2] == xmas[2] && puzzle[x+3][y+3] == xmas[3] {
					count++
				}
			}
		}
	}

	return count
}
