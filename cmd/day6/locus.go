package main

import (
	"image"
	"strconv"
	"strings"
)

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
