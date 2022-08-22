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
}

type guardRecord struct {
	id         int
	starts     []int
	ends       []int
	minsAsleep int
}

func part1(input []string) int {
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
			pts := guardShift.FindStringSubmatch(line)
			start, _ := strconv.Atoi(pts[1])
			g.starts = append(g.starts, start)
			lastStart = start
		case wakeUp.MatchString(line):
			pts := guardShift.FindStringSubmatch(line)
			end, _ := strconv.Atoi(pts[1])
			g.ends = append(g.ends, end)
			g.minsAsleep += end - lastStart
		}
	}

	return 0
}

func mostOftenAsleepAt(g *guardRecord) int {
	sort.IntSlice(g.starts)
	sort.IntSlice(g.ends)
}
