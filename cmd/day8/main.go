package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func getInt(i *bufio.Scanner) int {
	word := i.Text()
	num, err := strconv.Atoi(word)
	if err != nil {
		panic(fmt.Errorf("failed to convert token '%s' to int: %w", word, err))
	}

	return num
}

func main() {
	f, err := os.Open("inputs/day8.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))
	f.Seek(0, 0)
	fmt.Printf("Part 2: %d\n", part2(f))
}

func part1(r io.Reader) (acc int) {
	i := bufio.NewScanner(r)
	i.Split(bufio.ScanWords)

	acc = countNodes(i)

	return acc
}

func part2(r io.Reader) (acc int) {
	i := bufio.NewScanner(r)
	i.Split(bufio.ScanWords)

	acc = evaluateNode(i)

	return acc
}

func countNodes(i *bufio.Scanner) (acc int) {
	scan := func() {
		ok := i.Scan()
		if !ok {
			panic(errors.New("unexpected end of string"))
		}
	}

	scan()
	childCt := getInt(i)
	scan()
	metaCt := getInt(i)

	for n := 0; n < childCt; n++ {
		acc += countNodes(i)
	}

	for n := 0; n < metaCt; n++ {
		scan()
		acc += getInt(i)
	}
	return acc
}

func evaluateNode(i *bufio.Scanner) (acc int) {
	scan := func() {
		ok := i.Scan()
		if !ok {
			panic(errors.New("unexpected end of string"))
		}
	}

	scan()
	childCt := getInt(i)
	scan()
	metaCt := getInt(i)

	if childCt == 0 {
		for n := 0; n < metaCt; n++ {
			scan()
			acc += getInt(i)
		}
		return acc
	}

	children := make([]int, childCt)
	for n := 0; n < childCt; n++ {
		children[n] = evaluateNode(i)
	}

	for n := 0; n < metaCt; n++ {
		scan()
		idx := getInt(i)
		if idx == 0 || idx > len(children) {
			continue
		}

		acc += children[idx-1]
	}

	return acc
}
