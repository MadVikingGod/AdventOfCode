package intcode

import "fmt"

type executor func() bool
type Intcode struct {
	Memory         []int
	ProgramCounter int

	opcodes map[int]executor

	//This is just noodeling for registers that may be addded
	registers [10]int
}

func New(memory []int) *Intcode {
	ic := &Intcode{
		Memory: append([]int(nil), memory...),
	}
	ic.register()
	return ic
}

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

	ic.opcodes[1] = ic.binary(add2)
	ic.opcodes[2] = ic.binary(mul2)
	ic.opcodes[99] = executor(func() bool {
		return true
	})
}

func (ic *Intcode) binary(f func(int, int) int) executor {
	return func() bool {
		pc := ic.ProgramCounter
		imm1 := ic.ReadPtr(pc + 1)
		imm2 := ic.ReadPtr(pc + 2)

		out := f(imm1, imm2)

		ic.WritePtr(pc+3, out)
		ic.ProgramCounter += 4
		return false
	}
}

func add2(r1, r2 int) int {
	return r1 + r2
}
func mul2(r1, r2 int) int {
	return r1 * r2
}
