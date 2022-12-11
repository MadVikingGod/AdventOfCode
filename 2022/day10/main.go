package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed testInput.txt
var testInput string

func main() {
	sum := 0
	m := machine{
		program:   parseInstructions(input),
		registerX: 1,
	}
	m2 := machine{
		program:   parseInstructions(input),
		registerX: 1,
	}

	left := m.stepBy(20)
	sum += m.registerX * 20
	left = m.stepBy(left + 40)
	sum += m.registerX * 60
	left = m.stepBy(left + 40)
	sum += m.registerX * 100
	left = m.stepBy(left + 40)
	sum += m.registerX * 140
	left = m.stepBy(left + 40)
	sum += m.registerX * 180
	left = m.stepBy(left + 40)
	sum += m.registerX * 220

	println(sum)

	println(m2.drawLine())
	println(m2.drawLine())
	println(m2.drawLine())
	println(m2.drawLine())
	println(m2.drawLine())
	println(m2.drawLine())
}

type machine struct {
	ip      int
	hold    int
	program []instruction

	registerX int
}

func (m *machine) step() {
	if m.hold == 0 {
		m.program[m.ip].Execute(m)
		m.ip++
		m.hold = m.program[m.ip].Lenght()
	}
	m.hold--
}

func (m *machine) drawLine() string {
	out := &strings.Builder{}
	for i := 0; i < 40; i++ {
		if i >= m.registerX-1 && i <= m.registerX+1 {
			out.WriteRune('#')
		} else {
			out.WriteRune('.')
		}
		m.step()
	}
	return out.String()
}

func (m *machine) stepBy(count int) int {
	for count > m.program[m.ip].Lenght() {
		count -= m.program[m.ip].Lenght()
		m.program[m.ip].Execute(m)
		m.ip++
	}
	return count
}

type instruction interface {
	Lenght() int
	Execute(m *machine)
}

type noop struct{}

func (n noop) Lenght() int        { return 1 }
func (n noop) Execute(m *machine) {}

type addx struct {
	amount int
}

func (a addx) Lenght() int { return 2 }
func (a addx) Execute(m *machine) {
	m.registerX += a.amount
}

func parseInstructions(input string) []instruction {
	var inst []instruction
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "noop":
			inst = append(inst, noop{})
		case "addx":
			amount, _ := strconv.Atoi(parts[1])
			inst = append(inst, addx{amount})
		}
	}
	inst = append(inst, noop{}, noop{}, noop{})
	return inst
}
