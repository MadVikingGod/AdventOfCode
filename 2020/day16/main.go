package main

import "fmt"

func main() {
	invalids := 0
	validTickets := [][]int{}
	for _, ticket := range input {
		valid := true
		for _, x := range ticket {
			if !anyRule(x) {
				invalids += x
				valid = false
			}
		}
		if valid {
			validTickets = append(validTickets, ticket)
		}
	}
	fmt.Println(invalids)
	validMatrix := [][]int{}
	for _, rule := range rules {
		valids := []int{}
		for column := 0; column < len(input[0]); column++ {
			valid := true
			for i := 0; i < len(validTickets); i++ {
				if !rule(validTickets[i][column]) {
					valid = false
					break
				}
			}
			if valid {
				valids = append(valids, column)
			}
		}
		validMatrix = append(validMatrix, valids)
	}
	fmt.Println(validMatrix)
	fmt.Println(yourTicket[3] * yourTicket[17] * yourTicket[9] * yourTicket[13] * yourTicket[15] * yourTicket[1])
}

func anyRule(x int) bool {
	for _, rule := range rules {
		if rule(x) {
			return true
		}
	}
	return false
}

//
// output = [
// [3]
// [17]
// [9]
// [1]
// [13]
// [15]

// [18]
// [16]
// [5]
// [12]
// [14]
// [4]
// [8]
// [7]
// [19]
// [11]
// [0]
// [6]
// [2]
// [10]
// ]
