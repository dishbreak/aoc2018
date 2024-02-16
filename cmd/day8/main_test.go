package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"
	r := strings.NewReader(input)

	assert.Equal(t, 138, part1(r))
}
