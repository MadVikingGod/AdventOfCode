package intcode

import "fmt"

type Intcode struct {
	Memory         []int
	ProgramCounter int
}

func New(memory []int) *Intcode {
	return &Intcode{
		Memory: append([]int(nil), memory...),
	}
}

func (ic Intcode) Read(position int) int {
	return ic.Memory[position]
}
func (ic *Intcode) Write(position, data int) {
	ic.Memory[position] = data
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
	switch ic.Memory[ic.ProgramCounter] {
	case 1:
		ic.add()
		return false, nil
	case 2:
		ic.mul()
		return false, nil
	case 99:
		return true, nil
	default:
		return false, fmt.Errorf("not implemented, opcode %d at %d", ic.Memory[ic.ProgramCounter], ic.ProgramCounter)
	}
}

// add() returns [pc+1] + [pc+2] into [pc+3]
func (ic *Intcode) add() {
	pc := ic.ProgramCounter
	R1 := ic.Read(pc + 1)
	R2 := ic.Read(pc + 2)
	dest := ic.Read(pc + 3)
	ic.Write(dest, ic.Read(R1)+ic.Read(R2))
	ic.ProgramCounter += 4
}

func (ic *Intcode) mul() {
	pc := ic.ProgramCounter
	R1 := ic.Read(pc + 1)
	R2 := ic.Read(pc + 2)
	dest := ic.Read(pc + 3)
	ic.Write(dest, ic.Read(R1)*ic.Read(R2))
	ic.ProgramCounter += 4
}
