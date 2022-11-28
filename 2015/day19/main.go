package main

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

var molecule = `CRnSiRnCaPTiMgYCaPTiRnFArSiThFArCaSiThSiThPBCaCaSiRnSiRnTiTiMgArPBCaPMgYPTiRnFArFArCaSiRnBPMgArPRnCaPTiRnFArCaSiThCaCaFArPBCaCaPTiTiRnFArCaSiRnSiAlYSiThRnFArArCaSiRnBFArCaCaSiRnSiThCaCaCaFYCaPTiBCaSiThCaSiThPMgArSiRnCaPBFYCaCaFArCaCaCaCaSiThCaSiRnPRnFArPBSiThPRnFArSiRnMgArCaFYFArCaSiRnSiAlArTiTiTiTiTiTiTiRnPMgArPTiTiTiBSiRnSiAlArTiTiRnPMgArCaFYBPBPTiRnSiRnMgArSiThCaFArCaSiThFArPRnFArCaSiRnTiBSiThSiRnSiAlYCaFArPRnFArSiThCaFArCaCaSiThCaCaCaSiRnPRnCaFArFYPMgArCaPBCaPBSiRnFYPBCaFArCaSiAl`

func main() {
	replaces := parseReplaces(input)
	molecules := map[string]struct{}{}
	for _, r := range replaces {
		for _, m := range r.Exchange(molecule) {
			molecules[m] = struct{}{}
		}
	}
	println(len(molecules))
	molecules = map[string]struct{}{
		molecule: {},
	}
	count := 0
	for {
		if _, ok := molecules["e"]; ok {
			break
		}
		count++
		newMolecules := map[string]struct{}{}
		for m := range molecules {
			for _, r := range replaces {
				for _, newM := range r.ReverseExchange(m) {
					newMolecules[newM] = struct{}{}
				}
			}
		}
		molecules = newMolecules
		println(count, len(molecules))
	}
	println(count)
}

type replace struct {
	from string
	to   string
}

func (r replace) Exchange(s string) []string {
	var out []string
	for i := 0; i < len(s)-(len(r.from)-1); i++ {
		if s[i:i+len(r.from)] == r.from {
			out = append(out, s[:i]+r.to+s[i+len(r.from):])
		}
	}
	return out
}

func (r replace) ReverseExchange(s string) []string {
	var out []string
	for i := 0; i < len(s)-(len(r.to)-1); i++ {
		if s[i:i+len(r.to)] == r.to {
			out = append(out, s[:i]+r.from+s[i+len(r.to):])
		}
	}
	return out
}

func parseReplaces(s string) []replace {
	var out []replace
	for _, line := range strings.Split(s, "\n") {
		parts := strings.Split(line, " => ")
		out = append(out, replace{parts[0], parts[1]})
	}
	return out
}
