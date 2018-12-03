package helpers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

func GetInput(day int) ([]string, error) {
	reader, err := getHttpOrLocal(day)
	if err != nil {
		return nil, err
	}

	body := []byte{}
	_, err = reader.Read(body)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	inputs := strings.Split(string(body), "\n")
	return inputs, nil
}

func getHttpOrLocal(day int) (io.ReadCloser, error) {
	if os.Getenv("SESSION") == "" {
		return getLocal(day)
	}
	return getHttp(day)
}

func getLocal(day int) (*os.File, error) {
	filePath := path.Join(os.Getenv("HOME"), "go", "src", "github.com", "madvikinggod", "AdventOfCode", "2018", fmt.Sprintf("AoCDay%d", day), "input")
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func getHttp(day int) (io.ReadCloser, error) {
	url := "https://adventofcode.com/2018/day/%d/input"
	resp, err := http.Get(fmt.Sprintf(url, day))
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func GetInts(in []string) ([]int, error) {
	out := []int{}
	for _, s := range in {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		out = append(out, i)
	}
	return out, nil
}
