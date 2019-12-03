package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/madvikinggod/AdventOfCode/2018/helpers"
)

func main() {
	input, err := helpers.GetInput(4)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(len(input))
	guards := map[int][]int{}

	guard := 0
	sleep := 0

	for _, line := range input {
		if isGuard(line) {
			guard = guardNumber(line)
			if _, ok := guards[guard]; !ok {
				guards[guard] = make([]int, 60)
			}
		}
		if isAsleep(line) {
			sleep = time(line)
		}
		if isAwake(line) {
			for i := sleep; i < time(line); i++ {
				guards[guard][i]++
			}
		}
	}
	fmt.Println(guards)
	most := 0

	for g, times := range guards {
		if sum(times) > most {
			guard = g
			most = sum(times)
		}
	}
	fmt.Println(guard)
	fmt.Println(guard * max(guards[guard]))

	maxMinute := 0
	maxCount := 0
	for g, times := range guards {
		if times[max(times)] > maxCount {
			guard = g
			maxMinute = max(times)
			maxCount = times[maxMinute]
		}
	}
	fmt.Println(guard * maxMinute)

}

func sum(in []int) int {
	s := 0
	for _, i := range in {
		s = s + i
	}
	return s
}
func max(in []int) int {
	m := 0
	mv := 0
	for i, v := range in {
		if v > mv {
			m = i
			mv = v
		}
	}
	return m
}

// `[1518-02-24 23:58] Guard #2347 begins shift`
func time(l string) int {
	t, _ := strconv.Atoi(l[15:17])
	return t
}
func isGuard(l string) bool {
	return l[19:24] == "Guard"
}
func isAsleep(l string) bool {
	return l[19:24] == "falls"
}
func isAwake(l string) bool {
	return l[19:24] == "wakes"
}
func guardNumber(l string) int {
	f := strings.Fields(l)
	g, _ := strconv.Atoi(f[3][1:])
	return g
}
