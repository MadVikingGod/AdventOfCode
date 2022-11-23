package main

import "testing"

func TestLookAndSay(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		want string
	}{
		{
			name: "1",
			want: "11",
		},
		{
			name: "11",
			want: "21",
		},
		{
			name: "21",
			want: "1211",
		},
		{
			name: "1211",
			want: "111221",
		},
		{
			name: "111221",
			want: "312211",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LookAndSay(tt.name); got != tt.want {
				t.Errorf("LookAndSay(%s) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
