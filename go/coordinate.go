package main

type xy struct {
	x, y int
}

type xyz struct {
	x, y, z int
}

type xyNode struct {
	c xy
	d int
}

type xyGraph map[xy][]xy

type xyStack []xy

func (s xyStack) Push(c xy) xyStack {
	return append(s, c)
}

func (s xyStack) Pop() (xyStack, xy) {
	l := len(s)
	return s[:l-1], s[l-1]
}
