package main

var input = 33100000

func main() {

	divs := make([]int, input/10)
	for i := 2; i < len(divs); i++ {
		for j := i; j < len(divs); j += i {
			divs[j] += i * 10
		}

		if divs[i]+10 >= input {
			println(i)
			break
		}
	}

	divs = make([]int, input/10)
	for i := 1; i < len(divs); i++ {
		count := 0
		for j := i; j < len(divs) && count < 50; j += i {
			divs[j] += i * 11
			count++
		}

		if divs[i] >= input {
			println(i)
			break
		}
	}
}
