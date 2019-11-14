package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

var test_map *gameMap

func getTestMap() (*gameMap, error) {
	if test_map != nil {
		return test_map, nil
	}
	buf, err := ioutil.ReadFile("test_input")
	if err != nil {
		return nil, err
	}
	inputs := strings.Split(string(buf), "\n")

	test_map = NewGameMap(inputs)
	return test_map, nil
}

func TestNewGameMap(t *testing.T) {

	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "test_file",
			input: "test_input",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getTestMap()
			if err != nil {
				t.Error(err)
			}

			fmt.Print(got)
			fmt.Printf("%#v\n", got)
		})
	}
}

func Test_gameMap_FindDown(t *testing.T) {
	tests := []struct {
		name      string
		start     location
		end       location
		offBottom bool
	}{
		{
			name:      "start",
			start:     location{500, 0},
			end:       location{500, 6},
			offBottom: false,
		},
		{
			name:      "inf",
			start:     location{506, 9},
			end:       location{506, 13},
			offBottom: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := getTestMap()
			if err != nil {
				t.Error(err)
			}

			got, got1 := m.FindDown(tt.start)
			if !reflect.DeepEqual(got, tt.end) {
				t.Errorf("gameMap.FindDown() end = %v, want %v", got, tt.end)
			}
			if got1 != tt.offBottom {
				t.Errorf("gameMap.FindDown() offBottom = %v, want %v", got1, tt.offBottom)
			}
		})
	}
}
