package intcode

import (
	"reflect"
	"testing"
)

func TestIntcode_add(t *testing.T) {
	type fields struct {
		Memory         []int
		ProgramCounter int
	}
	type want struct {
		Memory         []int
		ProgramCounter int
	}
	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{
			name: "add",
			fields: fields{
				Memory:         []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
				ProgramCounter: 0,
			},
			want: want{
				Memory:         []int{1, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
				ProgramCounter: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ic := &Intcode{
				Memory:         tt.fields.Memory,
				ProgramCounter: tt.fields.ProgramCounter,
			}
			ic.binary(add2)()
			if ic.ProgramCounter != tt.want.ProgramCounter {
				t.Errorf("ProgramCounter does not match got %d, want %d", ic.ProgramCounter, tt.want.ProgramCounter)
			}
			if !reflect.DeepEqual(ic.Memory, tt.want.Memory) {
				t.Errorf("Memory doesn't match got %v, want %v", ic.Memory, tt.want.Memory)
			}
		})
	}
}

func TestIntcode_mul(t *testing.T) {
	type fields struct {
		Memory         []int
		ProgramCounter int
	}
	type want struct {
		Memory         []int
		ProgramCounter int
	}
	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{
			name: "mul",
			fields: fields{
				Memory:         []int{1, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
				ProgramCounter: 4,
			},
			want: want{
				Memory:         []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
				ProgramCounter: 8,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ic := &Intcode{
				Memory:         tt.fields.Memory,
				ProgramCounter: tt.fields.ProgramCounter,
			}
			ic.binary(mul2)()
			if ic.ProgramCounter != tt.want.ProgramCounter {
				t.Errorf("ProgramCounter does not match got %d, want %d", ic.ProgramCounter, tt.want.ProgramCounter)
			}
			if !reflect.DeepEqual(ic.Memory, tt.want.Memory) {
				t.Errorf("Memory doesn't match got %v, want %v", ic.Memory, tt.want.Memory)
			}
		})
	}
}

func TestIntcode_decode(t *testing.T) {
	type fields struct {
		Memory         []int
		ProgramCounter int
	}
	type want struct {
		Memory         []int
		ProgramCounter int
		halt           bool
		err            bool
	}
	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{
			name: "add",
			fields: fields{
				Memory:         []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
				ProgramCounter: 0,
			},
			want: want{
				Memory:         []int{1, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
				ProgramCounter: 4,
			},
		},
		{
			name: "mul",
			fields: fields{
				Memory:         []int{1, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
				ProgramCounter: 4,
			},
			want: want{
				Memory:         []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
				ProgramCounter: 8,
			},
		},
		{
			name: "halt",
			fields: fields{
				Memory:         []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
				ProgramCounter: 8,
			},
			want: want{
				Memory:         []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
				ProgramCounter: 8,
				halt:           true,
			},
		},
		{
			name: "oom",
			fields: fields{
				Memory:         []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
				ProgramCounter: 12,
			},
			want: want{
				Memory:         []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
				ProgramCounter: 12,
				halt:           false,
				err:            true,
			},
		},
		{
			name: "not implemented",
			fields: fields{
				Memory:         []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
				ProgramCounter: 0,
			},
			want: want{
				Memory:         []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
				ProgramCounter: 0,
				halt:           false,
				err:            true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ic := &Intcode{
				Memory:         tt.fields.Memory,
				ProgramCounter: tt.fields.ProgramCounter,
			}
			ic.register()
			got, err := ic.decode()
			if (err != nil) != tt.want.err {
				t.Errorf("Intcode.decode() error = %v, wantErr %v", err, tt.want.err)
				return
			}
			if got != tt.want.halt {
				t.Errorf("Intcode.decode() = %v, want %v", got, tt.want.halt)
			}
			if ic.ProgramCounter != tt.want.ProgramCounter {
				t.Errorf("ProgramCounter does not match got %d, want %d", ic.ProgramCounter, tt.want.ProgramCounter)
			}
			if !reflect.DeepEqual(ic.Memory, tt.want.Memory) {
				t.Errorf("Memory doesn't match got %v, want %v", ic.Memory, tt.want.Memory)
			}
		})
	}
}

func TestIntcode_Run(t *testing.T) {
	type fields struct {
		Memory         []int
		ProgramCounter int
	}
	tests := []struct {
		name    string
		Memory  []int
		want    int
		wantErr bool
	}{
		{
			name:    "add and mul",
			Memory:  []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			want:    3500,
			wantErr: false,
		},
		{
			name:    "not implmented",
			Memory:  []int{100, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ic := New(tt.Memory)
			got, err := ic.Run()
			if (err != nil) != tt.wantErr {
				t.Errorf("Intcode.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Intcode.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDay2_Run(b *testing.B) {
	input := []int{
		1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 10, 1, 19, 2, 9, 19, 23, 2, 13, 23, 27, 1, 6, 27, 31, 2, 6, 31, 35, 2, 13, 35, 39, 1, 39, 10, 43, 2, 43, 13, 47, 1, 9, 47, 51, 1, 51, 13, 55, 1, 55, 13, 59, 2, 59, 13, 63, 1, 63, 6, 67, 2, 6, 67, 71, 1, 5, 71, 75, 2, 6, 75, 79, 1, 5, 79, 83, 2, 83, 6, 87, 1, 5, 87, 91, 1, 6, 91, 95, 2, 95, 6, 99, 1, 5, 99, 103, 1, 6, 103, 107, 1, 107, 2, 111, 1, 111, 5, 0, 99, 2, 14, 0, 0,
	}
	for n := 0; n < b.N; n++ {
		memory := make([]int, len(input))
		copy(memory, input)
		ic := New(memory)
		ic.Run()
	}
}
