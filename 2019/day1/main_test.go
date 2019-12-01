package main

import (
	"strconv"
	"testing"
)

func Test_fuelToLaunch(t *testing.T) {

	tests := []struct {
		i int
		want int
	}{
		{
			i: 12,
			want: 2,
		},
		{
			i: 14,
			want: 2,
		},
		{
			i: 1969,
			want: 654,
		},
		{
			i: 100756,
			want: 33583,
		},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.i), func(t *testing.T) {
			if got := fuelToLaunch(tt.i); got != tt.want {
				t.Errorf("fuelToLaunch() = %v, want %v", got, tt.want)
			}
		})
	}
}
