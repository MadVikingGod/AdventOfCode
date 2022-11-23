package main

import (
	_ "embed"
	"testing"
)

func Test_findSumNumber(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "[1,2,3]",
			want: 6,
		},
		{
			name: `{"a":2,"b":4}`,
			want: 6,
		},
		{
			name: `[[[3]]]`,
			want: 3,
		},
		{
			name: `{"a":{"b":4},"c":-1}`,
			want: 3,
		},
		{
			name: `{"a":[-1,1]}`,
			want: 0,
		},
		{
			name: `[-1,{"a":1}]`,
			want: 0,
		},
		{
			name: `[]`,
			want: 0,
		},
		{
			name: `{}`,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSumNumber(marshal(tt.name)); got != tt.want {
				t.Errorf("findSumNumber(%s) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

// [1,2,3] still has a sum of 6.
// [1,{"c":"red","b":2},3] now has a sum of 4, because the middle object is ignored.
// {"d":"red","e":[1,2,3,4],"f":5} now has a sum of 0, because the entire structure is ignored.
// [1,"red",5] has a sum of 6, because "red" in an array has no effect.

func Test_findSumIgnoreRed(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: `[1,2,3]`,
			want: 6,
		},
		{
			name: `[1,{"c":"red","b":2},3]`,
			want: 4,
		},
		{
			name: `{"d":"red","e":[1,2,3,4],"f":5}`,
			want: 0,
		},
		{
			name: `[1,"red",5]`,
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSumIgnoreRed(marshal(tt.name)); got != tt.want {
				t.Errorf("findSumIgnoreRed() = %v, want %v", got, tt.want)
			}
		})
	}
}
