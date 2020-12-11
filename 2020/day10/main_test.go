package main

import (
	"sort"
	"testing"
)

var (
	testSmall = []int{
		16,
		10,
		15,
		5,
		1,
		11,
		7,
		19,
		6,
		12,
		4,
	}
	testBig = []int{
		28,
		33,
		18,
		42,
		31,
		14,
		46,
		20,
		48,
		47,
		24,
		23,
		49,
		45,
		19,
		38,
		39,
		11,
		1,
		32,
		25,
		35,
		8,
		17,
		7,
		9,
		4,
		2,
		34,
		10,
		3,
	}
)

func Test_diff(t *testing.T) {
	type args struct {
		l []int
	}
	tests := []struct {
		name  string
		args  args
		want1 int
		want3 int
	}{
		{
			name: "small",
			args: args{
				l: testSmall,
			},
			want1: 7,
			want3: 5,
		},
		{
			name: "big",
			args: args{
				l: testBig,
			},
			want1: 22,
			want3: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.l = append(tt.args.l, 0)
			sort.Ints(tt.args.l)
			got := count(diff(tt.args.l))
			if got[1] != tt.want1 {
				t.Errorf("count of 1s = %v, want %v", got[1], tt.want1)
			}
			if got[3]+1 != tt.want3 {
				t.Errorf("count of 3s = %v, want %v", got[3]+1, tt.want3)
			}
		})
	}
}

func Test_paths(t *testing.T) {
	type args struct {
		l []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "small",
			args: args{
				l: testSmall,
			},
			want: 8,
		},
		{
			name: "big",
			args: args{
				l: testBig,
			},
			want: 19208,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.l = append(tt.args.l, 0)
			sort.Ints(tt.args.l)
			if got := paths(tt.args.l); got != tt.want {
				t.Errorf("paths() = %v, want %v", got, tt.want)
			}
		})
	}
}
