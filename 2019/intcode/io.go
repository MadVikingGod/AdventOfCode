package intcode

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func (ic *Intcode) store(input io.Reader, writer writer) executor {
	if ic.scanner == nil {
		ic.scanner = bufio.NewScanner(input)
	}
	if ic.buffer == nil {
		ic.buffer = []string{}
	}

	return func() bool {
		if !ic.scanner.Scan() {
			fmt.Println(ic.Tag, "Input was closed before receiving")
			return true
		}
		text := ic.scanner.Text()
		text = strings.TrimSpace(text)

		val, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println(err)
			return true
		}

		pc := ic.ProgramCounter
		writer(pc+1, val)
		ic.ProgramCounter += 2
		return false
	}

}

func (ic *Intcode) output(val int) {
	// fmt.Println("Writing output:", ic.Tag, val)
	fmt.Fprintf(ic.Out, "%d\n", val)
}
