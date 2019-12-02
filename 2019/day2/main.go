package main

import (
	// "strings"
	"fmt"
	// "github.com/madvikinggod/AdventOfCode/2018/helpers"
)

func main() {
	// inputStr, err := helpers.GetInput(2)
	// if err != nil {
	// 	panic(err)
	// }
	// inputNum := strings.Split(inputStr[0], ",")
	input := []int {
		1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,10,1,19,2,9,19,23,2,13,23,27,1,6,27,31,2,6,31,35,2,13,35,39,1,39,10,43,2,43,13,47,1,9,47,51,1,51,13,55,1,55,13,59,2,59,13,63,1,63,6,67,2,6,67,71,1,5,71,75,2,6,75,79,1,5,79,83,2,83,6,87,1,5,87,91,1,6,91,95,2,95,6,99,1,5,99,103,1,6,103,107,1,107,2,111,1,111,5,0,99,2,14,0,0,
	}
	fmt.Println(run(input))

	for noun := 0; noun< 100; noun++ {
		for verb := 0; verb<100; verb ++{
			input[1] = noun
			input[2] = verb
			output := run(input)
			if output == 19690720 {
				fmt.Printf("100 * %d + %d = %d\n", noun, verb, 100*noun+verb)
			}
		}
	}
	

}

func run(input []int) int {
	memory := append([]int(nil), input...)
	position, halt := readOpcode(memory, 0)
	for !halt {
		position, halt = readOpcode(memory, position)
		
	}
	return memory[0]
}

func add(memory []int, position int) int {
	a := memory[position + 1]
	b := memory[position + 2]
	dest := memory[position + 3]
	memory[dest] = memory[a] + memory[b]
	return position + 4
}

func mul(memory []int, position int ) int {
	a := memory[position + 1]
	b := memory[position + 2]
	dest := memory[position + 3]
	memory[dest] = memory[a] * memory[b]
	return position + 4
}

func readOpcode(memory []int, position int) (int, bool) {
	switch memory[position] {
	case 1:
		return add(memory, position), false
	case 2:
		return mul(memory, position), false
	case 99:
		return 0, true
	}
	panic(fmt.Errorf("got invalid opcode %d at position %d", memory[position], position))
}

//4916076 too high