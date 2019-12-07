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
	return func() bool {
		var text string
		if len(ic.buffer) >0 {
			text = ic.buffer[0]
			ic.buffer = ic.buffer[1:]
		} else {
			fmt.Println("Reading input")
			scanner := bufio.NewScanner(input)
			if !scanner.Scan() {
				fmt.Println("No data")
				return true
			}
			text = scanner.Text()
			for scanner.Scan() {
				ic.buffer = append(ic.buffer, scanner.Text())
			}
			
		}

		pc := ic.ProgramCounter
				
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
	fmt.Fprintf(ic.Out, "%d\n", val)
}
