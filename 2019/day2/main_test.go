package main

import (
	"reflect"
	"testing"
)

func Test_add(t *testing.T) {
	type args struct {
		memory   []int
		position int
	}
	type want struct {
		memory []int
		position int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "add",
			args: args{
				memory: []int{1,9,10,3,2,3,11,0,99,30,40,50},
				position: 0,
			},
			want: want{
				memory: []int{1,9,10,70,2,3,11,0,99,30,40,50},
				position: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.memory, tt.args.position); got != tt.want.position {
				t.Errorf("add() = %v, want %v", got, tt.want.position)
			}
			t.Log(tt.args.memory)
			if !reflect.DeepEqual(tt.args.memory, tt.want.memory) {
				t.Errorf("memory doesn't match got %v, want %v", tt.args.memory, tt.want.memory)
			}
		})
	}
}

func Test_mul(t *testing.T) {
	type args struct {
		memory   []int
		position int
	}
	type want struct {
		memory []int
		position int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mul(tt.args.memory, tt.args.position); got != tt.want.position {
				t.Errorf("mul() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.args.memory, tt.want.memory) {
				t.Errorf("memory doesn't match got %v, want %v", tt.args.memory, tt.want.memory)
			}
		})
	}
}

func Test_readOpcode(t *testing.T) {
	type args struct {
		memory   []int
		position int
	}
	type want struct {
		memory []int
		position int
		stop bool
	}
	tests := []struct {
		name  string
		args  args
		want  want
	}{
		{
			name: "add",
			args: args{
				memory: []int{1,9,10,3,2,3,11,0,99,30,40,50},
				position: 0,
			},
			want: want{
				memory: []int{1,9,10,70,2,3,11,0,99,30,40,50},
				position: 4,
				stop: false,
			},
		},
		{
			name: "mul",
			args: args{
				memory: []int{1,9,10,70,2,3,11,0,99,30,40,50},
				position: 4,
			},
			want: want{
				memory: []int{3500,9,10,70,2,3,11,0,99,30,40,50},
				position: 8,
				stop: false,
			},
		},
		{
			name: "stop",
			args: args{
				memory: []int{3500,9,10,70,2,3,11,0,99,30,40,50},
				position: 8,
			},
			want: want{
				memory: []int{3500,9,10,70,2,3,11,0,99,30,40,50},
				position: 0,
				stop: true,
			},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := readOpcode(tt.args.memory, tt.args.position)
			if got != tt.want.position {
				t.Errorf("readOpcode() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want.stop {
				t.Errorf("readOpcode() got1 = %v, want %v", got1, tt.want.stop)
			}
			if !reflect.DeepEqual(tt.args.memory, tt.want.memory) {
				t.Errorf("memory doesn't match got %v, want %v", tt.args.memory, tt.want.memory)
			}
		})
	}
}
