package main

import (
	"testing"
)

var testIns = []string{
	"nop +0",
	"acc +1",
	"jmp +4",
	"acc +3",
	"jmp -3",
	"acc -99",
	"acc +1",
	"jmp -4",
	"acc +6",
}
var testInsTerminate = []string{
	"nop +0",
	"acc +1",
	"jmp +4",
	"acc +3",
	"jmp -3",
	"acc -99",
	"acc +1",
	"nop -4",
	"acc +6",
}

func TestComputer_Step(t *testing.T) {
	tests := []struct {
		name    string
		steps   int
		wantAcc int
		wantIp  int
	}{
		{
			name:    "nop",
			steps:   1,
			wantAcc: 0,
			wantIp:  1,
		},
		{
			name:    "acc",
			steps:   2,
			wantAcc: 1,
			wantIp:  2,
		},
		{
			name:    "jmp",
			steps:   3,
			wantAcc: 1,
			wantIp:  6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newComputer(testIns)
			for i := 0; i < tt.steps; i++ {
				c.Step()
			}
			if c.acc != tt.wantAcc {
				t.Errorf("comptuer acc != wantAcc, c.acc=%d, wantAcc=%d", c.acc, tt.wantAcc)
			}
			if c.ip != tt.wantIp {
				t.Errorf("comptuer ip != wantIp, c.ip=%d, wantIp=%d", c.ip, tt.wantIp)
			}

		})
	}
}

func Test_runTillLoop(t *testing.T) {

	c := newComputer(testIns)
	runTillLoop(c)
	if c.acc != 5 {
		t.Errorf("c.acc != 5, c.acc=%d", c.acc)
	}

}

func Test_doesTerminate(t *testing.T) {
	tests := []struct {
		name string
		ins  []string
		want bool
	}{
		{
			name: "doesn't",
			ins:  testIns,
			want: false,
		},
		{
			name: "does",
			ins:  testInsTerminate,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newComputer(tt.ins)
			if got := doesTerminate(c); got != tt.want {
				t.Errorf("doesTerminate() = %v, want %v", got, tt.want)
			}
		})
	}
}
