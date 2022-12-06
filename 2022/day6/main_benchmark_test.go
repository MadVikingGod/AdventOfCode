package main

import (
	"testing"
)

func Benchmark_findStartOfMessage(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		findStartOfMessage(input)
	}
}

func Benchmark_findStart(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		findStart(input)
	}
}
