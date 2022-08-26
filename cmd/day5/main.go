package main

import (
	"fmt"
	"sync"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day5.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input[0]))
	fmt.Printf("Part 2: %d\n", part2(input[0]))
}

func part1(input string) int {
	return reactPolymer([]byte(input))
}

func reactPolymer(units []byte) int {
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

func part2(input string) int {
	units := []byte(input)

	results := make(chan int)

	var wg sync.WaitGroup
	for t := 'A'; t <= 'Z'; t++ {
		wg.Add(1)
		go func(t byte) {
			defer wg.Done()
			chain := make([]byte, 0)

			for _, b := range units {
				if b == t || b == t+32 {
					continue
				}
				chain = append(chain, b)
			}

			results <- reactPolymer(chain)
		}(byte(t))
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	min := len(units)
	for val := range results {
		if val < min {
			min = val
		}
	}

	return min
}
