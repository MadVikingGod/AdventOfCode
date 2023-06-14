package main

import (
	"regexp"
	"strings"
)

func main() {

}

func findIngredient(ings, contents []stringSet) ([]stringSet, []stringSet, map[string]string) {
	knowIng := map[string]string{}
	for i, ing1 := range ings {
		for j := i + 1; j < len(ings); j++ {
			if ing1.intersection(ings[j]).len() == 1 && contents[i].intersection(contents[j]).len() == 1 {
				knowIng[ing1.intersection(ings[j]).top()] = contents[i].intersection(contents[j]).top()
			}
		}
	}
	return []stringSet{}, []stringSet{}, knowIng
}

var re = regexp.MustCompile(`(.*) \(contains (.*)\)`)

func parse(s string) (stringSet, stringSet) {
	match := re.FindStringSubmatch(s)
	ingredients := New(strings.Split(match[1], " ")...)
	allergens := New(strings.Split(match[2], ", ")...)
	return ingredients, allergens
}

type stringSet struct {
	set map[string]struct{}
}

func New(ss ...string) stringSet {
	set := stringSet{map[string]struct{}{}}
	for _, s := range ss {
		set.set[s] = struct{}{}
	}
	return set
}

func (s stringSet) has(in string) bool {
	_, ok := s.set[in]
	return ok
}

func (s stringSet) len() int {
	return len(s.set)
}

func (a stringSet) intersection(b stringSet) stringSet {
	inter := stringSet{map[string]struct{}{}}
	for s := range a.set {
		if _, ok := b.set[s]; ok {
			inter.set[s] = struct{}{}
		}
	}
	return inter
}
func (a stringSet) top() string {
	for k := range a.set {
		return k
	}
	return ""
}
func (a stringSet) remove(key string) {
	delete(a.set, key)
}
