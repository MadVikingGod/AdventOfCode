package intcode

import "fmt"

import "bufio"

import "os"

import "strconv"

import "strings"

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
	// fmt.Println(ic.Memory[0:10])
	halt, err := ic.decode()
	for !halt && err == nil {
		// pc := ic.ProgramCounter
		// max := pc+10
		// if max > len(ic.Memory) {
		// 	max = len(ic.Memory)
		// }
		// fmt.Println(ic.Memory[pc:max])
	
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

	ic.opcodes[1] = ic.trinary(ic.ReadPtr, ic.ReadPtr, add2)
	ic.opcodes[101] = ic.trinary(ic.Read, ic.ReadPtr, add2)
	ic.opcodes[1001] = ic.trinary(ic.ReadPtr, ic.Read, add2)
	ic.opcodes[1101] = ic.trinary(ic.Read, ic.Read, add2)
	ic.opcodes[2] = ic.trinary(ic.ReadPtr, ic.ReadPtr, mul2)
	ic.opcodes[102] = ic.trinary(ic.Read, ic.ReadPtr, mul2)
	ic.opcodes[1002] = ic.trinary(ic.ReadPtr, ic.Read, mul2)
	ic.opcodes[1102] = ic.trinary(ic.Read, ic.Read, mul2)

	ic.opcodes[5] =    ic.jmpeq(ic.ReadPtr, ic.ReadPtr)
	ic.opcodes[105] =  ic.jmpeq( ic.Read, ic.ReadPtr)
	ic.opcodes[1005] = ic.jmpeq(ic.ReadPtr, ic.Read)
	ic.opcodes[1105] = ic.jmpeq(ic.Read, ic.Read)
	ic.opcodes[6] =  ic.jmpneq(ic.ReadPtr, ic.ReadPtr)
	ic.opcodes[106] =  ic.jmpneq( ic.Read, ic.ReadPtr)
	ic.opcodes[1006] =  ic.jmpneq(ic.ReadPtr, ic.Read)
	ic.opcodes[1106] =  ic.jmpneq(ic.Read, ic.Read)
	

	ic.opcodes[7] = ic.trinary(ic.ReadPtr, ic.ReadPtr, less)
	ic.opcodes[107] = ic.trinary(ic.Read, ic.ReadPtr, less)
	ic.opcodes[1007] = ic.trinary(ic.ReadPtr, ic.Read, less)
	ic.opcodes[1107] = ic.trinary(ic.Read, ic.Read, less)
	ic.opcodes[8] = ic.trinary(ic.ReadPtr, ic.ReadPtr, equal)
	ic.opcodes[108] = ic.trinary(ic.Read, ic.ReadPtr, equal)
	ic.opcodes[1008] = ic.trinary(ic.ReadPtr, ic.Read, equal)
	ic.opcodes[1108] = ic.trinary(ic.Read, ic.Read, equal)
	

	ic.opcodes[3] = ic.store
	ic.opcodes[4] = ic.Uniary(ic.ReadPtr, output)
	ic.opcodes[104] = ic.Uniary(ic.Read, output)


	ic.opcodes[99] = executor(func() bool {
		return true
	})
}

func (ic *Intcode) trinary(reader1, reader2 func(int) int, f func(int, int) int) executor {
	return func() bool {
		
		pc := ic.ProgramCounter
		i1 := reader1(pc+1)
		i2 := reader2(pc+2)
		

		out := f(i1, i2)

		ic.WritePtr(pc+3, out)
		ic.ProgramCounter += 4
		return false
	}
}

func (ic *Intcode) jmpeq(reader1, reader2 func(int) int) executor {
	return func() bool {
		pc := ic.ProgramCounter
		val := reader1(pc+1)
		if val != 0 {
			dst := reader2(pc+2)
			ic.ProgramCounter = dst
			return false
		}
		ic.ProgramCounter +=3
		return false
	}
}
func (ic *Intcode) jmpneq(reader1, reader2 func(int) int) executor {
	return func() bool {
		pc := ic.ProgramCounter
		val := reader1(pc+1)
		if val == 0 {
			dst := reader2(pc+2)
			ic.ProgramCounter = dst
			return false
		}
		ic.ProgramCounter +=3
		return false
	}
}

func (ic *Intcode) store() bool {
	pc := ic.ProgramCounter
	reader := bufio.NewReader(os.Stdin)
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
func (ic *Intcode) Uniary(reader func(int)int, f func(int)) executor {
	return func() bool {
		pc := ic.ProgramCounter
		i1 := reader(pc+1)

		f(i1)

		ic.ProgramCounter +=2
		return false

	}
}

func output(val int) {
	fmt.Println("Output: ", val)	
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