package main

import (
	_ "embed"
	"encoding/json"
)

//go:embed input.txt
var input string

func main() {
	println(findSumNumber(marshal(input)))
	println(findSumIgnoreRed(marshal(input)))
}

// findSumNumber finds the sum of all numbers in the input.
func findSumNumber(input interface{}) int {
	switch i := input.(type) {
	case int:
		return i
	case float64:
		return int(i)
	case []interface{}:
		sum := 0
		for _, v := range i {
			sum += findSumNumber(v)
		}
		return sum
	case map[string]interface{}:
		sum := 0
		for _, v := range i {
			sum += findSumNumber(v)
		}
		return sum
	}
	return 0
}

func marshal(input string) interface{} {
	var v interface{}
	err := json.Unmarshal([]byte(input), &v)
	if err != nil {
		panic(err)
	}
	return v
}

func findSumIgnoreRed(input interface{}) int {
	switch i := input.(type) {
	case int:
		return i
	case float64:
		return int(i)
	case []interface{}:
		sum := 0
		for _, v := range i {
			sum += findSumIgnoreRed(v)
		}
		return sum
	case map[string]interface{}:
		sum := 0
		for _, v := range i {
			if v == "red" {
				return 0
			}
			sum += findSumIgnoreRed(v)
		}
		return sum
	}
	return 0
}
