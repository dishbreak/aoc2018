package main

import (
	"fmt"
	"image"
	"runtime"
	"sort"
	"sync"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day6.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
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

func closestToAll(l locus, input <-chan image.Point, cutoff int) <-chan image.Point {
	output := make(chan image.Point)

	var wg sync.WaitGroup
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for loc := range input {
				acc := 0
				for _, pt := range l.pts {
					acc += dist(loc, pt)
				}
				if acc >= cutoff {
					continue
				}
				output <- loc
			}
		}()
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

func part2(input []string) int {
	return safeRegionSize(input, 10000)
}

func safeRegionSize(input []string, cutoff int) int {
	l := toLocus(input)
	locs := generateLocations(l)
	safestPts := closestToAll(l, locs, cutoff)

	acc := 0
	for range safestPts {
		acc++
	}
	return acc
}
