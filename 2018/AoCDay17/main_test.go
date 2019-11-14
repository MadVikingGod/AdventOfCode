package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func Test_getStedyState(t *testing.T) {
	tests := []struct {
		name      string
		inputfile string
	}{
		// TODO: Add test cases.
		{
			name:      "punchout",
			inputfile: "input_punchout",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf, err := ioutil.ReadFile(tt.inputfile)
			if err != nil {
				t.Fatal("could not load map: " + tt.inputfile)
			}
			inputs := strings.Split(string(buf), "\n")

			gm := NewGameMap(inputs)
			getStedyState(gm)
			fmt.Print(gm)
			fmt.Println(gm.CountWater())
		})
	}
}
