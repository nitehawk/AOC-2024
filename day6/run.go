package main

import "fmt"

func main() {
	day := 0
	//puzzle := readInputMatrix("inputtest.txt")
	puzzle := readInputMatrix("input.txt")
	fmt.Println("Map loaded, x:", len(puzzle), " y:", len(puzzle[0]))

	fmt.Printf("Day %d, puzzle a: %d\n", day, puzzlea(puzzle))
	fmt.Printf("Day %d, puzzle b: %d\n", day, puzzleb(puzzle))
}
