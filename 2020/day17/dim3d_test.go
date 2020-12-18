package main

import (
	"reflect"
	"testing"
)

func Test_point3d_Range(t *testing.T) {
	type fields struct {
		x int
		y int
		z int
	}
	type args struct {
		p point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []point
	}{
		{
			name: "",
			fields: fields{
				x: 0,
				y: 0,
				z: 0,
			},
			args: args{
				p: point3d{1, 1, 1},
			},
			want: []point{
				point3d{0, 0, 0},
				point3d{0, 0, 1},
				point3d{0, 1, 0},
				point3d{0, 1, 1},
				point3d{1, 0, 0},
				point3d{1, 0, 1},
				point3d{1, 1, 0},
				point3d{1, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := point3d{
				x: tt.fields.x,
				y: tt.fields.y,
				z: tt.fields.z,
			}
			if got := a.Range(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("point3d.Range() = %v, want %v", got, tt.want)
			}
		})
	}
}
