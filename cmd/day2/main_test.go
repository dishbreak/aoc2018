package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}

	assert.Equal(t, 12, part1(input))
}

func TestAdjacent(t *testing.T) {
	type testCase struct {
		one, other string
		idx        int
		adj        bool
	}

	testCases := []testCase{
		{"axcye", "abcde", 1, false},
		{"fguij", "fghij", 2, true},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			adj, idx := adjacent(tc.one, tc.other)
			assert.Equal(t, tc.idx, idx)
			assert.Equal(t, tc.adj, adj)
		})
	}
}
