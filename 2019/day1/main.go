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
		f := fuelToLaunch(m)
		// This was added for part two
		sum += f + additionalFuel(f)
	}
	fmt.Println(sum)
}

func fuelToLaunch(i int) int {
	return (i/3)-2
}

func additionalFuel(start int) int {
	s := fuelToLaunch(start)
	// fmt.Println(s)
	if s <=0  {
		return 0
	}
	return s + additionalFuel(s)
}

//4918888 is too high
//4916076