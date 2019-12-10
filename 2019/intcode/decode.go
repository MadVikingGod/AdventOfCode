package intcode

type reader func(int) int

func (ic Intcode) Read(position int) int {
	return ic.Memory[position]
}
func (ic Intcode) ReadPtr(position int) int {
	return ic.Memory[ic.Memory[position]]
}
func (ic *Intcode) ReadBase(position int) int {
	address := ic.Base + ic.Memory[position]
	return ic.Memory[address]
}

type writer func(int, int)

func (ic *Intcode) Write(position, data int) {
	ic.Memory[position] = data
}
func (ic *Intcode) WritePtr(position, data int) {
	ic.Memory[ic.Memory[position]] = data
}
func (ic *Intcode) WriteBase(position, data int) {
	address := ic.Base + ic.Memory[position]
	ic.Memory[address] = data
}

// Note: currently instead of decoding the different modes at runtime they are decoded here
// A potentially useful change would be instead of registering each opcode with a different decode step
// put that common logic into the ic.decode.  Then each step would need to know how many parameters it had, and which can be immediate

func (ic *Intcode) registerTrinary(opcode int, f func(int, int) int) {

	for i, read1 := range []reader{ic.ReadPtr, ic.Read, ic.ReadBase} {
		for j, read2 := range []reader{ic.ReadPtr, ic.Read, ic.ReadBase} {
			for k, write3 := range []writer{ic.WritePtr, ic.Write, ic.WriteBase} {
				op := k*10000 + j*1000 + i*100 + opcode
				ic.opcodes[op] = ic.trinary(read1, read2, write3, f)

			}
		}
	}

}
func (ic *Intcode) registerBinary(opcode int, f func(reader, reader) executor) {

	for i, read1 := range []reader{ic.ReadPtr, ic.Read, ic.ReadBase} {
		for j, read2 := range []reader{ic.ReadPtr, ic.Read, ic.ReadBase} {
			op := +j*1000 + i*100 + opcode
			ic.opcodes[op] = f(read1, read2)
		}
	}

}

func (ic *Intcode) trinary(reader1, reader2 reader, writer writer, f func(int, int) int) executor {
	return func() bool {

		pc := ic.ProgramCounter
		i1 := reader1(pc + 1)
		i2 := reader2(pc + 2)

		out := f(i1, i2)

		writer(pc+3, out)
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
