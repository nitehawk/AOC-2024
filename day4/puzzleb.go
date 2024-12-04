package main

func puzzleb() int {
	count := 0
	m := byte('M')
	a := byte('A')
	s := byte('S')
	puzzle := readInputMatrix("input.txt")

	// Since the 'X' mas 'A' we're looking for can't be on the edge of the grid...
	for x := 1; x < len(puzzle)-1; x++ {
		for y := 1; y < len(puzzle[x])-1; y++ {
			// Only compare everything if we have a center
			if puzzle[x][y] == a {
				if puzzle[x-1][y-1] == m && puzzle[x+1][y-1] == m && puzzle[x+1][y+1] == s && puzzle[x-1][y+1] == s {
					count++
				} else if puzzle[x-1][y-1] == s && puzzle[x+1][y-1] == m && puzzle[x+1][y+1] == m && puzzle[x-1][y+1] == s {
					count++
				} else if puzzle[x-1][y-1] == s && puzzle[x+1][y-1] == s && puzzle[x+1][y+1] == m && puzzle[x-1][y+1] == m {
					count++
				} else if puzzle[x-1][y-1] == m && puzzle[x+1][y-1] == s && puzzle[x+1][y+1] == s && puzzle[x-1][y+1] == m {
					count++
				}
			}
		}
	}

	return count
}
