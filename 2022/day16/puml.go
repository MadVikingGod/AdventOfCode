package main

import (
	"bytes"
	"fmt"
)

func PrintActivity(t Tunnel) []byte {
	buf := bytes.Buffer{}
	buf.WriteString("@startuml\n\n[*] -> AA\n\n")
	seen := make(map[string]bool)
	queue := []string{"AA"}
	for len(queue) > 0 {
		name := queue[0]
		queue = queue[1:]
		if seen[name] {
			continue
		}
		seen[name] = true
		room := t[name]
		if room.flow > 0 {
			buf.WriteString(fmt.Sprintf("%s : %d\n", name, room.flow))
		}

		for _, c := range room.connections {
			if !seen[c] {
				buf.WriteString(fmt.Sprintf("%s --> %s\n", name, c))
			}
			queue = append(queue, c)
		}
	}

	buf.WriteString("\n@enduml\n")

	return buf.Bytes()
}

func minStateGraph(t minTunnel) []byte {
	buf := bytes.Buffer{}
	buf.WriteString("@startuml\n\n[*] -> AA\n\n")

	for n2, w := range t.weights["AA"] {
		buf.WriteString(fmt.Sprintf("AA --> %s : %d\n", n2, w))
	}

	seen := map[string]bool{"AA": true}
	for n1, m := range t.weights {
		if n1 == "AA" {
			continue
		}
		seen[n1] = true
		for n2, w := range m {
			if !seen[n2] {
				buf.WriteString(fmt.Sprintf("%s --> %s : %d\n", n1, n2, w))
			}
		}
	}

	for name, flow := range t.flows {
		buf.WriteString(fmt.Sprintf("%s : %d\n", name, flow))
	}

	buf.WriteString("\n@enduml\n")

	return buf.Bytes()
}
