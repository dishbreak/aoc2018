package main

import (
	"fmt"
	"image"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day3.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) (acc int) {
	space := make(map[image.Point]int)
	claims := make([]claim, len(input)-1)
	for i, item := range input[:len(input)-1] {
		c := parseClaim(item)
		claims[i] = c
	}

	for _, c := range claims {
		for x := c.origin.X; x < c.max.X; x++ {
			for y := c.origin.Y; y < c.max.Y; y++ {
				space[image.Pt(x, y)]++
			}
		}
	}

	for _, ct := range space {
		if ct > 1 {
			acc++
		}
	}

	return
}

func part2(input []string) int {
	space := make(map[image.Point]int)
	claims := make([]claim, len(input)-1)
	for i, item := range input[:len(input)-1] {
		c := parseClaim(item)
		claims[i] = c
	}

	for _, c := range claims {
		for x := c.origin.X; x < c.max.X; x++ {
			for y := c.origin.Y; y < c.max.Y; y++ {
				space[image.Pt(x, y)]++
			}
		}
	}

	isIntact := func(c claim) bool {
		for x := c.origin.X; x < c.max.X; x++ {
			for y := c.origin.Y; y < c.max.Y; y++ {
				if space[image.Pt(x, y)] != 1 {
					return false
				}
			}
		}
		return true
	}

	for _, c := range claims {
		if isIntact(c) {
			return c.id
		}
	}

	return -1
}

type claim struct {
	origin image.Point
	max    image.Point
	id     int
}

func parseClaim(input string) (c claim) {
	parts := strings.Fields(input)

	c.id, _ = strconv.Atoi(strings.TrimPrefix(parts[0], "#"))
	c.origin = toPoint(strings.TrimSuffix(parts[2], ":"), ",")
	c.max = c.origin.Add(toPoint(parts[3], "x"))
	return
}

func toPoint(input, delim string) (p image.Point) {
	parts := strings.Split(input, delim)
	p.X, _ = strconv.Atoi(parts[0])
	p.Y, _ = strconv.Atoi(parts[1])
	return
}
