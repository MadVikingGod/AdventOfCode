package main

import (
	"sync"

	"bytes"
	"fmt"
	"log"

	"github.com/madvikinggod/AdventOfCode/2018/helpers"
)

func main() {
	inputs, err := helpers.GetInput(7)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(len(inputs))

	deps := map[byte][]byte{}
	for _, line := range inputs {
		a, b := getDependancy(line)
		deps[b] = append(deps[b], a)
	}
	for i := byte('A'); i <= 'Z'; i++ {
		fmt.Println(string(i), string(deps[i]))
	}
	done := map[byte]bool{}

	buf := bytes.NewBufferString("")

	fmt.Printf("---\n---\n")

	for j := 0; j < 26; j++ {
		for i := byte('A'); i <= 'Z'; i++ {
			if !done[i] && isDone(i, done, deps) {
				done[i] = true
				buf.WriteByte(i)
				fmt.Println(string(i), string(deps[i]))
				break
			}
		}
	}

	fmt.Println(buf.String())
	fmt.Println(buf.Len())

	w := NewWorkers(5)
	done = map[byte]bool{0: true}
	added := map[byte]bool{}
	buf = bytes.NewBufferString("")
	tick := 0
	for buf.Len() < 26 {
		for i, time := range w.times {
			if time == 0 && !done[w.chars[i]] {
				//fmt.Printf("%4d: Finished %s\n", tick, string(w.chars[i]))
				done[w.chars[i]] = true
				buf.WriteByte(w.chars[i])
			}
		}
		for i := byte('A'); i <= 'Z'; i++ {
			if !added[i] && isDone(i, done, deps) {
				if w.add(i, 60+int(i)-int('A')+1) {
					added[i] = true

					fmt.Printf("%4d: Added %s\n", tick, string(i))
				}
			}
		}
		tick = tick + w.next()
		w.subtract(w.next())
	}
	fmt.Println(tick)
}

type workers struct {
	times []int
	chars []byte
}

func NewWorkers(count int) workers {
	return workers{
		times: make([]int, count),
		chars: make([]byte, count),
	}
}
func (w workers) add(char byte, time int) bool {
	for i, t := range w.times {
		if t <= 0 {
			w.times[i] = time
			w.chars[i] = char
			return true
		}
	}
	return false
}
func (w workers) next() int {
	count := 100
	for _, t := range w.times {
		if t > 0 && t < count {
			count = t
		}
	}
	if count == 100 {
		return 0
	}
	return count
}
func (w workers) subtract(count int) {
	for i := range w.times {
		if w.times[i] > 0 {
			w.times[i] = w.times[i] - count
		}
	}
}

func broadcast(input chan struct{}, out ...chan struct{}) {
	for {
		if _, ok := <-input; !ok {
			for _, c := range out {
				close(c)
			}
			return
		}
		for _, c := range out {
			select {
			case c <- struct{}{}:
			default:
			}
		}
	}
}

func worker(input <-chan byte, ticks <-chan struct{}, out chan<- byte, wg sync.WaitGroup) {
	for {
		c, ok := <-input
		if !ok {
			wg.Done()
			return
		}
		count := 60 + c
		for count > 0 {
			<-ticks
			count--
		}
		out <- c
	}
}

func isDone(uut byte, done map[byte]bool, deps map[byte][]byte) bool {
	if done[uut] {
		return true
	}
	for _, x := range deps[uut] {
		if !done[x] {
			return false
		}
	}
	return true
}
func getDependancy(input string) (byte, byte) {

	return input[5], input[36]
}
