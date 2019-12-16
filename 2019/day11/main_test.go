package main

import (
	"github.com/madvikinggod/AdventOfCode/2019/location"
	"io"
	"strings"
	"sync"
	"testing"
)

func Test_robot_run_turn(t *testing.T) {
	type fields struct {
		in  io.Reader
		dir location.Location
	}
	tests := []struct {
		name    string
		fields  fields
		wantDir location.Location
	}{
		{
			name: "Up turn left",
			fields: fields{
				in:  strings.NewReader("0\n0\n"),
				dir: up,
			},
			wantDir: left,
		},
		{
			name: "Up turn right",
			fields: fields{
				in:  strings.NewReader("0\n1\n"),
				dir: up,
			},
			wantDir: right,
		},
		{
			name: "down turn left",
			fields: fields{
				in:  strings.NewReader("0\n0\n"),
				dir: down,
			},
			wantDir: right,
		},
		{
			name: "down turn right",
			fields: fields{
				in:  strings.NewReader("0\n1\n"),
				dir: down,
			},
			wantDir: left,
		},
		{
			name: "left turn left",
			fields: fields{
				in:  strings.NewReader("0\n0\n"),
				dir: left,
			},
			wantDir: down,
		},
		{
			name: "left turn right",
			fields: fields{
				in:  strings.NewReader("0\n1\n"),
				dir: left,
			},
			wantDir: up,
		},
		{
			name: "right turn left",
			fields: fields{
				in:  strings.NewReader("0\n0\n"),
				dir: right,
			},
			wantDir: up,
		},
		{
			name: "right turn right",
			fields: fields{
				in:  strings.NewReader("0\n1\n"),
				dir: right,
			},
			wantDir: down,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &robot{
				in:   tt.fields.in,
				out:  &strings.Builder{},
				dir:  tt.fields.dir,
				hull: map[location.Location]int{},
			}
			wg := &sync.WaitGroup{}
			wg.Add(1)
			r.run(wg)
			if tt.wantDir != r.dir {
				t.Errorf("Robot did not turn the correct direction got %v, want %v", r.dir, tt.wantDir)
			}
		})
	}
}

func Test_robot_run_output(t *testing.T) {
	type fields struct {
		in   io.Reader
		dir  location.Location
		hull map[location.Location]int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Up turn left",
			fields: fields{
				in:  strings.NewReader("0\n1\n0\n0\n0\n1\n0\n0\n0\n1\n0\n0\n0\n1\n0\n0\n1\n0\n1\n0\n1\n0\n1\n0\n"),
				dir: up,

				hull: map[location.Location]int{
					location.New(1, 0): 1,
					location.New(1, 1): 0,
					location.New(2, 1): 1,
					location.New(2, 2): 0,
					location.New(3, 2): 1,
					location.New(3, 3): 0,
					location.New(4, 3): 1,
					location.New(4, 4): 0,
				},
			},
			want: "0\n1\n0\n1\n0\n1\n0\n1\n0\n0\n0\n0\n1\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &strings.Builder{}
			r := &robot{
				in:   tt.fields.in,
				out:  b,
				dir:  tt.fields.dir,
				hull: tt.fields.hull,
			}
			wg := &sync.WaitGroup{}
			wg.Add(1)
			r.run(wg)
			if tt.want != b.String() {
				t.Errorf("Robot did output the correct paint got %+q, want %+q", b.String(), tt.want)
			}
		})
	}
}
