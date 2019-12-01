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
	masses, err := helpers.GetInts(input)
	if err != nil {
		panic(err)
	}
	sum := 0
	for _,m := range masses {
		sum += fuelToLaunch(m)
	}
	fmt.Println(sum)
}

func fuelToLaunch(i int) int {
	return (i/3)-2
}