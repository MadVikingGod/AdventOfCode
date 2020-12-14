package main

import (
	"fmt"
	"math"
	"math/bits"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)

func main() {
	mask := mask{}
	mem := map[uint64]uint64{}

	for _, line := range input {
		if line[0:4] == "mask" {
			mask = newMask(line[7:])
		} else {
			matches := re.FindStringSubmatch(line)
			loc, _ := strconv.ParseUint(matches[1], 10, 64)
			val, _ := strconv.ParseUint(matches[2], 10, 64)
			mem[loc] = mask.Mask(val)
		}
	}
	fmt.Println(mask)
	sum := uint64(0)
	for _, val := range mem {
		sum += val
	}
	fmt.Println(sum)

	// Part2
	maskv2 := maskV2{}
	mem = map[uint64]uint64{}
	for _, line := range input {
		if line[0:4] == "mask" {
			maskv2 = newMaskV2(line[7:])
		} else {
			matches := re.FindStringSubmatch(line)
			loc, _ := strconv.ParseUint(matches[1], 10, 64)
			val, _ := strconv.ParseUint(matches[2], 10, 64)

			for _, address := range maskv2.Mask(loc) {
				mem[address] = val
			}
		}
	}
	sum = uint64(0)
	for _, val := range mem {
		sum += val
	}
	fmt.Println(sum)
}

type mask struct {
	clear uint64
	val   uint64
}

func newMask(s string) mask {
	s = strings.ToLower(s)
	m, _ := strconv.ParseUint(strings.ReplaceAll(strings.ReplaceAll(s, "1", "0"), "x", "1"), 2, 64)
	v, _ := strconv.ParseUint(strings.ReplaceAll(s, "x", "0"), 2, 64)
	return mask{
		clear: m,
		val:   v,
	}
}
func (m mask) Mask(x uint64) uint64 {
	return x&m.clear + m.val
}

type maskV2 struct {
	or    uint64
	float uint64
}

func newMaskV2(s string) maskV2 {
	s = strings.ToLower(s)
	or, _ := strconv.ParseUint(strings.ReplaceAll(s, "x", "0"), 2, 64)
	float, _ := strconv.ParseUint(strings.ReplaceAll(strings.ReplaceAll(s, "1", "0"), "x", "1"), 2, 64)
	return maskV2{
		or:    or,
		float: float,
	}
}

func (m maskV2) Mask(address uint64) []uint64 {
	out := make([]uint64, 1<<bits.OnesCount64(m.float))
	for i := range out {
		out[i] = address | m.or
	}
	k := 0
	for j := 0; j < 36; j++ {
		if (m.float>>j)&1 == 1 {
			//mask
			x := uint64(1) << j
			for i := range out {
				if i>>k%2 == 0 {
					//clear
					out[i] &= math.MaxUint64 ^ x // There isn't a complement, so use Xor
				} else {
					//set
					out[i] |= x
				}
			}
			k++
		}
	}
	return out
}
