package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPraseVarInt(t *testing.T) {

	input := "D2FE28"
	bits := parse(input)
	ver, id, bits := bits.parseHeader()

	assert.Equal(t, ver, 6)
	assert.Equal(t, id, 4)

	val, bits := parseVarInt(bits)

	assert.Equal(t, val, 2021)

}

// func TestPrasePacket(t *testing.T) {
// 	tests := []struct {
// 		input string
// 		want  packet
// 	}{
// 		{
// 			input: "38006F45291200",
// 			want: ,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.input, func(t *testing.T) {
// 			bits := parse(tt.input)
// 			pkt, _ := parsePacket(bits)

// 			assert.Equal(t, tt.want, pkt)
// 		})
// 	}
// }

func TestVersion(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "8A004A801A8002F478",
			want:  16,
		},
		{
			input: "620080001611562C8802118E34",
			want:  12,
		},
		{
			input: "C0015000016115A2E0802F182340",
			want:  23,
		},
		{
			input: "A0016C880162017C3686B18A3D4780",
			want:  31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			bits := parse(tt.input)
			pkt, _ := parsePacket(bits)
			assert.Equal(t, tt.want, pkt.Versions())
		})
	}
}

func TestValue(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"C200B40A82", 3},
		{"04005AC33890", 54},
		{"880086C3E88112", 7},
		{"CE00C43D881120", 9},
		{"D8005AC2A8F0", 1},
		{"F600BC2D8F", 0},
		{"9C005AC2F8F0", 0},
		{"9C0141080250320F1802104A08", 1},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			bits := parse(tt.input)
			pkt, _ := parsePacket(bits)
			assert.Equal(t, tt.want, pkt.Value())
		})
	}
}
