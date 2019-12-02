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
			ic.add()
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
			ic.mul()
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ic := &Intcode{
				Memory:         tt.fields.Memory,
				ProgramCounter: tt.fields.ProgramCounter,
			}
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
