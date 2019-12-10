package main

import (
	"github.com/madvikinggod/AdventOfCode/2019/location"
	"reflect"
	"strings"
	"testing"
)

func Test_bestLocation(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name  string
		args  args
		want  location.Location
		want1 int
	}{
		{
			name: "simple",
			args: args{
				input: `.#..#
.....
#####
....#
...##`,
			},
			want:  location.New(3, 4),
			want1: 8,
		},
		{
			name: "example 1",
			args: args{
				input: `......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`,
			},
			want:  location.New(5, 8),
			want1: 33,
		},
		{
			name: "example 2",
			args: args{
				input: `#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`,
			},
			want:  location.New(1, 2),
			want1: 35,
		},
		{
			name: "example 3",
			args: args{
				input: `.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`,
			},
			want:  location.New(6, 3),
			want1: 41,
		},
		{
			name: "example 4",
			args: args{
				input: `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`,
			},
			want:  location.New(11, 13),
			want1: 210,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			space := map[location.Location]int{}

			for y, line := range strings.Split(tt.args.input, "\n") {
				for x, c := range line {
					if c == '#' {
						space[location.New(x, y)] = 0
					}
				}
			}

			got, got1 := bestLocation(space)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bestLocation() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("bestLocation() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
