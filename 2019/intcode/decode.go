package intcode

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

// Note: currently instead of decoding the different modes at runtime they are decoded here
// A potentially useful change would be instead of registering each opcode with a different decode step
// put that common logic into the ic.decode.  Then each step would need to know how many parameters it had, and which can be immediate

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

func (ic *Intcode) Uniary(reader reader, f func(int)) executor {
	return func() bool {
		pc := ic.ProgramCounter
		i1 := reader(pc + 1)

		f(i1)

		ic.ProgramCounter += 2
		return false

	}
}
