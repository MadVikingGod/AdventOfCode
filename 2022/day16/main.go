package main

import (
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

//go:embed testInput.txt
var testInput string

func main() {

	tunnel := parseTunnel(testInput)
	puml := PrintActivity(tunnel)
	os.WriteFile("testInput.puml", puml, 0644)

	m := tunnel.toMinTunnel()
	os.WriteFile("testInputMin.puml", minStateGraph(m), 0644)
	fmt.Println(m.totalFlow(30, "AA", "DD", "BB", "JJ", "HH", "EE", "CC"))
	fmt.Println(m.findOptimal([]string{"AA"}))
	fmt.Println(m.findOptimal2(newState(m, 26)))

	tunnel = parseTunnel(input)
	puml = PrintActivity(tunnel)
	os.WriteFile("output.puml", puml, 0644)
	m = tunnel.toMinTunnel()
	os.WriteFile("outputMin.puml", minStateGraph(m), 0644)
	start := time.Now()
	fmt.Println(m.findOptimal([]string{"AA"}), time.Since(start))
	//1606 is too low
	//[AA UK CO IJ NA SE KF CS MN] 1862
	file, _ := os.Create("part2.pprof")
	pprof.StartCPUProfile(file)
	fmt.Println(m.findOptimal2(newState(m, 26)))
	pprof.StopCPUProfile()
	file.Close()
}

type Tunnel map[string]Room

func (t Tunnel) distance(from, to string) int {
	seen := make(map[string]bool)
	type state struct {
		name  string
		depth int
	}
	queue := []state{{name: from, depth: 0}}
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		if seen[s.name] {
			continue
		}
		seen[s.name] = true
		if s.name == to {
			return s.depth
		}
		for _, c := range t[s.name].connections {
			queue = append(queue, state{name: c, depth: s.depth + 1})
		}
	}
	return -1
}

func (t Tunnel) toMinTunnel() minTunnel {
	nodes := []string{}
	for n, r := range t {
		if r.flow > 0 {
			nodes = append(nodes, n)
		}
	}
	m := minTunnel{
		weights:     make(map[string]map[string]int),
		flatWeights: make(map[string]int),
		flows:       make(map[string]int),
	}
	m.weights["AA"] = map[string]int{}
	for _, n := range nodes {
		m.flows[n] = t[n].flow
		m.weights[n] = map[string]int{}
		m.weights["AA"][n] = t.distance("AA", n)
		m.flatWeights["AA"+n] = t.distance("AA", n)
	}
	for i := 0; i < len(nodes)-1; i++ {
		for j := i + 1; j < len(nodes); j++ {
			m.weights[nodes[i]][nodes[j]] = t.distance(nodes[i], nodes[j])
			m.weights[nodes[j]][nodes[i]] = t.distance(nodes[i], nodes[j])

			m.flatWeights[nodes[i]+nodes[j]] = t.distance(nodes[i], nodes[j])
			m.flatWeights[nodes[j]+nodes[i]] = t.distance(nodes[i], nodes[j])
		}
	}
	return m
}

type minTunnel struct {
	weights     map[string]map[string]int
	flatWeights map[string]int
	flows       map[string]int
}

func (t minTunnel) distance(from, to string) int {
	return t.weights[from][to]
}
func (t minTunnel) cumWeight(steps ...string) int {
	total := 0
	for i := 1; i < len(steps); i++ {
		total += t.weights[steps[i-1]][steps[i]]
	}
	return total
}

func (t minTunnel) totalFlow(max int, steps ...string) int {
	total := 0
	cumWeight := 0
	for i := 1; i < len(steps); i++ {
		s := steps[i]
		cumWeight += t.weights[steps[i-1]][s]
		total += (max - cumWeight - i) * t.flows[s]
	}
	return total
}

func (t minTunnel) findOptimal(current []string) int {
	best := 0
	seen := make(map[string]bool)
	for _, c := range current {
		seen[c] = true
	}
	from := current[len(current)-1]
	for to := range t.weights[from] {
		if seen[to] {
			continue
		}
		current = append(current, to)
		if len(current)+t.cumWeight(current...) > 30 {
			current = current[:len(current)-1]
			continue
		}
		score := t.findOptimal(current)
		if score > best {
			best = score
		}

		current = current[:len(current)-1]
	}
	if best == 0 {
		return t.totalFlow(30, current...)
	}
	return best
}

type state struct {
	tunnel          minTunnel
	max             int
	seen            map[string]bool
	self            []string
	selfWeight      int
	elephants       []string
	elephantsWeight int
}

func newState(t minTunnel, max int) *state {
	return &state{
		tunnel:    t,
		max:       max,
		seen:      make(map[string]bool),
		self:      []string{"AA"},
		elephants: []string{"AA"},
	}
}

func (s *state) flow() int {
	return s.tunnel.totalFlow(26, s.self...) + s.tunnel.totalFlow(26, s.elephants...)
}

func (s *state) addSelf(n string) bool {
	if s.seen[n] {
		return false
	}
	// w := s.tunnel.weights[s.self[len(s.self)-1]][n]
	w := s.tunnel.flatWeights[s.self[len(s.self)-1]+n]
	if len(s.self)+1+s.selfWeight+w > s.max {
		return false
	}
	s.self = append(s.self, n)
	s.selfWeight += w
	s.seen[n] = true
	return true
}

func (s *state) addElephant(n string) bool {
	if s.seen[n] {
		return false
	}
	// w := s.tunnel.weights[s.elephants[len(s.elephants)-1]][n]
	w := s.tunnel.flatWeights[s.elephants[len(s.elephants)-1]+n]
	if len(s.elephants)+1+s.elephantsWeight+w > s.max {
		return false
	}
	s.elephants = append(s.elephants, n)
	s.elephantsWeight += w
	s.seen[n] = true
	return true
}

func (s *state) popSelf() {
	last := s.self[len(s.self)-1]
	s.selfWeight -= s.tunnel.weights[s.self[len(s.self)-2]][last]
	s.self = s.self[:len(s.self)-1]
	s.seen[last] = false
}
func (s *state) popElephant() {
	last := s.elephants[len(s.elephants)-1]
	s.elephantsWeight -= s.tunnel.weights[s.elephants[len(s.elephants)-2]][last]
	s.elephants = s.elephants[:len(s.elephants)-1]
	s.seen[last] = false
}

func (t minTunnel) findOptimal2(s *state) int {
	startTime := time.Now()
	best := 0

	for to := range t.weights[s.self[len(s.self)-1]] {
		if !s.addSelf(to) {
			continue
		}
		for to := range t.weights[s.elephants[len(s.elephants)-1]] {
			if !s.addElephant(to) {
				continue
			}
			score := t.findOptimal2(s)
			if score > best {
				best = score
			}
			s.popElephant()
		}
		s.popSelf()
		if len(s.self) == 1 {
			fmt.Println(to, time.Since(startTime))
			startTime = time.Now()
		}
	}
	if best == 0 {
		f := s.flow()
		return f
	}
	return best
}

type Room struct {
	flow        int
	connections []string
}

func parseTunnel(input string) Tunnel {
	t := make(Tunnel)
	for _, line := range strings.Split(input, "\n") {
		name, room := parseRoom(line)
		t[name] = room
	}
	return t
}

var roomRegex = regexp.MustCompile(`Valve ([A-Z]{2}) has flow rate=([0-9]+); tunnels? leads? to valves? (.*)`)

func parseRoom(input string) (string, Room) {
	parts := roomRegex.FindStringSubmatch(input)
	name := parts[1]
	flow, _ := strconv.Atoi(parts[2])
	connections := parts[3]
	room := Room{flow: flow, connections: strings.Split(connections, ", ")}
	return name, room
}

type result struct {
	to    string
	steps int
	seen  map[string]bool
}
