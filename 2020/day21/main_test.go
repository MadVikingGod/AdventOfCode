package main

import (
	"reflect"
	"testing"
)

func Test_parse(t *testing.T) {

	tests := []struct {
		s     string
		want  stringSet
		want1 stringSet
	}{
		{
			s:     "mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
			want:  New("mxmxvkd", "kfcds", "sqjhc", "nhms"),
			want1: New("dairy", "fish"),
		},
		{
			s:     "trh fvjkl sbzzf mxmxvkd (contains dairy)",
			want:  New("trh", "fvjkl", "sbzzf", "mxmxvkd"),
			want1: New("dairy"),
		},
		{
			s:     "sqjhc fvjkl (contains soy)",
			want:  New("sqjhc", "fvjkl"),
			want1: New("soy"),
		},
		{
			s:     "sqjhc mxmxvkd sbzzf (contains fish)",
			want:  New("sqjhc", "mxmxvkd", "sbzzf"),
			want1: New("fish"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got, got1 := parse(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
