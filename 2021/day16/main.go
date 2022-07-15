package main

import (
	"fmt"
	"math"
)

func main() {
	bits := parse(input)
	pkt, _ := parsePacket(bits)
	fmt.Println(pkt.Versions())
	fmt.Println(pkt.Value())
}

type bits []int

func parse(input string) bits {
	output := []int{}
	lookup := map[rune][]int{
		'0': []int{0, 0, 0, 0},
		'1': []int{0, 0, 0, 1},
		'2': []int{0, 0, 1, 0},
		'3': []int{0, 0, 1, 1},
		'4': []int{0, 1, 0, 0},
		'5': []int{0, 1, 0, 1},
		'6': []int{0, 1, 1, 0},
		'7': []int{0, 1, 1, 1},
		'8': []int{1, 0, 0, 0},
		'9': []int{1, 0, 0, 1},
		'A': []int{1, 0, 1, 0},
		'B': []int{1, 0, 1, 1},
		'C': []int{1, 1, 0, 0},
		'D': []int{1, 1, 0, 1},
		'E': []int{1, 1, 1, 0},
		'F': []int{1, 1, 1, 1},
	}
	for _, r := range input {
		output = append(output, lookup[r]...)
	}
	return bits(output)
}

func (b bits) parseHeader() (int, int, bits) {
	version := combine(b, 3)
	b = b[3:]
	id := combine(b, 3)
	b = b[3:]
	return version, id, b
}

func combine(bits []int, len int) int {
	output := 0
	for i := 0; i < len; i++ {
		output = output<<1 + bits[i]
	}
	return output
}

func parseVarInt(bits []int) (int, []int) {
	cont := true
	value := 0
	for cont {
		cont = bits[0] == 1
		bits = bits[1:]
		value = value<<4 + combine(bits, 4)
		bits = bits[4:]
	}
	return value, bits
}

type packet struct {
	version int
	id      int

	value      int
	subpackets []packet
}

func parsePacket(bits bits) (packet, bits) {
	pkt := packet{}

	pkt.version, pkt.id, bits = bits.parseHeader()
	if pkt.id == 4 {
		pkt.value, bits = parseVarInt(bits)
		return pkt, bits
	}
	if bits[0] == 0 {
		bits = bits[1:]
		length := combine(bits, 15)
		bits = bits[15:]
		target := len(bits) - length
		var subPkt packet
		for len(bits) > target {
			subPkt, bits = parsePacket(bits)
			pkt.subpackets = append(pkt.subpackets, subPkt)
		}
		return pkt, bits
	}
	bits = bits[1:]
	count := combine(bits, 11)
	bits = bits[11:]
	var subPkt packet
	for i := 0; i < count; i++ {
		subPkt, bits = parsePacket(bits)
		pkt.subpackets = append(pkt.subpackets, subPkt)
	}
	return pkt, bits
}

func (p packet) Versions() int {
	ver := p.version
	for _, pkt := range p.subpackets {
		ver += pkt.Versions()
	}
	return ver
}

func (p packet) Value() int {
	switch p.id {
	case 0:
		sum := 0
		for _, pkt := range p.subpackets {
			sum += pkt.Value()
		}
		return sum
	case 1:
		prod := 1
		for _, pkt := range p.subpackets {
			prod *= pkt.Value()
		}
		return prod
	case 2:
		min := math.MaxInt
		for _, pkt := range p.subpackets {
			if pkt.Value() < min {
				min = pkt.Value()
			}
		}
		return min
	case 3:
		max := math.MinInt
		for _, pkt := range p.subpackets {
			if pkt.Value() > max {
				max = pkt.Value()
			}
		}
		return max
	case 4:
		return p.value
	case 5:
		if p.subpackets[0].Value() > p.subpackets[1].Value() {
			return 1
		} else {
			return 0
		}
	case 6:
		if p.subpackets[0].Value() < p.subpackets[1].Value() {
			return 1
		} else {
			return 0
		}
	case 7:
		if p.subpackets[0].Value() == p.subpackets[1].Value() {
			return 1
		} else {
			return 0
		}
	}

	return -1
}
