package main

import (
	_ "embed"
	"fmt"
	"os"
	"regexp"
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
	fmt.Println(m.findOptimal2([]string{"AA"}, []string{"AA"}))

	tunnel = parseTunnel(input)
	puml = PrintActivity(tunnel)
	os.WriteFile("output.puml", puml, 0644)
	m = tunnel.toMinTunnel()
	os.WriteFile("outputMin.puml", minStateGraph(m), 0644)
	fmt.Println(m.findOptimal([]string{"AA"}))
	//1606 is too low
	//[AA UK CO IJ NA SE KF CS MN] 1862
	fmt.Println(m.findOptimal2([]string{"AA"}, []string{"AA"}))
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
	m := minTunnel{weights: make(map[string]map[string]int), flows: make(map[string]int)}
	m.weights["AA"] = map[string]int{}
	for _, n := range nodes {
		m.flows[n] = t[n].flow
		m.weights[n] = map[string]int{}
		m.weights["AA"][n] = t.distance("AA", n)
	}
	for i := 0; i < len(nodes)-1; i++ {
		for j := i + 1; j < len(nodes); j++ {
			m.weights[nodes[i]][nodes[j]] = t.distance(nodes[i], nodes[j])
			m.weights[nodes[j]][nodes[i]] = t.distance(nodes[i], nodes[j])
		}
	}
	return m
}

type minTunnel struct {
	weights map[string]map[string]int
	flows   map[string]int
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

func (t minTunnel) findOptimal2(currentSelf, currentElephants []string) int {
	startTime := time.Now()
	best := 0
	seen := make(map[string]bool)
	for _, c := range currentSelf {
		seen[c] = true
	}
	for _, c := range currentElephants {
		if seen[c] && c != "AA" {
			panic("elephant in self")
		}
		seen[c] = true
	}
	fromSelf := currentSelf[len(currentSelf)-1]
	fromElephants := currentElephants[len(currentElephants)-1]
	for toSelf := range t.weights[fromSelf] {
		if len(currentSelf) == 1 {
			now := time.Now()
			fmt.Println(toSelf, now.Sub(startTime))
			startTime = now
		}
		if seen[toSelf] {
			continue
		}
		seen[toSelf] = true
		currentSelf = append(currentSelf, toSelf)
		if len(currentSelf)+t.cumWeight(currentSelf...) > 26 {
			currentSelf = currentSelf[:len(currentSelf)-1]
			seen[toSelf] = false
			continue
		}
		for toElephants := range t.weights[fromElephants] {
			if seen[toElephants] {
				continue
			}
			currentElephants = append(currentElephants, toElephants)
			if len(currentElephants)+t.cumWeight(currentElephants...) > 26 {
				currentElephants = currentElephants[:len(currentElephants)-1]
				continue
			}
			score := t.findOptimal2(currentSelf, currentElephants)
			if score > best {
				best = score
			}
			currentElephants = currentElephants[:len(currentElephants)-1]
		}

		currentSelf = currentSelf[:len(currentSelf)-1]
		seen[toSelf] = false
	}

	if best == 0 {
		return t.totalFlow(26, currentSelf...) + t.totalFlow(26, currentElephants...)
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
