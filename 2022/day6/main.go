package main

import _ "embed"

//go:embed input.txt
var input string

func main() {
	// Part 1
	println(findStart("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	println(findStart("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	println(findStart("nppdvjthqldpwncqszvftbrmjlhg"))
	println(findStart("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	println(findStart("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
	println(findStart(input))
	println(findStartOfMessage("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	println(findStartOfMessage("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	println(findStartOfMessage("nppdvjthqldpwncqszvftbrmjlhg"))
	println(findStartOfMessage("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	println(findStartOfMessage("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
	println(findStartOfMessage(input))
}

// findStart returns the position of the first non-repeating 4 character substring
func findStart(s string) int {
	for i := 3; i < len(s); i++ {
		if s[i-3] != s[i-2] && s[i-3] != s[i-1] && s[i-3] != s[i] &&
			s[i-2] != s[i-1] && s[i-2] != s[i] &&
			s[i-1] != s[i] {
			return i + 1
		}
	}
	return -1
}

// findStartofMessage returns the position of the first non-repeating 14 character substring
func findStartOfMessage(s string) int {
	for i := 13; i < len(s); i++ {
		for j := i - 13; j < i; j++ {
			for k := j + 1; k <= i; k++ {
				if s[j] == s[k] {
					goto next
				}
			}
		}
		return i + 1
	next:
	}
	return -1
}
