package intcode

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type executor func() bool
type Intcode struct {
	Memory         []int
	ProgramCounter int
	Base           int

	opcodes map[int]executor

	In      io.Reader
	buffer  []string
	Out     io.Writer
	scanner *bufio.Scanner

	Tag string
}

func New(memory []int) *Intcode {
	ic := &Intcode{
		Memory: append(append([]int(nil), memory...), make([]int, 4000)...),
		In:     os.Stdin,
		Out:    os.Stdout,
	}
	ic.Register()
	return ic
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

func (ic *Intcode) Register() {
	ic.scanner = nil

	if ic == nil {
		return
	}
	if ic.opcodes == nil {
		ic.opcodes = map[int]executor{}
	}

	ic.registerTrinary(1, add)
	ic.registerTrinary(2, mul)

	ic.registerBinary(5, ic.jmpneq)
	ic.registerBinary(6, ic.jmpeq)

	ic.registerTrinary(7, less)
	ic.registerTrinary(8, equal)

	ic.opcodes[3] = ic.store(ic.In, ic.WritePtr)
	ic.opcodes[203] = ic.store(ic.In, ic.WriteBase)
	ic.opcodes[4] = ic.Uniary(ic.ReadPtr, ic.output)
	ic.opcodes[104] = ic.Uniary(ic.Read, ic.output)
	ic.opcodes[204] = ic.Uniary(ic.ReadBase, ic.output)

	ic.opcodes[9] = ic.Uniary(ic.ReadPtr, ic.addBase)
	ic.opcodes[109] = ic.Uniary(ic.Read, ic.addBase)
	ic.opcodes[209] = ic.Uniary(ic.ReadBase, ic.addBase)

	ic.opcodes[99] = executor(func() bool {
		//fmt.Println("Stopping ", ic.Tag )
		return true
	})
}

// The Jump commands are special because they don't can modify the PC

func (ic *Intcode) jmpneq(reader1, reader2 reader) executor {
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
func (ic *Intcode) jmpeq(reader1, reader2 reader) executor {
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

func (ic *Intcode) addBase(val int) {
	ic.Base += val
}
