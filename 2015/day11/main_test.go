package main

import (
	"testing"
)

func Test_hasStraight(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "hijklmmn",
			want: true,
		},
		{
			name: "abcdffaa",
			want: true,
		},
		{
			name: "ghjaabcc",
			want: true,
		},
		{
			name: "abbceffg",
			want: false,
		},
		{
			name: "abbcegjk",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasStraight(tt.name); got != tt.want {
				t.Errorf("hasStraight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_noConfusing(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "hijklmmn",
			want: false,
		},
		{
			name: "abcdffaa",
			want: true,
		},
		{
			name: "ghjaabcc",
			want: true,
		},
		{
			name: "abbceffg",
			want: true,
		},
		{
			name: "abbcegjk",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := noConfusing(tt.name); got != tt.want {
				t.Errorf("noConfusing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoDoubles(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "hijklmmn",
			want: false,
		},
		{
			name: "abcdffaa",
			want: true,
		},
		{
			name: "ghjaabcc",
			want: true,
		},
		{
			name: "abbceffg",
			want: true,
		},
		{
			name: "abbcegjk",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoDoubles(tt.name); got != tt.want {
				t.Errorf("twoDoubles(%s) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func Test_next(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "abcdefgh",
			want: "abcdffaa",
		},
		{
			name: "ghijklmn",
			want: "ghjaabcc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := next(tt.name); got != tt.want {
				t.Errorf("next(%s) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func Test_increment(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "abcdefgh",
			want: "abcdefgi",
		},
		{
			name: "abcdexyz",
			want: "abcdexza",
		},
		{
			name: "abcdexzz",
			want: "abcdeyaa",
		},
		{
			name: "abcdezaz",
			want: "abcdezba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increment(tt.name); got != tt.want {
				t.Errorf("increment() = %v, want %v", got, tt.want)
			}
		})
	}
}
