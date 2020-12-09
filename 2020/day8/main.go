package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	c := newComputer(input)
	runTillLoop(c)
	fmt.Println(c.acc)
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
	c.Instructions[c.ip](c)
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

func (c *Computer) swapNopJmp()

func runTillLoop(c *Computer) {
	seen := map[int]struct{}{}
	ok := false
	for ; !ok; _, ok = seen[c.ip] {
		seen[c.ip] = struct{}{}
		c.Step()
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

type Instruction func(*Computer)

func standard(f Instruction) Instruction {
	return func(c *Computer) {
		f(c)
		c.ip += 1
	}
}
func Nop() Instruction {
	return standard(func(*Computer) {})
}
func Acc(i int) Instruction {
	return standard(func(c *Computer) {
		c.acc += i
	})
}
func Jmp(i int) Instruction {
	return func(c *Computer) {
		c.ip += i
	}
}

func newInstruction(s string) Instruction {
	sp := strings.Split(strings.ToLower(s), " ")
	switch sp[0] {
	case "nop":
		return Nop()
	case "acc":
		i, _ := strconv.Atoi(sp[1])
		return Acc(i)
	case "jmp":
		i, _ := strconv.Atoi(sp[1])
		return Jmp(i)
	default:
		panic(fmt.Sprintf("unkown command: \"%s\"", s))
	}
}
