package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var g = graph{}
var b = graph{}

func main() {
	// input := []string{
	// 	"light red bags contain 1 bright white bag, 2 muted yellow bags.",
	// 	"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
	// 	"bright white bags contain 1 shiny gold bag.",
	// 	"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
	// 	"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
	// 	"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
	// 	"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
	// 	"faded blue bags contain no other bags.",
	// 	"dotted black bags contain no other bags.",
	// }
	// input := []string{
	// 	"shiny gold bags contain 2 dark red bags.",
	// 	"dark red bags contain 2 dark orange bags.",
	// 	"dark orange bags contain 2 dark yellow bags.",
	// 	"dark yellow bags contain 2 dark green bags.",
	// 	"dark green bags contain 2 dark blue bags.",
	// 	"dark blue bags contain 2 dark violet bags.",
	// 	"dark violet bags contain no other bags.",
	// }

	for _, l := range input {
		s := strings.Split(l, " bags contain ")
		rules := newRules(s[1])
		g[s[0]] = rules
	}

	b = g.Reverse()
	fmt.Println(len(b.Find("shiny gold")))
	// fmt.Println(g.Find("shiny gold"))
	fmt.Println(g.Count("shiny gold"))
}

type Rule struct {
	count int
	color string
}

type graph map[string][]Rule

func (g graph) Reverse() graph {
	out := graph{}
	for key, rules := range g {
		for _, r := range rules {
			out.insert(r.color, Rule{r.count, key})
		}
	}
	return out
}

func (g graph) insert(color string, rule Rule) {
	r, ok := g[color]
	if !ok {
		g[color] = []Rule{rule}
		return
	}
	g[color] = append(r, rule)
}

func (g graph) Find(color string) map[string]struct{} {
	if _, ok := g[color]; !ok {
		return nil
	}
	found := map[string]struct{}{}
	for _, r := range g[color] {
		found[r.color] = struct{}{}
		others := g.Find(r.color)
		for o := range others {
			found[o] = struct{}{}
		}
	}
	return found

}

func (g graph) Count(color string) int {

	sum := 1
	for _, r := range g[color] {
		sum += r.count * g.Count(r.color)
	}
	return sum
}

var re = regexp.MustCompile(`(\d+) (\w+ \w+) bags?[,.]`)

func newRules(s string) []Rule {
	matches := re.FindAllStringSubmatch(s, -1)
	if len(matches) == 0 {
		return nil
	}
	rules := make([]Rule, len(matches))
	for i, match := range matches {
		count, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Printf("Could not read: %s", match[1])
			return nil
		}
		rules[i] = Rule{
			count: count,
			color: match[2],
		}
	}
	return rules
}
