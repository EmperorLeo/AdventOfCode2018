package main

import (
	"fmt"
	"strconv"
)

type lumberCollectionArea struct {
	state [][]rune
	round int
}

func day18() {
	leosWoods := lumberCollectionArea{}
	leosWoods.init()
	leosWoods.print()

	// Part 1
	for i := 0; i < 10; i++ {
		leosWoods.getNext()
		leosWoods.print()
		open, trees, lumberyards := leosWoods.getResourceCounts()
		fmt.Printf("Open = %d, Trees = %d, Lumberyards = %d.  Trees * Lumberyards = %d\n\n", open, trees, lumberyards, trees*lumberyards)
	}

	fmt.Print("^^^^ PART 1 ANSWER RIGHT ABOVE ^^^^\n\n\n\n\n\n\n")

	// Part 2 - I determined that the magic pattern is a cycle every 35 years
	var lastOpen, lastTrees, lastLumberyards int
	// 1000000000 mod 35 = 20, and im subtracting 10 from the last loop.  Also multiplying 35 by 50 to get far enough to cycle.
	for i := 0; i < 35*50+(20-10); i++ {
		// Continue the loop for 990 years to find a pattern
		leosWoods.getNext()
		// leosWoods.print()
		open, trees, lumberyards := leosWoods.getResourceCounts()
		fmt.Print(leosWoods.round)
		fmt.Println()
		fmt.Printf("Open Diff = %d, Trees Diff = %d, Lumberyards Diff = %d.  Trees * Lumberyards = %d\n", open-lastOpen, trees-lastTrees, lumberyards-lastLumberyards, trees*lumberyards)
		lastOpen, lastTrees, lastLumberyards = open, trees, lumberyards
	}
}

func (l *lumberCollectionArea) init() {
	lines := readFile("../input/day18.txt")
	l.state = make([][]rune, len(lines))
	for y, line := range lines {
		l.state[y] = []rune(line)
	}
}

func (l *lumberCollectionArea) getNext() {
	l.round++
	nextState := make([][]rune, len(l.state))
	for y, row := range l.state {
		nextState[y] = make([]rune, len(row))
		for x, item := range row {
			var openCount, treesCount, lumberyardCount int
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if !(i == 0 && j == 0) && l.inBounds(x+i, y+j) {
						adjacent := l.state[y+j][x+i]
						switch adjacent {
						case '.':
							openCount++
						case '#':
							lumberyardCount++
						case '|':
							treesCount++
						}
					}
				}
			}
			switch item {
			case '.':
				if treesCount >= 3 {
					nextState[y][x] = '|'
				} else {
					nextState[y][x] = '.'
				}
			case '#':
				if lumberyardCount >= 1 && treesCount >= 1 {
					nextState[y][x] = '#'
				} else {
					nextState[y][x] = '.'
				}
			case '|':
				if lumberyardCount >= 3 {
					nextState[y][x] = '#'
				} else {
					nextState[y][x] = '|'
				}
			}
		}
	}
	l.state = nextState
}

func (l *lumberCollectionArea) inBounds(x, y int) bool {
	return x >= 0 && y >= 0 && y < len(l.state) && x < len(l.state[y])
}

func (l *lumberCollectionArea) getResourceCounts() (int, int, int) {
	var openCount, treesCount, lumberyardCount int
	for _, row := range l.state {
		for _, item := range row {
			switch item {
			case '.':
				openCount++
			case '#':
				lumberyardCount++
			case '|':
				treesCount++
			}
		}
	}
	return openCount, treesCount, lumberyardCount
}

func (l *lumberCollectionArea) print() {
	fmt.Printf("After %d minutes:\n", l.round)
	for y, row := range l.state {
		if y == 0 {
			xAxis := "   "
			for i := 0; i < len(row); i++ {
				xAxis += strconv.Itoa(i % 10)
			}
			fmt.Print(xAxis)
			fmt.Println()
		}
		strToPrint := strconv.Itoa(y) + " "
		if y < 10 {
			strToPrint += " "
		}
		for _, item := range row {
			strToPrint += string(item)
		}
		fmt.Print(strToPrint)
		fmt.Println()
	}
	fmt.Println()
}
