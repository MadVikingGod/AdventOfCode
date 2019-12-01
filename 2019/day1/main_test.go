package main

import (
	"strconv"
	"testing"
)

func Test_fuelToLaunch(t *testing.T) {

	tests := []struct {
		i    int
		want int
	}{
		{
			i:    12,
			want: 2,
		},
		{
			i:    14,
			want: 2,
		},
		{
			i:    1969,
			want: 654,
		},
		{
			i:    100756,
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

func Test_additionalFuel(t *testing.T) {

	tests := []struct {
		start int
		want int
	}{
		{
			start: 2,
			want: 0,
		},
		{
			start: 654,
			want: 312,
		},
		{
			start: 33583,
			want: 16763,
		},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.start), func(t *testing.T) {
			if got := additionalFuel(tt.start); got != tt.want {
				t.Errorf("additionalFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}
