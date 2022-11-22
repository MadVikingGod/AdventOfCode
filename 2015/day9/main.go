package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {

	//128 is too high
	g := ParseInput(input)
	weight, path := g.ShortestPath()
	fmt.Println(weight, path)
	weight, path = g.LongestPath()
	fmt.Println(weight, path)
}

// Graph is a weighted graph of nodes.
type Graph struct {
	Edges map[string]map[string]int
}

// NewGraph returns a new graph.
func NewGraph() *Graph {
	return &Graph{
		Edges: map[string]map[string]int{},
	}
}

// AddEdge adds an edge to the graph.
func (g *Graph) AddEdge(from, to string, weight int) {
	if g.Edges[from] == nil {
		g.Edges[from] = map[string]int{}
	}
	g.Edges[from][to] = weight
	if g.Edges[to] == nil {
		g.Edges[to] = map[string]int{}
	}
	g.Edges[to][from] = weight
}

// Nodes returns the nodes in the graph.
func (g *Graph) Nodes() []string {
	var nodes []string
	for node := range g.Edges {
		nodes = append(nodes, node)
	}
	return nodes
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// permutations returns all permutations of a slice.
func permutations(s []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					arr[0], arr[n-1] = arr[n-1], arr[0]
				} else {
					arr[i], arr[n-1] = arr[n-1], arr[i]
				}
			}
		}
	}

	helper(s, len(s))
	return res
}

// ShortestPath returns the shortest path that visits all nodes.
func (g *Graph) ShortestPath() (int, []string) {
	var shortestPath []string
	var shortestWeight int
	for _, paths := range permutations(g.Nodes()) {
		weight := g.PathWeight(paths)
		if shortestWeight == 0 || weight < shortestWeight {
			shortestWeight = weight
			shortestPath = paths
		}
	}
	return shortestWeight, shortestPath
}

// LongestPath returns the longest path that visits all nodes.
func (g *Graph) LongestPath() (int, []string) {
	var longestPath []string
	var longestWeight int
	for _, paths := range permutations(g.Nodes()) {
		weight := g.PathWeight(paths)
		if weight > longestWeight {
			longestWeight = weight
			longestPath = paths
		}
	}
	return longestWeight, longestPath
}

// PathWeight returns the weight of a path.
func (g *Graph) PathWeight(path []string) int {
	var weight int
	for i := 1; i < len(path); i++ {
		weight += g.Edges[path[i-1]][path[i]]
	}
	return weight
}

// PaseInput parses the input. It returns a graph.
func ParseInput(in string) *Graph {
	g := NewGraph()
	for _, line := range strings.Split(in, "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		from := parts[0]
		to := parts[2]
		weight, _ := strconv.Atoi(parts[4])
		g.AddEdge(from, to, weight)
	}
	return g
}
