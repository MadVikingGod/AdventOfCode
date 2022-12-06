package main

import (
	_ "embed"
	"testing"
)

func Test_findStart(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{name: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", want: 7},
		{name: "bvwbjplbgvbhsrlpgdmjqwftvncz", want: 5},
		{name: "nppdvjthqldpwncqszvftbrmjlhg", want: 6},
		{name: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", want: 10},
		{name: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", want: 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findStart(tt.name); got != tt.want {
				t.Errorf("findStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findStartOfMessage(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{name: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", want: 19},
		{name: "bvwbjplbgvbhsrlpgdmjqwftvncz", want: 23},
		{name: "nppdvjthqldpwncqszvftbrmjlhg", want: 23},
		{name: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", want: 29},
		{name: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", want: 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findStartOfMessage(tt.name); got != tt.want {
				t.Errorf("findStartOfMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
