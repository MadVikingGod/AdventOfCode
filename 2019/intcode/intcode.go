package intcode

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type executor func() bool
type Intcode struct {
	Memory         []int
	ProgramCounter int

	opcodes map[int]executor

	in  io.Reader
	out io.Writer

	//This is just noodeling for registers that may be addded
	registers [10]int
}

func New(memory []int) *Intcode {
	ic := &Intcode{
		Memory: append([]int(nil), memory...),
		in:     os.Stdin,
		out:    os.Stdout,
	}
	ic.register()
	return ic
}

type reader func(int) int

func (ic Intcode) Read(position int) int {
	return ic.Memory[position]
}
func (ic Intcode) ReadPtr(position int) int {
	return ic.Memory[ic.Memory[position]]
}

func (ic *Intcode) Write(position, data int) {
	ic.Memory[position] = data
}
func (ic *Intcode) WritePtr(position, data int) {
	ic.Memory[ic.Memory[position]] = data
}

func (ic *Intcode) Run() (int, error) {
	halt, err := ic.decode()
	for !halt && err == nil {
		halt, err = ic.decode()
	}
	if err != nil {
		return 0, err
	}
	return ic.Read(0), nil

}
func (ic *Intcode) decode() (bool, error) {
	if ic.ProgramCounter >= len(ic.Memory) {
		return false, fmt.Errorf("program ran off memory: %d", ic.ProgramCounter)
	}
	oc := ic.Memory[ic.ProgramCounter]
	f, ok := ic.opcodes[oc]
	if !ok {
		return false, fmt.Errorf("not implemented, opcode %d at %d", oc, ic.ProgramCounter)
	}
	return f(), nil

}

func (ic *Intcode) register() {
	if ic == nil {
		return
	}
	if ic.opcodes == nil {
		ic.opcodes = map[int]executor{}
	}

	ic.registerTrinary(1, add2)
	ic.registerTrinary(2, mul2)

	ic.registerBinary(5, ic.jmpeq)
	ic.registerBinary(6, ic.jmpneq)

	ic.registerTrinary(7, less)
	ic.registerTrinary(8, equal)

	ic.opcodes[3] = ic.store(ic.in)
	ic.opcodes[4] = ic.Uniary(ic.ReadPtr, ic.output)
	ic.opcodes[104] = ic.Uniary(ic.Read, ic.output)

	ic.opcodes[99] = executor(func() bool {
		return true
	})
}

func (ic *Intcode) registerTrinary(opcode int, f func(int, int) int) {
	ic.opcodes[opcode] = ic.trinary(ic.ReadPtr, ic.ReadPtr, f)
	ic.opcodes[100+opcode] = ic.trinary(ic.Read, ic.ReadPtr, f)
	ic.opcodes[1000+opcode] = ic.trinary(ic.ReadPtr, ic.Read, f)
	ic.opcodes[1100+opcode] = ic.trinary(ic.Read, ic.Read, f)
}
func (ic *Intcode) registerBinary(opcode int, f func(reader, reader) executor) {
	ic.opcodes[opcode] = f(ic.ReadPtr, ic.ReadPtr)
	ic.opcodes[100+opcode] = f(ic.Read, ic.ReadPtr)
	ic.opcodes[1000+opcode] = f(ic.ReadPtr, ic.Read)
	ic.opcodes[1100+opcode] = f(ic.Read, ic.Read)

}

func (ic *Intcode) trinary(reader1, reader2 reader, f func(int, int) int) executor {
	return func() bool {

		pc := ic.ProgramCounter
		i1 := reader1(pc + 1)
		i2 := reader2(pc + 2)

		out := f(i1, i2)

		ic.WritePtr(pc+3, out)
		ic.ProgramCounter += 4
		return false
	}
}

func (ic *Intcode) jmpeq(reader1, reader2 reader) executor {
	return func() bool {
		pc := ic.ProgramCounter
		val := reader1(pc + 1)
		if val != 0 {
			dst := reader2(pc + 2)
			ic.ProgramCounter = dst
			return false
		}
		ic.ProgramCounter += 3
		return false
	}
}
func (ic *Intcode) jmpneq(reader1, reader2 reader) executor {
	return func() bool {
		pc := ic.ProgramCounter
		val := reader1(pc + 1)
		if val == 0 {
			dst := reader2(pc + 2)
			ic.ProgramCounter = dst
			return false
		}
		ic.ProgramCounter += 3
		return false
	}
}

func (ic *Intcode) store(input io.Reader) executor {
	return func() bool {
		pc := ic.ProgramCounter
		reader := bufio.NewReader(input)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return true
		}
		text = strings.TrimSpace(text)
		val, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println(err)
			return true
		}

		ic.WritePtr(pc+1, val)
		ic.ProgramCounter += 2
		return false
	}

}
func (ic *Intcode) Uniary(reader reader, f func(int)) executor {
	return func() bool {
		pc := ic.ProgramCounter
		i1 := reader(pc + 1)

		f(i1)

		ic.ProgramCounter += 2
		return false

	}
}

func (ic *Intcode) output(val int) {
	fmt.Fprintf(ic.out, "Output: %d\n", val)
}

func add2(r1, r2 int) int {
	return r1 + r2
}
func mul2(r1, r2 int) int {
	return r1 * r2
}
func less(r1, r2 int) int {
	if r1 < r2 {
		return 1
	}
	return 0
}
func equal(r1, r2 int) int {
	if r1 == r2 {
		return 1
	}
	return 0
}
