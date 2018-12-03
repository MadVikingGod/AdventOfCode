package main

import (
	"fmt"
	"log"

	"github.com/madvikinggod/AdventOfCode/2018/helpers"
)

func main() {

	i, err := helpers.GetInput(1)
	if err != nil {
		log.Panic(err)
	}

	input, err := helpers.GetInts(i)
	if err != nil {
		log.Panic(err)
	}

	found := map[int]bool{0: true}
	sum := 0
	var exit *int
	for exit == nil {
		for _, f := range input {
			sum = sum + f
			if ok := found[sum]; ok {
				exit = &sum
				break
			}
			found[sum] = true
		}

	}
	fmt.Println(len(found))
	fmt.Println(*exit)
}
