package main

import (
	"log"
	"os"
	"runtime/pprof"
	"testing"

	"github.com/MadVikingGod/AdventOfCode/2018/helpers"
)

func TestRecurse(t *testing.T) {
	inputs, err := helpers.GetInput(5)
	if err != nil {
		log.Panic(err)
	}

	input := []byte(inputs[0])
	expected := iterative(input)
	got := recuse(input)
	if expected != got {
		t.Errorf("recuse got the wrong length: expected %d, got %d", expected, got)
	}
}

func BenchmarkIterative(b *testing.B) {
	f, _ := os.Create("iterative.proto.gz")
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	// run the Fib function b.N times
	inputs, err := helpers.GetInput(5)
	if err != nil {
		log.Panic(err)
	}

	input := []byte(inputs[0])
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		iterative(input)
	}
}

func BenchmarkRecursie(b *testing.B) {
	f, _ := os.Create("recursive.proto.gz")
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	// run the Fib function b.N times
	inputs, err := helpers.GetInput(5)
	if err != nil {
		log.Panic(err)
	}

	input := []byte(inputs[0])
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		recuse(input)
	}
}
