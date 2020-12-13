package main

import (
	"strconv"
	"testing"
)

func TestLeft(t *testing.T) {

	type args struct {
		d int
	}
	tests := []struct {
		name string
		args args
		want point
	}{
		{
			name: "90",
			args: args{
				d: 90,
			},
			want: point{-5, 1},
		},
		{
			name: "180",
			args: args{
				d: 180,
			},
			want: point{-1, -5},
		},
		{
			name: "270",
			args: args{
				d: 270,
			},
			want: point{5, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ship{dir: point{1, 5}}
			got := Left(tt.args.d)
			got(s)
			if s.dir != tt.want {
				t.Errorf("Left() = %v, want %v", s.dir, tt.want)
			}
		})
	}
}

func TestRight(t *testing.T) {
	type args struct {
		d int
	}
	tests := []struct {
		name string
		args args
		want point
	}{
		{
			name: "270",
			args: args{
				d: 270,
			},
			want: point{-5, 1},
		},
		{
			name: "180",
			args: args{
				d: 180,
			},
			want: point{-1, -5},
		},
		{
			name: "90",
			args: args{
				d: 90,
			},
			want: point{5, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ship{dir: point{1, 5}}
			got := Right(tt.args.d)
			got(s)
			if s.dir != tt.want {
				t.Errorf("RotateRight() = %v, want %v", s.dir, tt.want)
			}
		})
	}
}

func TestLeftAll(t *testing.T) {
	tests := []struct {
		name    string
		start   point
		ammount int
		stop    point
	}{
		{
			name:    "N90",
			start:   north,
			ammount: 90,
			stop:    west,
		},
		{
			name:    "N180",
			start:   north,
			ammount: 180,
			stop:    south,
		},
		{
			name:    "N270",
			start:   north,
			ammount: 270,
			stop:    east,
		},
		{
			name:    "E90",
			start:   east,
			ammount: 90,
			stop:    north,
		},
		{
			name:    "E180",
			start:   east,
			ammount: 180,
			stop:    west,
		},
		{
			name:    "E270",
			start:   east,
			ammount: 270,
			stop:    south,
		},
		{
			name:    "S90",
			start:   south,
			ammount: 90,
			stop:    east,
		},
		{
			name:    "S180",
			start:   south,
			ammount: 180,
			stop:    north,
		},
		{
			name:    "S270",
			start:   south,
			ammount: 270,
			stop:    west,
		},
		{
			name:    "W90",
			start:   west,
			ammount: 90,
			stop:    south,
		},
		{
			name:    "W180",
			start:   west,
			ammount: 180,
			stop:    east,
		},
		{
			name:    "W270",
			start:   west,
			ammount: 270,
			stop:    north,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ship{dir: tt.start}
			f := Left(tt.ammount)
			f(s)
			if s.dir != tt.stop {
				t.Errorf("Left() = %v, want %v", s.dir, tt.stop)
			}
		})
	}
}

func TestShip(t *testing.T) {
	type testcase struct {
		instruction string
		ship        ship
	}
	input := []testcase{
		{"F10", ship{
			dir:       east,
			location:  point{10, 0},
			waypoint:  point{10, 1},
			location2: point{100, 10}},
		},
		{"N3", ship{
			dir:       east,
			location:  point{10, 3},
			waypoint:  point{10, 4},
			location2: point{100, 10}},
		},
		{"F7", ship{
			dir:       east,
			location:  point{17, 3},
			waypoint:  point{10, 4},
			location2: point{170, 38}},
		},
		{"R90", ship{
			dir:       south,
			location:  point{17, 3},
			waypoint:  point{4, -10},
			location2: point{170, 38}},
		},
		{"F11", ship{
			dir:       south,
			location:  point{17, -8},
			waypoint:  point{4, -10},
			location2: point{214, -72}},
		},
	}
	s := &ship{
		dir: east,

		waypoint: point{10, 1},
	}

	for _, tc := range input {
		x, _ := strconv.Atoi(tc.instruction[1:])
		funcs[tc.instruction[0]](x)(s)

		if s.location != tc.ship.location || s.dir != tc.ship.dir {
			t.Errorf("Ship not at the correct destination in part1 for instruction %s, got = %v, want = %v", tc.instruction, s, tc.ship)
		}
		if s.location2 != tc.ship.location2 || s.waypoint != tc.ship.waypoint {
			t.Errorf("Ship not at the correct destination in part2 for instruction %s, got = %v, want = %v", tc.instruction, s, tc.ship)
		}
	}

	if dist := s.location2.Distance(point{}); dist != 286 {
		t.Errorf("ship ended at the wong location, got %d, want 286", dist)
	}
}
