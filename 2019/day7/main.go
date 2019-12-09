package main

import (
	"fmt"
	"io"
	"strconv"

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

	max := 0

	for a := 5; a < 10; a++ {
		for b := 5; b < 10; b++ {
			for c := 5; c < 10; c++ {
				for d := 5; d < 10; d++ {
					for e := 5; e < 10; e++ {

						if a == b || a == c || a == d || a == e ||
							b == c || b == d || b == e || c == d ||
							c == e || d == e {
							continue
						}

						buffaRead, buffaWrite := io.Pipe()
						buffbRead, buffbWrite := io.Pipe()
						buffcRead, buffcWrite := io.Pipe()
						buffdRead, buffdWrite := io.Pipe()
						buffeRead, buffeWrite := io.Pipe()

						buff := &strings.Builder{}

						//outeWriter := io.MultiWriter(buff,  &logger{"e"})
						outeWriter := io.MultiWriter(buff, buffeWrite)

						wg := &sync.WaitGroup{}
						wg.Add(4)
						wg2 := &sync.WaitGroup{}
						wg2.Add(1)

						//inaReader := io.MultiReader(strings.NewReader(fmt.Sprintf("%d\n0\n", a)))

						inaReader := io.MultiReader(strings.NewReader(fmt.Sprintf("%d\n0\n", a)), buffeRead)
						inbReader := io.MultiReader(strings.NewReader(fmt.Sprintf("%d\n", b)), buffaRead)
						incReader := io.MultiReader(strings.NewReader(fmt.Sprintf("%d\n", c)), buffbRead)
						indReader := io.MultiReader(strings.NewReader(fmt.Sprintf("%d\n", d)), buffcRead)
						ineReader := io.MultiReader(strings.NewReader(fmt.Sprintf("%d\n", e)), buffdRead)

						go feedbackAmp(inaReader, buffaWrite, wg, "a")
						go feedbackAmp(inbReader, buffbWrite, wg, "b")
						go feedbackAmp(incReader, buffcWrite, wg, "c")
						go feedbackAmp(indReader, buffdWrite, wg, "d")
						go feedbackAmp(ineReader, outeWriter, wg2, "e")
						wg.Wait()

						buffaWrite.Close()
						buffbWrite.Close()
						buffcWrite.Close()
						buffdWrite.Close()
						buffeWrite.Close()

						outputs := strings.Split(buff.String(), "\n")

						test, err := strconv.Atoi(outputs[len(outputs)-2])
						if err != nil {
							fmt.Println("ERROR:", a, b, c, d, err)
						}
						if test > max {
							fmt.Println("OUTPUTS", a, b, c, d, e, outputs[len(outputs)-2])
							max = test
						}

					}
				}
			}
		}
	}

	fmt.Println("MAX", max)
}

func feedbackAmp(read io.Reader, write io.Writer, wg *sync.WaitGroup, tag string) {

	mem := make([]int, len(input))
	copy(mem, input)
	ic := intcode.Intcode{
		Memory: mem,
		In:     read,
		Out:    write,
		Tag:    tag,
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

//4305190943
