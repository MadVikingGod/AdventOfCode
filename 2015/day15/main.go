package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(maxScore(parseInput(input)))
	fmt.Println(maxScoreWithCalories(parseInput(input)))
}

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

type recipe struct {
	ingredients map[ingredient]int
}

func (r recipe) score() int {
	capacity := 0
	durability := 0
	flavor := 0
	texture := 0
	for ingredient, amount := range r.ingredients {
		capacity += ingredient.capacity * amount
		durability += ingredient.durability * amount
		flavor += ingredient.flavor * amount
		texture += ingredient.texture * amount
	}
	if capacity < 0 {
		capacity = 0
	}
	if durability < 0 {
		durability = 0
	}
	if flavor < 0 {
		flavor = 0
	}
	if texture < 0 {
		texture = 0
	}
	return capacity * durability * flavor * texture
}

func parseInput(input string) []ingredient {
	ingredients := []ingredient{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		ingredient := ingredient{}
		fmt.Sscanf(line, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &ingredient.name, &ingredient.capacity, &ingredient.durability, &ingredient.flavor, &ingredient.texture, &ingredient.calories)
		ingredients = append(ingredients, ingredient)
	}
	return ingredients
}

func maxScore(ingredients []ingredient) int {
	maxScore := 0
	for i := 1; i <= 97; i++ {
		for j := 1; j <= 100-i; j++ {
			for k := 1; k <= 100-i-j; k++ {
				l := 100 - i - j - k
				score := recipe{
					ingredients: map[ingredient]int{
						ingredients[0]: i,
						ingredients[1]: j,
						ingredients[2]: k,
						ingredients[3]: l,
					},
				}.score()
				if score > maxScore {
					maxScore = score
				}
			}
		}
	}
	return maxScore
}

func maxScoreWithCalories(ingredients []ingredient) int {
	maxScore := 0
	for i := 1; i <= 97; i++ {
		for j := 1; j <= 100-i; j++ {
			for k := 1; k <= 100-i-j; k++ {
				l := 100 - i - j - k
				score := recipe{
					ingredients: map[ingredient]int{
						ingredients[0]: i,
						ingredients[1]: j,
						ingredients[2]: k,
						ingredients[3]: l,
					},
				}.score()
				calories := ingredients[0].calories*i + ingredients[1].calories*j + ingredients[2].calories*k + ingredients[3].calories*l
				if calories == 500 && score > maxScore {
					maxScore = score
				}
			}
		}
	}
	return maxScore
}
