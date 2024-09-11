package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day7.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %s\n", part1(input))
}

type node struct {
	label string
	next  map[string]*node
	prev  map[string]*node
}

type graph struct {
	nodes map[string]*node
	start []*node
}

func toGraph(input []string) *graph {
	g := &graph{
		nodes: make(map[string]*node),
		start: make([]*node, 0),
	}

	noDeps := make(map[string]bool)

	type rule struct {
		dependent, dependency string
	}

	rules := make([]rule, 0)
	for _, line := range input {
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		rules = append(rules, rule{
			dependent:  parts[7],
			dependency: parts[1],
		})
		noDeps[parts[1]] = true
		noDeps[parts[7]] = true
	}

	for l := range noDeps {
		g.nodes[l] = &node{
			label: l,
			next:  make(map[string]*node),
			prev:  make(map[string]*node),
		}
	}

	for _, r := range rules {
		delete(noDeps, r.dependent)
		dependency := g.nodes[r.dependency]
		dependent := g.nodes[r.dependent]
		dependency.next[dependent.label] = dependent
		dependent.prev[dependency.label] = dependency
	}

	for d := range noDeps {
		g.start = append(g.start, g.nodes[d])
	}

	return g
}

func part1(input []string) string {
	var sb strings.Builder

	g := toGraph(input)
	q := make([]*node, 0)
	q = append(q, g.start...)

	for len(q) > 0 {
		sort.Slice(q, func(i, j int) bool {
			return q[i].label < q[j].label
		})

		n := q[0]
		q = q[1:]

		sb.WriteString(n.label)
		for _, c := range n.next {
			delete(c.prev, n.label)
			// if there's still dependencies, we can't enqueue this child yet
			if len(c.prev) > 0 {
				continue
			}
			q = append(q, c)
		}
	}

	return sb.String()
}
