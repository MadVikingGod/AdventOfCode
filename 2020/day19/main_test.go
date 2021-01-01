package main

import (
	"reflect"
	"testing"
)

func Test_parseRule(t *testing.T) {

	tests := []struct {
		i    int
		want []string
	}{
		{
			i:    54,
			want: []string{"a"},
		},
		{
			i:    117,
			want: []string{"b"},
		},
		{
			i:    56,
			want: []string{"a", "b"},
		},
		{
			i:    104,
			want: []string{"ab", "bb", "ba"},
		},
	}
	for _, tt := range tests {
		t.Run(rules[tt.i], func(t *testing.T) {
			if got := parseRule(tt.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseRegex(t *testing.T) {

	tests := []struct {
		i    int
		want string
	}{
		{
			i:    54,
			want: "a",
		},
		{
			i:    117,
			want: "b",
		},
		{
			i:    56,
			want: "(?:a|b)",
		},
		{
			i:    104,
			want: "(?:(?:a|b)b|ba)",
		},
	}
	for _, tt := range tests {
		t.Run(rules[tt.i], func(t *testing.T) {
			if got := parseRegex(tt.i); got != tt.want {
				t.Errorf("parseRegex() = %v, want %v", got, tt.want)
			}
		})
	}
}