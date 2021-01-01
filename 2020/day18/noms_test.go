package main

import "testing"

func TestParse(t *testing.T) {

	tests := []struct {
		line string
		want int
	}{
		{
			line: "1 + 2 * 3 + 4 * 5 + 6",
			want: 71,
		},
		{
			line: "1 + (2 * 3) + (4 * (5 + 6))",
			want: 51,
		},
	}
	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			if got := Parse(tt.line); got != tt.want {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
