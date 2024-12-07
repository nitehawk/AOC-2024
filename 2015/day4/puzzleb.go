package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func puzzleb(inF string) int {
	key := aoclib.ReadSimpleInput(inF)
	tgtbeg := "000000"

	for i := 1; ; i++ {
		try := fmt.Sprintf("%s%d", key, i)
		hash := md5.Sum([]byte(try))
		hashstr := hex.EncodeToString(hash[:])
		if strings.Compare(tgtbeg, hashstr[:6]) == 0 {
			return i
		}
	}
}
