package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/nitehawk/advent-of-code/aoclib"
)

func intWalk(v []interface{}, skip string) int {
	sum := 0
	for _, val := range v {
		switch val.(type) {
		case int, float64:
			sum += int(val.(float64))
		case map[string]interface{}:
			sum += mapWalk(val.(map[string]interface{}), skip)
		case []interface{}:
			sum += intWalk(val.([]interface{}), skip)
		}
	}
	return sum
}
func mapWalk(v map[string]interface{}, skip string) int {
	sum := 0
	// check the map for skip string
	for _, val := range v {
		if val == skip {
			return sum
		}
	}

	for _, val := range v {
		switch val.(type) {
		case int, float64:
			sum += int(val.(float64))
		case map[string]interface{}:
			sum += mapWalk(val.(map[string]interface{}), skip)
		case []interface{}:
			sum += intWalk(val.([]interface{}), skip)
		}
	}
	return sum
}

// walk the JSON and add any numbers found
func jsonWalk(js string, skip string) int {
	sum := 0

	dec := json.NewDecoder(strings.NewReader(js))
	for dec.More() {
		var v interface{}
		if err := dec.Decode(&v); err != nil {
			fmt.Println(v)
			panic(err)
		}
		switch v.(type) {
		case int, float64:
			sum += int(v.(float64))
		case []interface{}:
			sum += intWalk(v.([]interface{}), skip)
		case map[string]interface{}:
			sum += mapWalk(v.(map[string]interface{}), skip)
		}

	}
	return sum
}

func puzzlea(inF string) int {
	puz := aoclib.ReadSimpleInput(inF)
	return jsonWalk(puz, "zzzzzz")
}
