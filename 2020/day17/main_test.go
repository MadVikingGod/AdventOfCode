package main

import (
	"fmt"
	"testing"
)

func Test_dim_Step(t *testing.T) {
	dimension := New3d([]string{
		".#.",
		"..#",
		"###",
	})

	out := dimension.Step()
	want := 11
	count := 0
	count = out.Actives()
	if count != want {
		fmt.Println(out)
		t.Errorf("dim.Step() count = %d, want 11", count)
	}

	out = out.Step()
	want = 21
	count = out.Actives()
	if count != want {
		fmt.Println(out)
		t.Errorf("dim.Step()*2 count = %d, want 21", count)
	}

	out = out.Step() //step3
	out = out.Step() //step4
	out = out.Step() //step5
	out = out.Step() //step6
	want = 112
	count = out.Actives()
	if count != want {
		fmt.Println(out)
		t.Errorf("dim.Step()*2 count = %d, want 112", count)
	}

}
func Test_dim_Step4d(t *testing.T) {
	dimension := New4d([]string{
		".#.",
		"..#",
		"###",
	})
	out := dimension.Step()
	count := out.Actives()
	if count != 29 {
		fmt.Println(out)
		t.Errorf("dim.Step()*2 count = %d, want 29", count)
	}
	out = out.Step() //step2
	count = out.Actives()
	if count != 60 {
		fmt.Println(out)
		t.Errorf("dim.Step()*2 count = %d, want 60", count)
	}
	out = out.Step() //step3
	out = out.Step() //step4
	out = out.Step() //step5
	out = out.Step() //step6

	count = out.Actives()
	if count != 848 {
		fmt.Println(out)
		t.Errorf("dim.Step()*2 count = %d, want 848", count)
	}
}
