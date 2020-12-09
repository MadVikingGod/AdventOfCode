package main

import (
	"reflect"
	"testing"
)

func TestLimitedSet_isValid1(t *testing.T) {
	set := new(25)
	set.add(20)
	for i := 1; i < 20; i++ {
		set.add(i)
	}
	for i := 21; i < 26; i++ {
		set.add(i)
	}
	if !set.isValid(45) {
		t.Errorf("45 should be valid")
	}
	set.add(45)
	if set.isValid(65) {
		t.Errorf("65 should not be valid")
	}

}
func TestLimitedSet_isValid2(t *testing.T) {
	set := new(5)
	items := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}
	for _, i := range items {
		if !set.isValid(i) && i != 127 {
			t.Error("only 127 should not be valid")
		}
		set.add(i)
	}
}

func Test_find(t *testing.T) {
	type args struct {
		target int
		l      []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "127",
			args: args{
				target: 127,
				l: []int{
					35,
					20,
					15,
					25,
					47,
					40,
					62,
					55,
					65,
					95,
					102,
					117,
					150,
					182,
					127,
					219,
					299,
					277,
					309,
					576,
				},
			},
			want: []int{15, 25, 47, 40},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find(tt.args.target, tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("find() = %v, want %v", got, tt.want)
			}
		})
	}
}
