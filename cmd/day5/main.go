package main

import (
	"fmt"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day5.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input[0]))
}

func part1(input string) int {
	units := []byte(input)

	for changed := true; changed; {
		changed = false
		skips := make([]bool, len(units))

		for i := range units {
			if i == 0 {
				continue
			}
			if d := int(units[i]) - int(units[i-1]); d == 32 || d == -32 {
				skips[i], skips[i-1] = true, true
				changed = true
				break
			}
		}

		if !changed {
			continue
		}

		newUnits := make([]byte, 0)
		for i := range units {
			if skips[i] {
				continue
			}
			newUnits = append(newUnits, units[i])
		}
		units = newUnits
	}

	return len(units)
}
