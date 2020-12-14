package main

import (
	"reflect"
	"testing"
)

func Test_mask_Mask(t *testing.T) {

	type args struct {
		x uint64
	}
	tests := []struct {
		name string
		mask string
		args args
		want uint64
	}{
		{
			name: "11",
			mask: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			args: args{
				x: 11,
			},
			want: 73,
		},
		{
			name: "101",
			mask: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			args: args{
				x: 101,
			},
			want: 101,
		},
		{
			name: "0",
			mask: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			args: args{
				x: 0,
			},
			want: 64,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := newMask(tt.mask)
			if got := m.Mask(tt.args.x); got != tt.want {
				t.Errorf("mask.Mask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maskV2_Mask(t *testing.T) {

	type args struct {
		address uint64
	}
	tests := []struct {
		name string
		mask string
		args args
		want []uint64
	}{
		{
			name: "42",
			mask: "000000000000000000000000000000X1001X",
			args: args{
				address: 42,
			},
			want: []uint64{
				26, 27, 58, 59,
			},
		},
		{
			name: "26",
			mask: "00000000000000000000000000000000X0XX",
			args: args{
				address: 26,
			},
			want: []uint64{
				16, 17, 18, 19, 24, 25, 26, 27,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := newMaskV2(tt.mask)

			if got := m.Mask(tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maskV2.Mask() = %v, want %v", got, tt.want)
			}
		})
	}
}
