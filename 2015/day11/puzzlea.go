package main

import (
	"fmt"
	"strings"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func validatePass(pass string) bool {
	// Rule 0
	if len(pass) != 8 {
		return false
	}
	// Rule 1 -- sequence
	pbyte := []byte(pass)
	good := false
	for i := 0; i < len(pbyte)-2; i++ {
		if pbyte[i] == pbyte[i+1]-1 && pbyte[i+1] == pbyte[i+2]-1 {
			good = true
			break
		}
	}
	if !good {
		return false
	}

	// Rule 2
	if strings.ContainsAny(pass, "iol") {
		return false
	}

	// Rule 3 - 2 * pair letters
	dbc := 0
	for l := "a"; l <= "z"; l = string(byte(l[0]) + 1) {
		if strings.Contains(pass, l+l) {
			dbc++
		}
	}
	return dbc >= 2
}

func nextPass(pass string) string {
	pbyte := []byte(pass)
	for i := len(pbyte) - 1; i >= 0; i-- {
		if pbyte[i] == 'z' {
			pbyte[i] = 'a'
		} else {
			pbyte[i]++
			break
		}
	}
	return string(pbyte)
}

func puzzlea(inF string) string {
	passlist := aoclib.ReadStringSlice(inF)

	for _, p := range passlist {
		fmt.Printf("%s : %t\n", p, validatePass(p))
	}
	newpass := nextPass(passlist[0])
	gen := 0
	for {
		if validatePass(newpass) {
			gen++
			if gen == 2 {
				return newpass
			}
		}
		newpass = nextPass(newpass)
	}
}
