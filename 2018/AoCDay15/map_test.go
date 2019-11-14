package main

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/go-test/deep"
)

func Test_gameMap_attack(t *testing.T) {
	type fields struct {
		x       int
		y       int
		walls   map[location]mapSquare
		goblins map[location]mapSquare
		elfs    map[location]mapSquare
	}
	type args struct {
		loc location
		dmg int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "attack",
			fields: fields{
				elfs: map[location]mapSquare{
					location{1, 1}: mapSquare{hitpoints: 200},
				},
			},
			args: args{
				loc: location{1, 1},
				dmg: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &gameMap{
				x:       tt.fields.x,
				y:       tt.fields.y,
				walls:   tt.fields.walls,
				goblins: tt.fields.goblins,
				elfs:    tt.fields.elfs,
			}
			m.attack(tt.args.loc, tt.args.dmg)
			if m.elfs[location{1, 1}].hitpoints != 197 {
				t.Error("did not subtract attack")
			}
		})
	}
}

func Test_inputs(t *testing.T) {
	tests := []struct {
		name      string
		inputfile string
		n         int
		elfs      map[location]mapSquare
		goblins   map[location]mapSquare
	}{
		{
			name:      "steps",
			inputfile: "test_files/input.steps",
			n:         1,
			elfs:      map[location]mapSquare{{3, 4}: {elf, 200, 3}},
			goblins: map[location]mapSquare{
				{1, 2}: {goblin, 200, 3},
				{1, 6}: {goblin, 200, 3},
				{2, 4}: {goblin, 197, 3},
				{3, 7}: {goblin, 200, 3},
				{4, 2}: {goblin, 200, 3},
				{6, 1}: {goblin, 200, 3},
				{6, 4}: {goblin, 200, 3},
				{6, 7}: {goblin, 200, 3},
			},
		},
		{
			name:      "attack1",
			inputfile: "test_files/input.attack",
			n:         1,
			elfs: map[location]mapSquare{
				{2, 4}: {elf, 197, 3},
				{4, 5}: {elf, 197, 3},
			},
			goblins: map[location]mapSquare{
				{1, 3}: {goblin, 200, 3},
				{2, 5}: {goblin, 197, 3},
				{3, 3}: {goblin, 200, 3},
				{3, 5}: {goblin, 197, 3},
			},
		},
		{
			name:      "attack2",
			inputfile: "test_files/input.attack",
			n:         2,
			elfs: map[location]mapSquare{
				{2, 4}: {elf, 188, 3},
				{4, 5}: {elf, 194, 3},
			},
			goblins: map[location]mapSquare{
				{1, 4}: {goblin, 200, 3},
				{2, 3}: {goblin, 200, 3},
				{2, 5}: {goblin, 194, 3},
				{3, 5}: {goblin, 194, 3},
			},
		},
		{
			name:      "attack23",
			inputfile: "test_files/input.attack",
			n:         23,
			elfs: map[location]mapSquare{
				{4, 5}: {elf, 131, 3},
			},
			goblins: map[location]mapSquare{
				{1, 4}: {goblin, 200, 3},
				{2, 3}: {goblin, 200, 3},
				{2, 5}: {goblin, 131, 3},
				{3, 5}: {goblin, 131, 3},
			},
		},
		{
			name:      "attack24",
			inputfile: "test_files/input.attack",
			n:         24,
			elfs: map[location]mapSquare{
				{4, 5}: {elf, 128, 3},
			},
			goblins: map[location]mapSquare{
				{1, 3}: {goblin, 200, 3},
				{3, 3}: {goblin, 200, 3},
				{2, 4}: {goblin, 131, 3},
				{3, 5}: {goblin, 128, 3},
			},
		},
		{
			name:      "attack25",
			inputfile: "test_files/input.attack",
			n:         25,
			elfs: map[location]mapSquare{
				{4, 5}: {elf, 125, 3},
			},
			goblins: map[location]mapSquare{
				{1, 2}: {goblin, 200, 3},
				{4, 3}: {goblin, 200, 3},
				{2, 3}: {goblin, 131, 3},
				{3, 5}: {goblin, 125, 3},
			},
		},
		{
			name:      "attack47",
			inputfile: "test_files/input.attack",
			n:         47,
			elfs:      map[location]mapSquare{},
			goblins: map[location]mapSquare{
				{1, 1}: {goblin, 200, 3},
				{5, 5}: {goblin, 200, 3},
				{2, 2}: {goblin, 131, 3},
				{3, 5}: {goblin, 59, 3},
			},
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

			for i := 0; i < tt.n; i++ {
				gameStep(gm)
			}

			deep.CompareUnexportedFields = true
			if diff := deep.Equal(tt.elfs, gm.elfs); diff != nil {
				t.Error("elfs ", diff)
			}
			if diff := deep.Equal(tt.goblins, gm.goblins); diff != nil {
				t.Error("goblins ", diff)
			}

		})
	}
}

func Test_games(t *testing.T) {
	tests := []struct {
		name      string
		inputfile string
		rounds    int
		hitpoints int
	}{
		{
			name:      "input1",
			inputfile: "test_files/input1",
			rounds:    37,
			hitpoints: 982,
		},
		{
			name:      "input2",
			inputfile: "test_files/input2",
			rounds:    46,
			hitpoints: 859,
		},
		{
			name:      "input3",
			inputfile: "test_files/input3",
			rounds:    35,
			hitpoints: 793,
		},
		{
			name:      "input4",
			inputfile: "test_files/input4",
			rounds:    54,
			hitpoints: 536,
		},
		{
			name:      "input5",
			inputfile: "test_files/input5",
			rounds:    20,
			hitpoints: 937,
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

			rounds := runGame(gm)
			if rounds != tt.rounds {
				t.Errorf("Number of rounds doesn't match. expected %d, got %d ", tt.rounds, rounds)
			}
			hp := gm.getHealth()
			if hp != tt.hitpoints {
				t.Errorf("hitpoints do not match, expected %d, got %d", tt.hitpoints, hp)
			}
		})
	}

}
