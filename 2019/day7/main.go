package main

import (
	"fmt"
	"io"
	// "strconv"
	"strings"
	"sync"

	"github.com/madvikinggod/AdventOfCode/2019/intcode"
)

type logger struct {
	path string
}

func (l *logger) Write(in []byte) (int, error) {
	fmt.Printf("log %s: %+q\n", l.path, string(in))
	return 0, nil
}

func main() {

	// max := 0

	// for a := 0; a < 5; a++ {
	// 	for b := 0; b < 5; b++ {
	// 		for c := 0; c < 5; c++ {
	// 			for d := 0; d < 5; d++ {
	// 				for e := 0; e < 5; e++ {

	// 					if a == b || a == c || a == d || a == e ||
	// 						b == c || b == d || b == e || c == d ||
	// 						c == e || d == e {
	// 						continue
	// 					}

	// 					o, _ := amp(strconv.Itoa(a), "0")

	// 					o, _ = amp(strconv.Itoa(b), o)

	// 					o, _ = amp(strconv.Itoa(c), o)

	// 					o, _ = amp(strconv.Itoa(d), o)

	// 					o, _ = amp(strconv.Itoa(e), o)

	// 					x, _ := strconv.Atoi(o)
	// 					if x > max {
	// 						max = x
	// 						// fmt.Println(a,b,c,d,e)
	// 					}

	// 				}
	// 			}
	// 		}
	// 	}
	// }

	// fmt.Println(max)

	// a, b, c, d, e := 9, 8, 7, 6, 5
	a, b, c, d, e := 4,3,2,1,0

	buffaRead, buffaWrite := io.Pipe()
	buffbRead, buffbWrite := io.Pipe()
	buffcRead, buffcWrite := io.Pipe()
	buffdRead, buffdWrite := io.Pipe()
	buffeRead, buffeWrite := io.Pipe()

	buff := &strings.Builder{}

	outaWriter := io.MultiWriter(buffaWrite, &logger{"a"})
	outbWriter := io.MultiWriter(buffbWrite, &logger{"b"})
	outcWriter := io.MultiWriter(buffcWrite, &logger{"c"})
	outdWriter := io.MultiWriter(buffdWrite, &logger{"d"})
	outeWriter := io.MultiWriter(buff, buffeWrite, &logger{"e"})

	wg := &sync.WaitGroup{}
	wg.Add(5)

	go feedbackAmp(buffeRead, outaWriter, wg)
	go feedbackAmp(buffaRead, outbWriter, wg)
	go feedbackAmp(buffbRead, outcWriter, wg)
	go feedbackAmp(buffcRead, outdWriter, wg)
	go feedbackAmp(buffdRead, outeWriter, wg)

	fmt.Fprintf(outdWriter, "%d\n", e)
	fmt.Fprintf(outcWriter, "%d\n", d)
	fmt.Fprintf(outbWriter, "%d\n", c)
	fmt.Fprintf(outaWriter, "%d\n", b)
	fmt.Fprintf(outeWriter, "%d\n0\n", a)


}

func feedbackAmp(read io.Reader, write io.Writer, wg *sync.WaitGroup) {

	mem := make([]int, len(input))
	copy(mem, input)
	ic := intcode.Intcode{
		Memory: mem,
		In:     read,
		Out:    write,
	}
	ic.Register()

	ic.Run()

	wg.Done()
}

func amp(setting, in string) (string, error) {
	setting = strings.TrimSpace(setting)
	in = strings.TrimSpace(in)

	memory := make([]int, len(input))
	copy(memory, input)

	buff := &strings.Builder{}

	ic := intcode.Intcode{
		Memory: memory,
		In:     strings.NewReader(setting + "\n" + in + "\n"),
		Out:    buff,
	}
	ic.Register()

	_, err := ic.Run()

	return strings.TrimSpace(buff.String()), err

}
