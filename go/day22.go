package main

import (
	"fmt"
	"math"
)

type xyWithTool struct {
	c    xy
	tool int
}

type xyToolWithDepth struct {
	xyTool xyWithTool
	depth  int
}

type cave struct {
	erosionLevelMap [][]int
	depth           int
	target          xy
	tools           map[int][]bool
	edgeIterator    []xy
}

func day22() {
	target := xy{13, 704}
	depth := 9465
	smallestRegionSize := int(math.Max(float64(target.y), float64(target.x))) + 200 // add 200 to give yourself some space for part 2
	geoIndexMap := make([][]int, smallestRegionSize)
	for i := range geoIndexMap {
		geoIndexMap[i] = make([]int, smallestRegionSize)
		for j := range geoIndexMap[i] {
			geoIndexMap[i][j] = -1
		}
	}
	tools := map[int][]bool{
		0: []bool{false, true, true},
		1: []bool{true, true, false},
		2: []bool{true, false, true}}
	iter := []xy{xy{-1, 0}, xy{1, 0}, xy{0, -1}, xy{0, 1}}

	c := cave{geoIndexMap, depth, target, tools, iter}

	c.scanGeologicIndices()

	smallestRectRiskLevel := c.calculateSmallestRectangleRiskLevel()
	fmt.Printf("Smallest rectangle risk level is %d.\n", smallestRectRiskLevel)

	shortestPathRescueLength := c.getShortestRescuePathLength()
	fmt.Printf("Shortest path rescue length is %d.\n", shortestPathRescueLength)
	c.print()
}

func (c *cave) scanGeologicIndices() {
	for y := 0; y < len(c.erosionLevelMap); y++ {
		for x := 0; x < len(c.erosionLevelMap); x++ {
			var geologicIndex int
			if x == 0 && y == 0 {
				geologicIndex = 0
			} else if x == 0 {
				geologicIndex = y * 48271
			} else if y == 0 {
				geologicIndex = x * 16807
			} else if c.target.x == x && c.target.y == y {
				geologicIndex = 0
			} else {
				geologicIndex = c.erosionLevelMap[y-1][x] * c.erosionLevelMap[y][x-1]
			}

			// fmt.Printf("Geologic index at (%d, %d) = %d\n", x, y, geologicIndex)
			c.erosionLevelMap[y][x] = (geologicIndex + c.depth) % 20183
		}
	}
}

func (c *cave) calculateSmallestRectangleRiskLevel() int {
	totalRisk := 0
	for y := 0; y <= c.target.y; y++ {
		for x := 0; x <= c.target.x; x++ {
			totalRisk += c.erosionLevelMap[y][x] % 3
		}
	}
	return totalRisk
}

func (c *cave) getShortestRescuePathLength() int {
	shortestPath := math.MaxInt32
	queue := make([]xyToolWithDepth, 0)
	queue = append(queue, xyToolWithDepth{xyWithTool{xy{0, 0}, 2}, 0})
	visited := map[xyWithTool]int{xyWithTool{xy{0, 0}, 2}: -1}
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		edges := c.getEdges(curNode)
		for _, e := range edges {
			if visited[e.xyTool] == 0 || visited[e.xyTool] > e.depth {
				queue = append(queue, e)
				visited[e.xyTool] = e.depth
			}
		}
		if curNode.xyTool.c.x == c.target.x && curNode.xyTool.c.y == c.target.y && shortestPath > curNode.depth && curNode.xyTool.tool == 2 {
			shortestPath = curNode.depth
		}
	}
	return shortestPath
}

func (c *cave) getEdges(source xyToolWithDepth) []xyToolWithDepth {
	edges := make([]xyToolWithDepth, 0, 6) // capacity = 4 possible adjacent squares, 2 possible tool switches
	for _, iter := range c.edgeIterator {
		coord := xy{source.xyTool.c.x + iter.x, source.xyTool.c.y + iter.y}
		if coord.y < 0 || coord.y >= len(c.erosionLevelMap) || coord.x < 0 || coord.x >= len(c.erosionLevelMap[coord.y]) {
			continue // out of bounds
		}
		destinationTerrain := c.erosionLevelMap[coord.y][coord.x] % 3
		if !c.tools[destinationTerrain][source.xyTool.tool] {
			continue // cannot go to destination with current tool
		}

		edges = append(edges, xyToolWithDepth{xyWithTool{coord, source.xyTool.tool}, source.depth + 1})
	}

	sourceTerrain := c.erosionLevelMap[source.xyTool.c.y][source.xyTool.c.x] % 3
	for t := 0; t < 3; t++ {
		if t != source.xyTool.tool && c.tools[sourceTerrain][t] {
			// add the same xy coord with a different tool and a deeper depth
			edges = append(edges, xyToolWithDepth{xyWithTool{source.xyTool.c, t}, source.depth + 7})
		}
	}

	return edges
}

func (c *cave) print() {
	for y := 0; y <= c.target.y; y++ {
		str := ""
		for x := 0; x <= c.target.x; x++ {
			val := c.erosionLevelMap[y][x] % 3
			if x == 0 && y == 0 {
				val = -1
			}
			if x == c.target.x && y == c.target.y {
				val = -2
			}
			switch val {
			case 0:
				str += "."
			case 1:
				str += "="
			case 2:
				str += "|"
			case -1:
				str += "M"
			case -2:
				str += "T"
			default:
			}
		}
		fmt.Print(str)
		fmt.Println()
	}
}

func printQueue(queue []xyToolWithDepth) {
	for _, q := range queue {
		fmt.Printf("(%d, %d) - %d,%d || ", q.xyTool.c.x, q.xyTool.c.y, q.xyTool.tool, q.depth)
	}
	fmt.Println()
}
