package main

import "fmt"

func puzzleb(m [][]byte) int {
	g := byte('^')
	o := byte('#')

	sx, sy := findPos(m, g)
	fmt.Println("Obs: ", string(o), " Guard ", sx, ",", sy)

	return 0
}
