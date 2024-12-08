package main

import (
	"fmt"
	"io"
	"os"
)

func day1readInput(inputname string) (cola []int, colb []int) {
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
