package main

import "fmt"

func main() {
	day := 1
	inputs := []string{"input.txt", "inputtest.txt"}
	fmt.Printf("Day %d, puzzle a: %d\n", day, puzzlea(inputs[0]))
	fmt.Printf("Day %d, puzzle b: %d\n", day, puzzleb(inputs[0]))
}
