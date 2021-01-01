package main

import "fmt"

func main() {
	value := 0
	for _, line := range input {
		value += Parse(line)
	}
	fmt.Println(value)
}
