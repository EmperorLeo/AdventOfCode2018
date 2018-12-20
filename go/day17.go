package main

import (
	"fmt"
	"math"
)

type stack []*xy

type xy struct {
	x, y int
}

type ground struct {
	grid          [][]rune
	xOffset, minY int
	stk           stack
}

func day17() {
	g := ground{}
	g.init()
	g.startFlow()
	g.print()
	still, running := g.getNumberWaterUnits()

	// Part 1
	fmt.Printf("%d water units.\n", still+running)

	// Part 2
	fmt.Printf("%d water units left after all the spring stops running.\n", still)
}

func (g *ground) init() {
	lines := readFile("../input/day17.txt")
	clay := []xy{}
	var minX, maxX, minY, maxY int
	minX, minY = math.MaxInt32, math.MaxInt32
	for _, line := range lines {
		var l1, l2 string
		var n1, n2, n3 int
		fmt.Sscanf(line, "%1s=%d, %1s=%d..%d", &l1, &n1, &l2, &n2, &n3)
		if l1 == "y" {
			for i := n2; i <= n3; i++ {
				if n1 > maxY {
					maxY = n1
				}
				if n1 < minY {
					minY = n1
				}
				if i > maxX {
					maxX = i
				}
				if i < minX {
					minX = i
				}
				clay = append(clay, xy{i, n1})
			}
		} else {
			for i := n2; i <= n3; i++ {
				if i > maxY {
					maxY = i
				}
				if i < minY {
					minY = i
				}
				if n1 > maxX {
					maxX = n1
				}
				if n1 < minX {
					minX = n1
				}
				clay = append(clay, xy{n1, i})
			}
		}
	}

	g.xOffset = 1 - minX
	g.minY = minY
	g.stk = stack{}
	g.grid = make([][]rune, maxY+1)
	for i := range g.grid {
		g.grid[i] = make([]rune, maxX-minX+3)
		for j := range g.grid[i] {
			g.grid[i][j] = '.'
		}
	}
	for _, coord := range clay {
		g.grid[coord.y][coord.x+g.xOffset] = '#'
	}
}

func (g *ground) startFlow() {
	var currentCoord *xy
	g.stk = g.stk.Push(&xy{500 + g.xOffset, 1})
	g.grid[0][500+g.xOffset] = '+'
	var counter int
	for len(g.stk) > 0 {
		counter++
		g.stk, currentCoord = g.stk.Pop()
		x, y := currentCoord.x, currentCoord.y
		if counter > 1000000 {
			fmt.Printf("Stuck coordinate is (%d, %d)\n", x, y)
		}
		char := g.grid[y][x]
		if char == '.' {
			g.grid[y][x] = '|'
			if len(g.grid) <= y+1 || g.grid[y+1][x] == '|' {
				// Out of bounds, or just hit more running water, so get the next item on the stack
			} else if g.grid[y+1][x] == '#' || g.grid[y+1][x] == '~' {
				// Hit some clay or some still water, so should fill up more
				g.fill(currentCoord)
			} else { // water is still going down
				g.stk = g.stk.Push(currentCoord)
				g.stk = g.stk.Push(&xy{x, y + 1})
			}
		} else if char == '|' {
			if len(g.grid) > y+1 && g.grid[y+1][x] != '|' {
				g.fill(currentCoord)
			}
		}
	}
}

func (g *ground) fill(coord *xy) {
	defer func() {
		if r := recover(); r != nil {
			g.print()
			panic(r)
		}
	}()
	x, y := coord.x, coord.y
	leftBound, rightBound := x, x
	var hitLeft, hitRight bool
	for {
		if g.grid[y+1][leftBound] != '#' && g.grid[y+1][leftBound] != '~' {
			g.stk = g.stk.Push(&xy{leftBound, y})
			leftBound++
			break
		}
		if g.grid[y][leftBound-1] == '#' {
			hitLeft = true
			break
		}
		leftBound--
	}
	for {
		if g.grid[y+1][rightBound] != '#' && g.grid[y+1][rightBound] != '~' {
			g.stk = g.stk.Push(&xy{rightBound, y})
			rightBound--
			break
		}
		if g.grid[y][rightBound+1] == '#' {
			hitRight = true
			break
		}
		rightBound++
	}
	var fillChar rune
	if hitLeft && hitRight {
		fillChar = '~'
	} else {
		fillChar = '|'
	}
	for i := leftBound; i <= rightBound; i++ {
		g.grid[y][i] = fillChar
	}
}

func (g *ground) print() {
	for _, row := range g.grid {
		strToPrint := ""
		for _, val := range row {
			strToPrint += string(val)
		}
		fmt.Print(strToPrint)
		fmt.Println()
	}
}

func (g *ground) getNumberWaterUnits() (int, int) {
	var stillCount int
	var runningCount int
	for y, row := range g.grid {
		for _, val := range row {
			if g.minY <= y {
				if val == '|' {
					runningCount++
				}
				if val == '~' {
					stillCount++
				}
			}
		}
	}
	return stillCount, runningCount
}

func (s stack) Push(coord *xy) stack {
	return append(s, coord)
}

func (s stack) Pop() (stack, *xy) {
	l := len(s)
	return s[:l-1], s[l-1]
}
