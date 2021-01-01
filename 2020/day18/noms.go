package main

import (
	"strconv"
	"strings"
)

type nom interface {
	Eval(int) nom
	Value() int
}

type stack []nom

func (s *stack) Push(n nom) {
	*s = append(*s, n)
}
func (s *stack) Pop() nom {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

func New() *stack {
	return &stack{None{}}
}

func Parse(line string) int {
	s := New()
	for _, token := range strings.Split(line, " ") {
		evalToken(token, s)
	}
	return (*s)[len(*s)-1].Value()
}

func evalToken(token string, s *stack) {
	// (
	if token[0] == '(' {
		s.Push(None{})
		evalToken(token[1:], s)
		return
	}
	// )
	if token[len(token)-1] == ')' {
		evalToken(token[:len(token)-1], s)
		b := s.Pop()
		a := s.Pop()
		s.Push(a.Eval(b.Value()))
		return
	}

	if token == "+" {
		a := s.Pop()
		s.Push(Plus(a.Value()))
		return
	}
	if token == "*" {
		a := s.Pop()
		s.Push(Mul(a.Value()))
		return
	}
	a := s.Pop()
	b, _ := strconv.Atoi(token)
	s.Push(a.Eval(b))
}

type Val int

var _ nom = Val(0)

func (v Val) Eval(int) nom {
	return v
}
func (v Val) Value() int {
	return int(v)
}

type None struct{}

var _ nom = None{}

func (None) Eval(i int) nom {
	return Val(i)
}
func (None) Value() int {
	return 0
}

type Plus int

var _ nom = Plus(0)

func (p Plus) Eval(i int) nom {
	return Val(int(p) + i)
}
func (p Plus) Value() int {
	return int(p)
}

type Mul int

var _ nom = Mul(0)

func (m Mul) Eval(i int) nom {
	return Val(int(m) * i)
}
func (m Mul) Value() int {
	return int(m)
}
