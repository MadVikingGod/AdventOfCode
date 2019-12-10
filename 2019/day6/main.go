package main

import (
	"fmt"
	"github.com/madvikinggod/AdventOfCode/2018/helpers"
	"strings"
)

type node struct {
	depth  int
	parent string
}

type graph map[string]node

func (g graph) depth(n string) int {

	node, ok := g[n]
	if !ok {
		return 0
	}
	if node.depth > 0 {
		return node.depth
	}
	node.depth = g.depth(node.parent) + 1
	g[n] = node
	return node.depth

}

func (g graph) lineage(n string) graph {
	out := graph{}
	if _, ok := g[n]; !ok {
		return out
	}

	for g[n].parent != "" {
		out[n] = g[n]
		n = g[n].parent
	}
	return out

}

func main() {
	input, err := helpers.GetInput(6)
	if err != nil {
		panic(err)
	}

	g := graph{}

	for _, orbit := range input {
		planets := strings.Split(orbit, ")")

		g[planets[1]] = node{parent: planets[0]}

	}

	sum := 0
	for node := range g {
		sum += g.depth(node)
	}
	fmt.Println(sum)

	youl := g.lineage(g["YOU"].parent)
	sanl := g.lineage(g["SAN"].parent)

	max := 0

	for name, y := range youl {
		if _, ok := sanl[name]; ok && y.depth > max {
			max = y.depth
		}
	}

	// This has the -2 because the depth of the parrent is 1 less then the depth of You and Santa
	fmt.Println(g["YOU"].depth - max + g["SAN"].depth - max - 2)
}
