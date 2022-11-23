package main

import (
	"crypto/md5"
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(getHash(input))
	fmt.Println(getHash2(input))
}

func getHash(input string) int {
	h := md5.New()
	for i := 0; i < 1000000000; i++ {
		h.Write([]byte(input + strconv.Itoa(i)))
		if fmt.Sprintf("%x", h.Sum(nil))[:5] == "00000" {
			return i
		}
		h.Reset()
	}
	return -1
}

func getHash2(input string) int {
	h := md5.New()
	for i := 0; i < 1000000000; i++ {
		h.Write([]byte(input + strconv.Itoa(i)))
		if fmt.Sprintf("%x", h.Sum(nil))[:6] == "000000" {
			return i
		}
		h.Reset()
	}
	return -1
}
