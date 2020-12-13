package main

import "testing"

func Test_pair_isValid(t *testing.T) {
	type fields struct {
		val int
		pos int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "",
			fields: fields{
				val: 7,
				pos: 0,
			},
			args: args{
				i: 1068788,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				val: 7,
				pos: 0,
			},
			args: args{
				i: 1068789,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				val: 13,
				pos: 1,
			},
			args: args{
				i: 1068788,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				val: 13,
				pos: 1,
			},
			args: args{
				i: 1068789,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				val: 59,
				pos: 4,
			},
			args: args{
				i: 1068788,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				val: 59,
				pos: 4,
			},
			args: args{
				i: 1068789,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pair{
				val: tt.fields.val,
				pos: tt.fields.pos,
			}
			p.isValid(tt.args.i)
		})
	}
}
