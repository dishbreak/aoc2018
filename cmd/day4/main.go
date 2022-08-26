package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day4.txt")
	if err != nil {
		panic(err)
	}

	sort.Strings(input)
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

type guardRecord struct {
	id         int
	starts     []int
	ends       []int
	minsAsleep int
}

func part1(input []string) int {
	barracks := parseBarracks(input)
	var g *guardRecord
	max := -1
	for _, val := range barracks {
		if max < val.minsAsleep {
			max = val.minsAsleep
			g = val
		}
	}

	m, _ := sleepiestMinute(g)

	return g.id * m
}

func parseBarracks(input []string) map[int]*guardRecord {
	var g *guardRecord
	var lastStart int

	barracks := make(map[int]*guardRecord)

	var guardShift = regexp.MustCompile(`^\[1518-\d\d-\d\d \d\d:\d\d] Guard #(\d+) begins shift$`)
	var fallAsleep = regexp.MustCompile(`^\[1518-\d\d-\d\d \d\d:(\d\d)] falls asleep$`)
	var wakeUp = regexp.MustCompile(`^\[1518-\d\d-\d\d \d\d:(\d\d)] wakes up$`)

	for _, line := range input {
		switch {
		case line == "":
			continue
		case guardShift.MatchString(line):
			pts := guardShift.FindStringSubmatch(line)
			id, err := strconv.Atoi(pts[1])
			if err != nil {
				panic(err)
			}

			var ok bool
			g, ok = barracks[id]
			if !ok {
				g = &guardRecord{
					id:     id,
					starts: make([]int, 0),
					ends:   make([]int, 0),
				}
				barracks[id] = g
			}
		case fallAsleep.MatchString(line):
			pts := fallAsleep.FindStringSubmatch(line)
			start, _ := strconv.Atoi(pts[1])
			g.starts = append(g.starts, start)
			lastStart = start
		case wakeUp.MatchString(line):
			pts := wakeUp.FindStringSubmatch(line)
			end, _ := strconv.Atoi(pts[1])
			g.ends = append(g.ends, end)
			g.minsAsleep += end - lastStart
		}
	}

	for id, g := range barracks {
		if len(g.starts) == 0 || len(g.ends) == 0 {
			delete(barracks, id)
		}
	}

	return barracks
}

func sleepiestMinute(g *guardRecord) (int, int) {
	sort.Ints(g.starts)
	sort.Ints(g.ends)

	starts := make([]int, len(g.starts))
	copy(starts, g.starts)

	ends := make([]int, len(g.ends))
	copy(ends, g.ends)

	acc := 0
	start, end := starts[0], ends[0]
	starts, ends = starts[1:], ends[1:]

	max, atMin := -1, -1
	for {
		if start < end {
			acc++
			if acc > max {
				max = acc
				atMin = start
			}
			if len(starts) == 0 {
				break
			}
			start = starts[0]
			starts = starts[1:]
			continue
		}
		acc--
		end = ends[0]
		ends = ends[1:]
	}

	return atMin, acc
}

func part2(input []string) int {
	barracks := parseBarracks(input)

	var g *guardRecord
	minute := 0
	max := -1
	for _, val := range barracks {
		m, count := sleepiestMinute(val)
		if count > max {
			g = val
			max = count
			minute = m
		}
	}

	return g.id * minute
}
