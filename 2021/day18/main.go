package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {

	e := parse("[1,2]")
	fmt.Println(e.String())
	e = parse("[[1,2],3]")
	fmt.Println(e.String())
	e = parse("[9,[8,7]]")
	fmt.Println(e.String())
	e = parse("[[1,9],[8,5]]")
	fmt.Println(e.String())
	e = parse("[[[[1,2],[3,4]],[[5,6],[7,8]]],9]")
	fmt.Println(e.String())
}

type elements struct {
	value int
	depth int
	left  *elements
	right *elements
}

func parse(input string) *elements {
	output := &elements{}
	index := output
	depth := 0
	for _, r := range input {
		switch r {
		case '[':
			depth += 1
		case ']':
			depth -= 1
		case ',':
		default:
			e := &elements{
				value: int(r - '0'),
				depth: depth,
				left:  index,
			}
			index.right = e
			index = index.right
		}
	}
	output = output.right
	output.left = nil
	return output
}

func (e *elements) Add(x, y int) {
	idx := e
	for idx.right != nil {
		idx.depth += 1
		idx = idx.right
	}
	idx.depth += 1
	idx.right = &elements{
		value: x,
		depth: 2,
		left:  idx,
	}
	idx = idx.right
	idx.right = &elements{
		value: y,
		depth: 2,
		left:  idx,
	}
}

func (e *elements) String() string {
	buf := bytes.Buffer{}
	idx := e
	for i := 0; i < idx.depth; i++ {
		buf.WriteRune('[')
	}
	buf.WriteString(strconv.Itoa(idx.value))
	idx = idx.right

	for idx.right != nil {
		if idx.left.depth == idx.depth {
			buf.WriteRune(',')
		}
		if idx.left.depth < idx.depth {
			buf.WriteString(",")
			for i := 0; i < idx.depth-idx.left.depth; i++ {
				buf.WriteRune('[')
			}
		}
		if idx.left.depth > idx.depth {
			for i := 0; i < idx.left.depth-idx.depth; i++ {
				buf.WriteRune(']')
			}
			buf.WriteRune(',')
		}

		buf.WriteString(strconv.Itoa(idx.value))
		idx = idx.right
	}

	buf.WriteRune(',')
	buf.WriteString(strconv.Itoa(idx.value))
	for i := 0; i < idx.depth; i++ {
		buf.WriteRune(']')
	}

	return buf.String()
}
