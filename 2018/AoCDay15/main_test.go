package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func Test_findElfPower(t *testing.T) {
	t.SkipNow()
	tests := []struct {
		name      string
		inputfile string
		want      int
		want1     int
		want2     int
	}{
		{
			name:      "demo",
			inputfile: "test_files/input.attack",
			want:      15,
			want1:     29,
			want2:     172,
		},
		{
			name:      "input2",
			inputfile: "test_files/input2",
			want:      4,
			want1:     33,
			want2:     948,
		},
		{
			name:      "input3",
			inputfile: "test_files/input3",
			want:      15,
			want1:     37,
			want2:     94,
		},
		{
			name:      "input4",
			inputfile: "test_files/input4",
			want:      12,
			want1:     39,
			want2:     166,
		},
		{
			name:      "input5",
			inputfile: "test_files/input5",
			want:      34,
			want1:     30,
			want2:     38,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf, err := ioutil.ReadFile(tt.inputfile)
			if err != nil {
				t.Error("error reading file: ", err)
			}
			inputs := strings.Split(string(buf), "\n")

			gm := readMap(inputs)

			got, got1, got2 := findElfPower(gm)
			if got != tt.want {
				t.Errorf("findElfPower() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findElfPower() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("findElfPower() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
