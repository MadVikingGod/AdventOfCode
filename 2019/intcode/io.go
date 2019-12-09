package intcode

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func (ic *Intcode) store(input io.Reader) executor {
	if ic.buffer == nil {
		ic.buffer = []string{}
	}
	scanner := bufio.NewScanner(input)
	return func() bool {
		if !scanner.Scan() {
			fmt.Println(ic.Tag, "Input was closed before receiving")
			return true
		}
		text := scanner.Text()
		text = strings.TrimSpace(text)

		val, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println(err)
			return true
		}

		pc := ic.ProgramCounter
		ic.WritePtr(pc+1, val)
		ic.ProgramCounter += 2
		return false
	}

}

func (ic *Intcode) output(val int) {
	//fmt.Println("Writing output:", ic.Tag, val)
	fmt.Fprintf(ic.Out, "%d\n", val)
}
