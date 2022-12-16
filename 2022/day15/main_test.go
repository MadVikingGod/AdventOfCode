package main

import "testing"

func BenchmarkPart1(b *testing.B) {

	sensors := parseSensors(input)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		CountFull(2000000, sensors)
	}
}

func BenchmarkPart2(b *testing.B) {
	sensors := parseSensors(input)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		findBeacon(4000000, sensors)
	}
}
