package main

import "fmt"

func main() {
	day := 0
	inP := "input.txt"
	inT := "inputtest.txt"
	fmt.Printf("Day %d, puzzle a: %d\n", day, puzzlea(inP))
	fmt.Printf("Day %d, puzzle b: %d\n", day, puzzleb(inT))
}
