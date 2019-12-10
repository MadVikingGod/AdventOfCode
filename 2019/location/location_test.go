package location

import (
	"reflect"
	"testing"
)

func Test_gcd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "no common",
			args: args{
				a: 2,
				b: 3,
			},
			want: 1,
		},
		{
			name: "common",
			args: args{
				a: 2,
				b: 4,
			},
			want: 2,
		},
		{
			name: "reverse",
			args: args{
				a: 4,
				b: 2,
			},
			want: 2,
		},
		{
			name: "Negative",
			args: args{
				a: -8,
				b: 4,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocation_Direction(t *testing.T) {
	type fields struct {
		x int
		y int
	}
	type args struct {
		l2 Location
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Location
	}{
		{
			name: "neg",
			fields: fields{
				x: 5,
				y: 5,
			},
			args: args{
				l2: Location{
					X: 3,
					Y: 3,
				},
			},
			want: Location{
				X: -1,
				Y: -1,
			},
		},
		{
			name: "angle",
			fields: fields{
				x: 5,
				y: 5,
			},
			args: args{
				l2: Location{
					X: 3,
					Y: 4,
				},
			},
			want: Location{
				X: -2,
				Y: -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Location{
				X: tt.fields.x,
				Y: tt.fields.y,
			}
			if got := l.Direction(tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Direction() = %v, want %v", got, tt.want)
			}
		})
	}
}
