package intcode

import (
	"io"
	"reflect"
	"strings"
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
		{
			name: "add immediate",
			fields: fields{
				Memory:         []int{1101, 100, -1, 4, 0},
				ProgramCounter: 0,
			},
			want: want{
				Memory:         []int{1101, 100, -1, 4, 99},
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
			ic.Register()
			halt, err := ic.decode()
			if halt {
				t.Errorf("add should not halt execution")
			}
			if err != nil {
				t.Errorf("add should not cause error, got %v", err)
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
		{
			name: "mul immediate",
			fields: fields{
				Memory:         []int{1002, 4, 3, 4, 33},
				ProgramCounter: 0,
			},
			want: want{
				Memory:         []int{1002, 4, 3, 4, 99},
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
			ic.Register()
			halt, err := ic.decode()
			if halt {
				t.Errorf("mul should not halt execution")
			}
			if err != nil {
				t.Errorf("mul should not cause error, got %v", err)
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
			ic.Register()
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
		Memory []int
		in     io.Reader
		out    io.Writer
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name: "add and mul",
			fields: fields{
				Memory: []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			},
			want:    3500,
			wantErr: false,
		},
		{
			name: "not implmented",
			fields: fields{
				Memory: []int{999999, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ic := New(tt.fields.Memory)
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

func TestIntcode_Run_input_output(t *testing.T) {
	type fields struct {
		Memory []int
		in     io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		wantOut string
	}{
		{
			name: "add and mul",
			fields: fields{
				Memory: []int{3, 0, 4, 0, 99},
				in:     strings.NewReader("5\n"),
			},
			wantOut: "5\n",
		},
		{
			name: "position mode != 8",
			fields: fields{
				Memory: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
				in:     strings.NewReader("5\n"),
			},
			wantOut: "0\n",
		},
		{
			name: "position mode == 8",
			fields: fields{
				Memory: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
				in:     strings.NewReader("8\n"),
			},
			wantOut: "1\n",
		},
		{
			name: "position mode < 8",
			fields: fields{
				Memory: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
				in:     strings.NewReader("5\n"),
			},
			wantOut: "1\n",
		},
		{
			name: "position mode !< 8",
			fields: fields{
				Memory: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
				in:     strings.NewReader("129\n"),
			},
			wantOut: "0\n",
		},
		{
			name: "immediate mode != 8",
			fields: fields{
				Memory: []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
				in:     strings.NewReader("5\n"),
			},
			wantOut: "0\n",
		},
		{
			name: "immediate mode == 8",
			fields: fields{
				Memory: []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
				in:     strings.NewReader("8\n"),
			},
			wantOut: "1\n",
		},
		{
			name: "immediate mode < 8",
			fields: fields{
				Memory: []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
				in:     strings.NewReader("5\n"),
			},
			wantOut: "1\n",
		},
		{
			name: "immediate mode !< 8",
			fields: fields{
				Memory: []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
				in:     strings.NewReader("129\n"),
			},
			wantOut: "0\n",
		},

		{
			name: "position mode jmp 0",
			fields: fields{
				Memory: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
				in:     strings.NewReader("0\n"),
			},
			wantOut: "0\n",
		},
		{
			name: "position mode jmp 1",
			fields: fields{
				Memory: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
				in:     strings.NewReader("129\n"),
			},
			wantOut: "1\n",
		},
		{
			name: "immediate mode jmp 0",
			fields: fields{
				Memory: []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
				in:     strings.NewReader("0\n"),
			},
			wantOut: "0\n",
		},
		{
			name: "immediate mode jmp 1",
			fields: fields{
				Memory: []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
				in:     strings.NewReader("129\n"),
			},
			wantOut: "1\n",
		},

		{
			name: "day5 < 8",
			fields: fields{
				Memory: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
					1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
					999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
				in: strings.NewReader("4\n"),
			},
			wantOut: "999\n",
		},
		{
			name: "day5 == 8",
			fields: fields{
				Memory: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
					1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
					999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
				in: strings.NewReader("8\n"),
			},
			wantOut: "1000\n",
		},
		{
			name: "day5 > 8",
			fields: fields{
				Memory: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
					1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
					999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
				in: strings.NewReader("129\n"),
			},
			wantOut: "1001\n",
		},
		{
			name: "day9 base mode quine",
			fields: fields{
				Memory: append([]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}, make([]int, 100)...),
				in:     strings.NewReader(""),
			},
			wantOut: "109\n1\n204\n-1\n1001\n100\n1\n100\n1008\n100\n16\n101\n1006\n101\n0\n99\n",
		},
		{
			name: "day9 base mode large number",
			fields: fields{
				Memory: []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
				in:     strings.NewReader(""),
			},
			wantOut: "1219070632396864\n",
		},
		{
			name: "day9 base mode large number",
			fields: fields{
				Memory: []int{104, 1125899906842624, 99},
				in:     strings.NewReader(""),
			},
			wantOut: "1125899906842624\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := &strings.Builder{}
			ic := Intcode{
				Memory: tt.fields.Memory,
				In:     tt.fields.in,
				Out:    output,
			}
			ic.Register()
			_, err := ic.Run()
			if err != nil {
				t.Errorf("Intcode.Run() error = %v, wantErr nil", err)
				return
			}
			if output.String() != tt.wantOut {
				t.Errorf("output didn't match got %+q, wanted %+q", output.String(), tt.wantOut)
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
		_, _ = ic.Run()
	}
}
