package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/madvikinggod/AdventOfCode/2018/helpers"
)

func main() {
	inputs, err := helpers.GetInput(5)
	if err != nil {
		log.Panic(err)
	}

	input := []byte(inputs[0])
	fmt.Println(input[:5])

	out, cont := react(input)
	for cont {
		out, cont = react(out)
	}
	//fmt.Println(string(out))
	fmt.Println(len(out))

	shortest := len(out)
	for _, char := range []byte("abcdefghijklmnopqrstuvwxyz") {
		out, cont = react(remove(input, char))
		for cont {
			out, cont = react(out)
		}
		if len(out) < shortest {
			shortest = len(out)
		}

		fmt.Println(string(char), len(out))
	}
	fmt.Println(shortest)
}

func iterative(input []byte) int {
	out, cont := react(input)
	for cont {
		out, cont = react(out)
	}
	return len(out)
}

func recuse(input []byte) int {
	return len(reactRecursive(input))
}

func reactRecursive(input []byte) []byte {
	for i := 1; i < len(input); i++ {
		if input[i-1]^byte(' ') == input[i] {
			return reactRecursive(append(input[:i-1], input[i+1:]...))
		}
	}
	return input
}

func react(input []byte) ([]byte, bool) {
	for i := 1; i < len(input); i++ {
		if input[i-1]^byte(' ') == input[i] {
			return append(input[:i-1], input[i+1:]...), true
		}
	}
	return input, false
}

func remove(input []byte, char byte) []byte {
	i := string(input)
	return []byte(strings.Replace(strings.Replace(i, string(char), "", -1), string(char^byte(' ')), "", -1))
}
