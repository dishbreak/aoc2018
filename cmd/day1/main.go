package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day1.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) (acc int) {

	for _, item := range input {
		if item == "" {
			continue
		}
		item = strings.TrimPrefix(item, "+")
		delta, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		acc += delta
	}
	return
}

func part2(input []string) (acc int) {

	hitMap := make(map[int]bool)

	for i := 0; ; i++ {
		item := input[i%len(input)]
		if item == "" {
			continue
		}
		item = strings.TrimPrefix(item, "+")
		delta, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		acc += delta
		if hitMap[acc] {
			return
		}
		hitMap[acc] = true
	}
}
