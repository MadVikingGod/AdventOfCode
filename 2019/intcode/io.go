package intcode

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

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

func (ic *Intcode) output(val int) {
	fmt.Fprintf(ic.out, "Output: %d\n", val)
}
