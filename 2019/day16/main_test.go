package main

import (
	"gonum.org/v1/gonum/mat"
	"reflect"
	"testing"
)

func Test_createMat(t *testing.T) {
	type args struct {
		len int
	}
	tests := []struct {
		name string
		args args
		want *mat.Dense
	}{
		{
			name: "mat 4",
			args: args{
				len: 4,
			},
			want: mat.NewDense(4,4, []float64{
					1,0,-1,0,
					0,1, 1,0,
					0,0,1,1,
					0,0,0,1,
			}),
		},
		{
			name: "mat 8",
			args: args{
				len: 8,
			},
			want: mat.NewDense(8,8, []float64{
					1, 0,-1, 0, 1, 0,-1, 0,
					0, 1, 1, 0, 0,-1,-1, 0,
					0, 0, 1, 1, 1, 0, 0, 0,
					0, 0, 0, 1, 1, 1, 1, 0,
					0, 0, 0, 0, 1, 1, 1, 1,
					0, 0, 0, 0, 0, 1, 1, 1,
					0, 0, 0, 0, 0, 0, 1, 1,
					0, 0, 0, 0, 0, 0, 0, 1,
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createMat(tt.args.len); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createMat() = %v, want %v", got, tt.want)
			}
		})
	}
}