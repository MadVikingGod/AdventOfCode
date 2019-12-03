package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/madvikinggod/AdventOfCode/2018/helpers"
)

func main() {
	inputs, err := helpers.GetInput(8)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(len(strings.Split(inputs[0], " ")))
	data := make([]int, 0, len(strings.Split(inputs[0], " ")))
	for _, i := range strings.Split(inputs[0], " ") {
		d, _ := strconv.Atoi(i)
		data = append(data, d)
	}

	root, data := parseNode(data)
	fmt.Println(len(data))
	fmt.Println(root.sum())
	fmt.Println(root.value())
}

type node struct {
	number     int
	childCount int
	metaCount  int
	children   []*node
	metadata   []int
}

func (n *node) sum() int {
	s := 0
	for _, cn := range n.children {
		s = s + cn.sum()
	}
	return s + sum(n.metadata)
}
func (n *node) value() int {
	if n.childCount == 0 {
		return sum(n.metadata)
	}
	s := 0
	for _, m := range n.metadata {
		if m-1 < n.childCount {
			s = s + n.children[m-1].value()
		}
	}
	return s
}

func (n *node) AddNode(newNode *node) {
	n.children = append(n.children, newNode)
}

func parseNode(data []int) (*node, []int) {
	n := &node{
		childCount: data[0],
		metaCount:  data[1],
	}
	data = data[2:]
	for i := 0; i < n.childCount; i++ {
		var cn *node
		cn, data = parseNode(data)
		n.children = append(n.children, cn)
	}
	n.metadata = data[:n.metaCount]
	return n, data[n.metaCount:]
}

func sum(a []int) (x int) {
	for _, b := range a {
		x = x + b
	}
	return x
}
