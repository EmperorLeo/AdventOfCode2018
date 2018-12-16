package main

import (
	"fmt"
	"math"
	"sort"
)

type coordinate struct {
	x, y int
}

type unit struct {
	x, y, hp, ap int
	badGuy       bool
}

type cavern struct {
	units   []*unit
	goblins []*unit
	elves   []*unit
	caveMap [][]bool
	turn    int
}

type coordinateDepth struct {
	c     coordinate
	depth int
}

func day16() {
	game := cavern{}
	game.init()
	for !game.winnerFound() {
		game.sortUnits()
		game.incrementTurn()
		game.cleanUpTheDead()
	}
	winner, totalHp := game.getWinner()
	fmt.Printf("Team %q won with outcome (totalHp * rounds) = %d * %d = %d", winner, totalHp, game.turn, totalHp*game.turn)
}

func (c *cavern) init() {
	lines := readFile("../input/day15.txt")
	caveMap := make([][]bool, len(lines))
	units := []*unit{}
	goblins := []*unit{}
	elves := []*unit{}
	for y, line := range lines {
		caveMap[y] = make([]bool, len(line))
		for x, char := range line {
			switch char {
			case '#':
				caveMap[y][x] = false
			case '.':
				caveMap[y][x] = true
			case 'E':
				caveMap[y][x] = false
				elf := unit{x, y, 200, 3, false}
				elves = append(elves, &elf)
				units = append(units, &elf)
			case 'G':
				caveMap[y][x] = false
				goblin := unit{x, y, 200, 3, true}
				goblins = append(goblins, &goblin)
				units = append(units, &goblin)
			}
		}
	}
}

func (c *cavern) incrementTurn() {
	c.turn++
	for _, u := range c.units {
		if u.hp <= 0 {
			// Need to deal with this if the unit has not been cleaned up yet
			continue
		}

		var enemies []*unit
		if u.badGuy {
			enemies = c.elves
		} else {
			enemies = c.goblins
		}
		u.takeTurn(c.caveMap, enemies)
	}
}

func (c *cavern) cleanUpTheDead() {
	dead := []*unit{}
	for _, u := range c.units {
		if u.hp <= 0 {
			dead = append(dead, u)
		}
	}

	for _, deadUnit := range dead {
		var unitIndex int
		for i, u := range c.units {
			if u == deadUnit {
				unitIndex = i
				break
			}
		}
		c.units = remove(c.units, unitIndex)
		if deadUnit.badGuy {
			for i, u := range c.goblins {
				if u == deadUnit {
					unitIndex = i
					break
				}
			}
			c.goblins = remove(c.goblins, unitIndex)
		} else {
			for i, u := range c.elves {
				if u == deadUnit {
					unitIndex = i
					break
				}
			}
			c.elves = remove(c.elves, unitIndex)
		}
	}
}

func (c *cavern) winnerFound() bool {
	return len(c.elves) == 0 || len(c.goblins) == 0
}

func (c *cavern) getWinner() (rune, int) {
	var winner rune
	if len(c.elves) == 0 {
		winner = 'G'
	} else {
		winner = 'E'
	}

	var totalHp int
	for _, unit := range c.units {
		totalHp += unit.hp
	}

	return winner, totalHp
}

func (c *cavern) sortUnits() {
	sort.SliceStable(c.units, func(i, j int) bool {
		if c.units[i].y < c.units[j].y {
			return true
		} else if c.units[i].y > c.units[j].y {
			return false
		} else {
			return c.units[i].x < c.units[j].x
		}
	})
}

func (u *unit) takeTurn(graph [][]bool, enemies []*unit) {
	var adjacentToTarget *unit
	var destinationCoord coordinate
	var currentNextCoord coordinate
	currentPathLength := math.MaxInt32
	for _, enemy := range enemies {
		if u.isAdjacentTo(enemy) {
			if adjacentToTarget == nil || adjacentToTarget != nil && compareAdjacent(enemy, adjacentToTarget) {
				adjacentToTarget = enemy
			}
			currentPathLength = 0
		} else if currentPathLength > 0 {
			for i := -1; i <= 1; i += 2 {
				for j := -1; j <= 1; j += 2 {
					coord := coordinate{u.x + i, u.y + j}
					nextCoord, pathLength := u.getShortestPath(coord, graph)
					if pathLength != -1 {
						if currentPathLength > pathLength || currentPathLength == pathLength && compareReadingDistance(coord, destinationCoord) {
							destinationCoord = coord
							currentPathLength = pathLength
							currentNextCoord = nextCoord
						}
					}
				}
			}
		}

		// If an adjacent enemy or a path was found
		if adjacentToTarget != nil {
			u.attack(adjacentToTarget)
		} else if currentPathLength != math.MaxInt32 {
			// Open up the area once the unit leaves
			graph[u.y][u.x] = true
			u.x = currentNextCoord.x
			u.y = currentNextCoord.y
			// Close up the area once the unit enters
			graph[u.y][u.x] = false
		}
	}
}

func (u *unit) isAdjacentTo(enemy *unit) bool {
	return u.x == enemy.x && int(math.Abs(float64(u.y-enemy.y))) == 1 || u.y == enemy.y && int(math.Abs(float64(u.x-enemy.x))) == 1
}

func (u *unit) getShortestPath(c coordinate, graph [][]bool) (coordinate, int) {
	// Implement BFS algorithm to find shortest path (modified Dijkstra's algorithm based on the distance always being 1)
	initialCoordinate := coordinate{u.x, u.y}
	visitedMap := map[coordinate]int{initialCoordinate: -1}
	queue := []coordinateDepth{}
	for len(queue) > 0 {
		queueItem := queue[0]
		queue = queue[1:]
		newDepth := queueItem.depth + 1
		if queueItem.depth == -1 {
			// Take care of first item, which has a depth of -1
			newDepth = 1
		}

		for i := -1; i <= 1; i += 2 {
			for j := -1; j <= 1; j += 2 {
				coord := coordinate{queueItem.c.x + i, queueItem.c.y + j}
				if coord.canTraverse(graph) && visitedMap[coord] == 0 {
					visitedMap[coord] = newDepth
					queue = append(queue, coordinateDepth{coord, newDepth})
				}
				if coord == c {
					return coord, newDepth
				}
			}
		}
	}

	return coordinate{}, -1
}

func (u *unit) attack(enemy *unit) {
	enemy.hp -= u.ap
}

func (c *coordinate) canTraverse(graph [][]bool) bool {
	if c.y < 0 || c.y >= len(graph) {
		return false
	}

	if c.x < 0 || c.x >= len(graph[c.y]) {
		return false
	}

	return graph[c.y][c.x]
}

func compareAdjacent(a *unit, b *unit) bool {
	if a.hp < b.hp {
		return true
	} else if a.hp > b.hp {
		return false
	} else {
		if a.y < b.y {
			return true
		} else if a.y > b.y {
			return false
		} else {
			return a.x < b.x
		}
	}
}

func compareReadingDistance(a coordinate, b coordinate) bool {
	if a.y < b.y {
		return true
	} else if a.y > b.y {
		return false
	} else {
		return a.x < b.x
	}
}

func remove(units []*unit, index int) []*unit {
	units[index] = units[len(units)-1]
	return units[:len(units)-1]
}
