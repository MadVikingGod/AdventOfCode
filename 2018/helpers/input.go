package helpers

import (
	"fmt"
	"io"
	"io/ioutil"
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

	body, err := ioutil.ReadAll(reader)
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
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	filePath := correctInputFile(wd, day)
	return os.Open(filePath)

}

func correctInputFile(cwd string, Day int) string {
	day := strconv.Itoa(Day)
	if strings.HasSuffix(cwd, day) {
		return path.Join(cwd, "input")
	}
	return path.Join(cwd, "AoCDay"+day, "input")
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
