package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func readInput(inputname string) (cola []int, colb []int) {
	file, err := os.Open(inputname)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", inputname, err))
	}

	var na int
	var nb int

	for {
		_, err := fmt.Fscanf(file, "%d %d\n", &na, &nb)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println(err)
			panic(fmt.Sprintf("Scan Failed %s: %v", inputname, err))
		}
		cola = append(cola, na)
		colb = append(colb, nb)
	}
}

func main() {
	// Read the input data
	// 2 columns of numbers
	cola, colb := readInput("input1")

	//	cola := []int{1, 2, 3, 4}
	//	colb := []int{5, 6, 9, 4}

	// Sort the input arrays
	sort.Ints(cola)
	sort.Ints(colb)

	var dist int = 0

	// Loop through the lists
	for i := 0; i < len(cola); i++ {
		a := cola[i]
		b := colb[i]
		d := a - b
		if d < 0 {
			d = -d
		}
		dist += d
	}

	println(dist)
}
