package main

import (
	"fmt"
	"github.com/madvikinggod/AdventOfCode/2018/helpers"
	"strconv"
	"strings"
)

func main() {

	input, _ := helpers.GetInput(14)
	recipies = parse(input)

	fmt.Println(countOre(1))
	fmt.Println("------------")
	for i := 1572302; countOre(i) < 1000000000000; i += 1 {
		fmt.Println(i, countOre(i))
	}

}

func countOre(fuel int) int {
	need := map[string]int{
		"FUEL": fuel,
	}

	for !isFinished(need) {
		for res, needAmmount := range need {
			if res == "ORE" {
				continue
			}
			if isReady(res, need) {
				delete(need, res)
				q := quant(needAmmount, recipies[res].quantity)
				for name, recipeAmmount := range recipies[res].recipe {
					need[name] += q * recipeAmmount
				}
			}
		}

	}
	return need["ORE"]
}

//963630

func isFinished(need map[string]int) bool {
	if len(need) > 1 {
		return false
	}
	if _, ok := need["ORE"]; ok {
		return true
	}
	return false
}

func isReady(res string, need map[string]int) bool {
	if res == "ORE" {
		return true
	}
	stuff := map[string]int{}
	for k, v := range need {
		stuff[k] = v
	}
	for len(stuff) > 0 {
		for name := range stuff {
			recipe := recipies[name].recipe
			if _, ok := recipe[res]; ok {
				return false
			}
			delete(stuff, name)
			for component := range recipe {
				if component == "ORE" {
					continue
				}
				stuff[component] = 1
			}
		}
	}
	return true
}

func quant(need, step int) int {
	q := need / step
	if need%step != 0 {
		q += 1
	}

	return q
}

var recipies = map[string]resource{}

type resource struct {
	quantity int
	recipe   map[string]int
}

//157 ORE => 5 NZVS
//165 ORE => 6 DCFZ
//165 ORE => 2 GPVTF
//179 ORE => 7 PSHF
//177 ORE => 5 HKGWZ
//12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ
//7 DCFZ, 7 PSHF => 2 XJWVT
//3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT
//44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL
func parse(input []string) map[string]resource {
	recipies := map[string]resource{}
	for _, line := range input {
		res := strings.Split(line, " => ")
		outString := res[1]
		outStrings := strings.Split(outString, " ")
		name := outStrings[1]
		quant, _ := strconv.Atoi(outStrings[0])

		in := map[string]int{}
		inString := res[0]
		for _, inStr := range strings.Split(inString, ", ") {
			i := strings.Split(inStr, " ")
			n, _ := strconv.Atoi(i[0])
			in[i[1]] = n
		}

		recipies[name] = resource{quantity: quant, recipe: in}
	}
	return recipies

}
