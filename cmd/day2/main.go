package main

import (
	"fmt"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %s\n", part2(input))
}

func part1(input []string) int {
	var count2, count3 int
	for _, item := range input {
		if item == "" {
			continue
		}
		c2, c3 := checkSum(item)
		count2 += c2
		count3 += c3
	}
	return count2 * count3
}

func checkSum(name string) (int, int) {
	charCounts := make([]int, 26)
	for _, c := range name {
		charCounts[c-'a']++
	}

	counterCounts := make([]int, len(name))
	for _, ct := range charCounts {
		counterCounts[ct]++
	}

	if counterCounts[2] > 1 {
		counterCounts[2] = 1
	}

	if counterCounts[3] > 1 {
		counterCounts[3] = 1
	}

	return counterCounts[2], counterCounts[3]
}

func part2(input []string) string {
	for i := range input[:len(input)-1] {
		for j := range input[:len(input)-1] {
			if i == j {
				continue
			}
			adj, idx := adjacent(input[i], input[j])
			if !adj {
				continue
			}
			return input[i][:idx] + input[i][idx+1:]
		}
	}
	return "n/a"
}

func adjacent(one, other string) (adj bool, idx int) {
	idx = -1
	for i := range one {
		if one[i] == other[i] {
			continue
		}
		idx = i
		break
	}

	if idx == -1 {
		return
	}

	adj = one[idx+1:] == other[idx+1:]
	return
}
