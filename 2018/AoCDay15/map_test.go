package main

import "testing"

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
