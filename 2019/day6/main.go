package main


import (
	"github.com/madvikinggod/AdventOfCode/2018/helpers"
	"strings"
	"fmt"
)

type node struct {
	name string
	children []*node
	parent *node
}

func New(name string) *node {
	return &node{
		name: name,
		children: []*node{},
	}
}

func (n *node) depth() int {
	if n.name == "COM" {
		return 0
	}
	return 1 + n.parent.depth()
}

func (n *node) lineage() map[string]int {
	count :=0 
	current := n
	lin := map[string]int {
		n.name: 0,
	}
	for current.name != "COM" {
		current = current.parent
		count++
		lin[current.name] = count
	}
	return lin
}

func (n *node) dist(dest *node) int {
	linN := n.lineage()
	linD := dest.lineage()
	//find min common ancestor
	min := 9000
	ddpeth := 0
	for name, depth := range linN {
		d,ok := linD[name]
		if ok && depth < min {
			min = depth
			ddpeth = d
		}
	}

	return min + ddpeth -2

}


func main() {
	input, err := helpers.GetInput(6)
	if err != nil {
		panic(err)
	}

	input = []string {
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
		"K)YOU",
		"I)SAN",
	}


	head := &node{
		name: "COM",
		children: []*node{},
	}
	graph := map[string]*node {
		"COM": head,
	}
	for _, orbit := range input {
		planets := strings.Split(orbit, ")")
		c, ok := graph[planets[0]]
		if  !ok {
			c = New(planets[0])
			graph[planets[0]] = c
		}
		o, ok := graph[planets[1]]
		if  !ok {
			o = New(planets[1])
			graph[planets[1]] = o
		}
		o.parent = c
		c.children = append(c.children, o)
	}
	fmt.Println(len(graph))
	fmt.Println(graph["COM"].children[0].name)

	sum :=0
	for _, node := range graph {
		sum += node.depth()
	}
	fmt.Println(sum)

	you := graph["YOU"]
	san := graph["SAN"]
	youl := you.lineage()
	sanl := san.lineage()
	
	
	common := map[string][]int{}
	for name,y := range youl {
		if s,ok := sanl[name]; ok{
			common[name] = []int{y,s}
		}
	}
	min := 900
	for _,v := range common {
		if v[0]<min {
			fmt.Println(v)
			min = v[0]
		}
	}

	fmt.Println(common)
	fmt.Println(len(common))

	fmt.Println(you.dist(san))
}

//551