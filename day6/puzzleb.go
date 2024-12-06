package main

import (
	"fmt"

	aoclib "github.com/nitehawk/AOC-2024/aoclib"
)

func puzzleb(inF string) int {
	m := aoclib.ReadInputMatrix(inF)

	g := []byte("^")
	o := []byte("#")
	s := []byte("x") // Replace each step we take with this character

	sy, sx := aoclib.FindPos(m, g[0])
	fmt.Println("Obs: ", string(o), " Guard ", sy, ",", sx, "  ", s)

	return 0
}
