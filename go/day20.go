package main

import "fmt"

type room struct {
	north, east, south, west bool
}

type elfComplex struct {
	building         map[xy]*room
	traversingHelper map[rune]xy
}

func day20() {
	complex := elfComplex{}
	complex.init()
	longestShortestPath, numShortestPathsLongerThan1000 := complex.getLongestShortestPath()
	fmt.Printf("Longest shortest path is %d.\n", longestShortestPath)
	fmt.Printf("Number of shortest paths longer than 1000 is %d.\n", numShortestPathsLongerThan1000)
}

func (complex *elfComplex) init() {
	regex := readFile("../input/day20.txt")[0]
	building := map[xy]*room{}
	traversingHelper := map[rune]xy{'N': xy{0, -1}, 'E': xy{1, 0}, 'S': xy{0, 1}, 'W': xy{-1, 0}}
	curCoord := xy{0, 0}
	building[curCoord] = &room{}
	stk := xyStack{}
	for _, c := range regex {
		switch c {
		case '$':
		case '^':
		case '(':
			stk = stk.Push(curCoord)
		case ')':
			stk, curCoord = stk.Pop()
		case '|':
			// I'm too lazy to implement stack.Peek()
			stk, curCoord = stk.Pop()
			stk = stk.Push(curCoord)
		default:
			newCoord := xy{traversingHelper[c].x + curCoord.x, traversingHelper[c].y + curCoord.y}
			r := building[newCoord]
			if r == nil {
				r = &room{}
				building[newCoord] = r
			}
			old := building[curCoord]
			if c == 'E' {
				r.west = true
				old.east = true
			} else if c == 'N' {
				r.south = true
				old.north = true
			} else if c == 'W' {
				r.east = true
				old.west = true
			} else if c == 'S' {
				r.north = true
				old.south = true
			} else {
				panic("Bad input.")
			}
			curCoord = newCoord
		}
	}
	complex.building = building
	complex.traversingHelper = traversingHelper
}

func (complex elfComplex) getLongestShortestPath() (int, int) {
	// Compute shortest path for each node 😭(V^3)
	// Research on wikipedia has led me to believe that this is the best way to do this unfortunately https://en.wikipedia.org/wiki/Floyd–Warshall_algorithm
	longestPath := 0
	numLongerThan1000 := 0
	// longestCoord := xy{0, 0}
	for k := range complex.building {
		queue := []xyNode{xyNode{xy{0, 0}, 0}}
		visited := map[xy]bool{}
		for len(queue) > 0 { // Most likely gonna break out of this earlier b/c all of the weights of the edges are 1
			curCoord := queue[0]
			curRoom := complex.building[curCoord.c]

			if curCoord.c == k {
				// Found the room
				if longestPath < curCoord.d {
					longestPath = curCoord.d
					// longestCoord = curCoord.c
				}
				if curCoord.d >= 1000 {
					numLongerThan1000++
				}
				break // break out of while loop early cause why not
			}

			queue = queue[1:]
			if curRoom.east {
				n := xyNode{xy{curCoord.c.x + 1, curCoord.c.y}, curCoord.d + 1}
				if !visited[n.c] {
					queue = append(queue, n)
					visited[n.c] = true
				}
			}
			if curRoom.west {
				n := xyNode{xy{curCoord.c.x - 1, curCoord.c.y}, curCoord.d + 1}
				if !visited[n.c] {
					queue = append(queue, n)
					visited[n.c] = true
				}
			}
			if curRoom.north {
				n := xyNode{xy{curCoord.c.x, curCoord.c.y - 1}, curCoord.d + 1}
				if !visited[n.c] {
					queue = append(queue, n)
					visited[n.c] = true
				}
			}
			if curRoom.south {
				n := xyNode{xy{curCoord.c.x, curCoord.c.y + 1}, curCoord.d + 1}
				if !visited[n.c] {
					queue = append(queue, n)
					visited[n.c] = true
				}
			}
		}
	}
	return longestPath, numLongerThan1000
}
