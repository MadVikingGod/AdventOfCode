package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	c := newComputer(input)
	cPart1 := c.Copy()
	doesTerminate(cPart1)
	fmt.Println(cPart1.acc)

	for i := range c.Instructions {
		c2 := c.Copy()
		c2.swapNopJmp(i)
		if doesTerminate(c2) {
			fmt.Println(c2.acc)
			break
		}
	}
}

type Computer struct {
	acc          int
	ip           int
	Instructions []Instruction
}

func newComputer(input []string) *Computer {
	ins := make([]Instruction, len(input))
	for i, s := range input {
		ins[i] = newInstruction(s)
	}
	return &Computer{Instructions: ins}
}

func (c *Computer) Step() {
	c.Instructions[c.ip].run(c)
}

func (c *Computer) Copy() *Computer {
	ins := make([]Instruction, len(c.Instructions))
	copy(ins, c.Instructions)
	return &Computer{
		acc:          c.acc,
		ip:           c.ip,
		Instructions: ins,
	}
}

func (c *Computer) swapNopJmp(i int) {
	c.Instructions[i] = swapNopJmp(c.Instructions[i])
}

func swapNopJmp(ins Instruction) Instruction {
	switch in := ins.(type) {
	case *Nop:
		return &Jmp{i: in.i}
	case *Jmp:
		return &Nop{i: in.i}
	default:
		return ins
	}
}

func doesTerminate(c *Computer) bool {
	seen := map[int]struct{}{}
	ok := false
	for ; !ok; _, ok = seen[c.ip] {
		seen[c.ip] = struct{}{}
		c.Step()
		if c.ip >= len(c.Instructions) {
			return true
		}
	}
	return false
}

type Instruction interface {
	run(*Computer)
}

type Nop struct {
	i int
}

func (*Nop) run(c *Computer) {
	c.ip += 1
}

type Acc struct {
	i int
}

func (a *Acc) run(c *Computer) {
	c.acc += a.i
	c.ip += 1
}

type Jmp struct {
	i int
}

func (j *Jmp) run(c *Computer) {
	c.ip += j.i
}

func newInstruction(s string) Instruction {
	sp := strings.Split(strings.ToLower(s), " ")
	i, _ := strconv.Atoi(sp[1])
	switch sp[0] {
	case "nop":
		return &Nop{i: i}
	case "acc":
		return &Acc{i: i}
	case "jmp":
		return &Jmp{i: i}
	default:
		panic(fmt.Sprintf("unkown command: \"%s\"", s))
	}
}
