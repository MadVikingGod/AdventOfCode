package main

import (
	"testing"
)

func Test_countIncreases(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "default",
			args: args{list: []int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countIncreases(tt.args.list); got != tt.want {
				t.Errorf("countIncreases() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countWindowIncreases(t *testing.T) {
	type args struct {
		list       []int
		windowSize int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "default",
			args: args{
				list: []int{
					199,
					200,
					208,
					210,
					200,
					207,
					240,
					269,
					260,
					263,
				},
				windowSize: 3,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countWindowIncreases(tt.args.list, tt.args.windowSize); got != tt.want {
				t.Errorf("countWindowIncreases() = %v, want %v", got, tt.want)
			}
		})
	}
}
