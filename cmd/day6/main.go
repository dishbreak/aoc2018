package main

import (
	"fmt"
	"image"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day6.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
}

type locus struct {
	min, max image.Point
	pts      []image.Point
}

func (l locus) OnBoundary(p image.Point) bool {
	return p.X == l.min.X || p.X == l.max.X || p.Y == l.min.Y || p.Y == l.max.Y
}

func toLocus(input []string) locus {
	l := locus{
		min: image.Pt(10000, 10000),
		max: image.Pt(0, 0),
		pts: make([]image.Point, 0),
	}

	for _, line := range input {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ", ")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		p := image.Pt(x, y)
		l.pts = append(l.pts, p)

		if l.min.X > x {
			l.min.X = x
		} else if l.max.X < x {
			l.max.X = x
		}

		if l.min.Y > y {
			l.min.Y = y
		} else if l.max.Y < y {
			l.max.Y = y
		}

	}

	return l
}

func generateLocations(l locus) <-chan image.Point {
	output := make(chan image.Point)

	go func() {
		defer close(output)
		for x := l.min.X; x <= l.max.X; x++ {
			for y := l.min.Y; y <= l.max.Y; y++ {
				p := image.Pt(x, y)
				output <- p
			}
		}
	}()

	return output
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func dist(a, b image.Point) int {
	v := b.Sub(a)
	return abs(v.X) + abs(v.Y)
}

type distRecord struct {
	dist int
	pt   image.Point
	loc  image.Point
}

func nearestPoint(l locus, input <-chan image.Point) <-chan distRecord {
	output := make(chan distRecord)

	var wg sync.WaitGroup
	for i := 0; i < runtime.NumCPU()-1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for loc := range input {
				records := make([]distRecord, 0)
				for _, p := range l.pts {
					records = append(records, distRecord{dist(loc, p), p, loc})
				}
				sort.Slice(records, func(i, j int) bool {
					return records[i].dist < records[j].dist
				})

				//if there's a tie for distances, we don't have a winner.
				if records[0].dist == records[1].dist {
					continue
				}
				output <- records[0]
			}
		}()
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

func largestArea(l locus, input <-chan distRecord) <-chan int {
	output := make(chan int)

	go func() {
		defer close(output)
		hits := make(map[image.Point]int)
		infinitePt := make(map[image.Point]bool)

		for _, p := range l.pts {
			if l.OnBoundary(p) {
				infinitePt[p] = true
			}
		}

		for p := range input {
			hits[p.pt]++
			if l.OnBoundary(p.loc) {
				infinitePt[p.pt] = true
			}
		}

		maxArea := -1
		for p, area := range hits {
			if infinitePt[p] {
				continue
			}
			if maxArea < area {
				maxArea = area
			}
		}

		output <- maxArea
	}()

	return output
}

func part1(input []string) int {
	l := toLocus(input)
	locs := generateLocations(l)
	nearestPts := nearestPoint(l, locs)
	result := largestArea(l, nearestPts)

	return <-result
}
