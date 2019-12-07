package intcode

import (
	"fmt"
	"io"
	"os"
)

type executor func() bool
type Intcode struct {
	Memory         []int
	ProgramCounter int

	opcodes map[int]executor

	In  io.Reader
	buffer []string
	Out io.Writer
}

func New(memory []int) *Intcode {
	ic := &Intcode{
		Memory: append([]int(nil), memory...),
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

	ic.opcodes[3] = ic.store(ic.In)
	ic.opcodes[4] = ic.Uniary(ic.ReadPtr, ic.output)
	ic.opcodes[104] = ic.Uniary(ic.Read, ic.output)

	ic.opcodes[99] = executor(func() bool {
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
