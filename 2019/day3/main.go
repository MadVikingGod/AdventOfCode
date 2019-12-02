package main

import (
	"fmt"
	"github.com/madvikinggod/AdventOfCode/2018/helpers"
)

func main() {
	input, err := helpers.GetInput(1)
	if err!= nil {
		panic(err)
	}
	fmt.Println(input)
}
