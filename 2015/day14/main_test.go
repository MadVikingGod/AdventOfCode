package main

import (
	_ "embed"
	"testing"
)

func Test_score(t *testing.T) {
	type args struct {
		reindeers   []reindeer
		maxDistance int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				reindeers: []reindeer{
					{"Comet", 14, 10, 127},
					{"Dancer", 16, 11, 162},
				},
				maxDistance: 1000,
			},
			want: 689,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := score(tt.args.reindeers, tt.args.maxDistance); got != tt.want {
				t.Errorf("score() = %v, want %v", got, tt.want)
			}
		})
	}
}
