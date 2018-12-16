package main

import (
	"fmt"
	"math"
	"sort"
)

var coordIncrementor = []coordinate{coordinate{0, -1}, coordinate{-1, 0}, coordinate{1, 0}, coordinate{0, 1}}

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

func day15() {
	// Part 1
	game := cavern{}
	elfAttackPower := 3
	game.init(elfAttackPower)
	for !game.winnerFound() {
		game.sortUnits()
		game.incrementTurn()
		game.cleanUpTheDead()
		game.print()
	}
	winner, totalHp := game.getWinner()
	fmt.Printf("Team %q won with outcome (totalHp * rounds) = %d * %d = %d", winner, totalHp, game.turn, totalHp*(game.turn))
	fmt.Println()

	// Part 2
	elvesWin := false
	for !elvesWin {
		elfAttackPower++
		game = cavern{}
		game.init(elfAttackPower)
		initialElfCount := len(game.elves)
		for !game.winnerFound() {
			game.sortUnits()
			game.incrementTurn()
			game.cleanUpTheDead()
		}
		winner, totalHp = game.getWinner()
		fmt.Printf("Team %q won with outcome (totalHp * rounds) = %d * %d = %d. Elves had %d AP.", winner, totalHp, game.turn, totalHp*(game.turn), elfAttackPower)
		fmt.Println()

		if winner == 'E' && len(game.elves) == initialElfCount {
			elvesWin = true
		}
	}
}

func (c *cavern) init(elfAttackPower int) {
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
				elf := unit{x, y, 200, elfAttackPower, false}
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
	c.units = units
	c.elves = elves
	c.goblins = goblins
	c.caveMap = caveMap
}

func (c *cavern) incrementTurn() {
	c.turn++
	for _, u := range c.units {
		if u.hp <= 0 {
			// Need to deal with this if the unit has not been cleaned up yet
			// ALSO - EDGE CASE:
			// ####
			// ##E#
			// #GG#
			// ####
			// THE LAST UNIT IS DEAD, so should not check c.winnerFound() within the round if the last unit is dead.
			continue
		}

		// If a full round was not completed, decrement turn for return value.
		if c.winnerFound() {
			c.turn--
			break
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
		// Commented out code introduced a nasty bug - a unit could have moved into that spot before the dead were cleaned 😬
		// c.caveMap[deadUnit.y][deadUnit.x] = true
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
	return len(c.elves) == 0 || len(c.goblins) == 0 || allUnitsDead(c.elves) || allUnitsDead(c.goblins)
}

func (c *cavern) getWinner() (rune, int) {
	var winner rune
	if len(c.elves) == 0 || allUnitsDead(c.elves) {
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

func (c *cavern) print() {
	fmt.Printf("ROUND %d RESULTS:", c.turn)
	fmt.Println()
	for y, row := range c.caveMap {
		stringToPrint := ""
		combatInfo := ""
		for x, val := range row {
			var curUnit *unit
			for _, u := range c.units {
				if u.x == x && u.y == y {
					curUnit = u
					break
				}
			}

			if curUnit != nil {
				if curUnit.badGuy {
					stringToPrint += "G"
					combatInfo += fmt.Sprintf("G(%d) ", curUnit.hp)
				} else {
					stringToPrint += "E"
					combatInfo += fmt.Sprintf("E(%d) ", curUnit.hp)
				}
			} else if val {
				stringToPrint += "."
			} else {
				stringToPrint += "#"
			}
		}

		stringToPrint += " " + combatInfo

		fmt.Print(stringToPrint)
		fmt.Println()
	}
	fmt.Println()
}

func (u *unit) takeTurn(graph [][]bool, enemies []*unit) {
	var adjacentToTarget *unit
	var destinationCoord coordinate
	var currentNextCoord coordinate
	currentPathLength := math.MaxInt32
	for _, enemy := range enemies {
		if enemy.hp <= 0 {
			// Enemy not cleaned up yet
			continue
		}

		if u.isAdjacentTo(enemy) {
			if adjacentToTarget == nil || adjacentToTarget != nil && compareAdjacent(enemy, adjacentToTarget) {
				adjacentToTarget = enemy
			}
			currentPathLength = 0
		} else if currentPathLength > 0 {
			for _, incrementor := range coordIncrementor {
				coord := coordinate{enemy.x + incrementor.x, enemy.y + incrementor.y}
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
		u.attack(adjacentToTarget, graph)
	} else if currentPathLength != math.MaxInt32 {
		// Open up the area once the unit leaves
		graph[u.y][u.x] = true
		u.x = currentNextCoord.x
		u.y = currentNextCoord.y
		// Close up the area once the unit enters
		graph[u.y][u.x] = false

		// Handle case where unit is now adjacent to enemy
		// Loop over enemies twice 😰
		var enemyToAttack *unit
		for _, enemy := range enemies {
			if u.isAdjacentTo(enemy) && (enemyToAttack == nil || compareAdjacent(enemy, enemyToAttack)) {
				enemyToAttack = enemy
			}
		}
		if enemyToAttack != nil {
			u.attack(enemyToAttack, graph)
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
	backwardsMap := map[coordinate]*coordinate{}
	queue := []coordinateDepth{}
	top := coordinate{u.x, u.y - 1}
	left := coordinate{u.x - 1, u.y}
	right := coordinate{u.x + 1, u.y}
	bottom := coordinate{u.x, u.y + 1}
	// Kinda hacky so the backwards map ends up at these nodes
	if top.canTraverse(graph) {
		queue = append(queue, coordinateDepth{top, 1})
		visitedMap[top] = 1
		if top == c {
			return top, 1
		}
	}
	if left.canTraverse(graph) {
		queue = append(queue, coordinateDepth{left, 1})
		visitedMap[left] = 1
		if left == c {
			return left, 1
		}
	}
	if right.canTraverse(graph) {
		queue = append(queue, coordinateDepth{right, 1})
		visitedMap[right] = 1
		if right == c {
			return right, 1
		}
	}
	if bottom.canTraverse(graph) {
		queue = append(queue, coordinateDepth{bottom, 1})
		visitedMap[bottom] = 1
		if bottom == c {
			return bottom, 1
		}
	}

	for len(queue) > 0 {
		queueItem := queue[0]
		// fmt.Printf("Testing coordinate")
		queue = queue[1:]
		newDepth := queueItem.depth + 1
		for _, incrementor := range coordIncrementor {
			coord := coordinate{queueItem.c.x + incrementor.x, queueItem.c.y + incrementor.y}
			// fmt.Print(coord.canTraverse(graph))
			// fmt.Println()
			if coord.canTraverse(graph) && visitedMap[coord] == 0 {
				visitedMap[coord] = newDepth
				queue = append(queue, coordinateDepth{coord, newDepth})
				backwardsMap[coord] = &queueItem.c

				if coord == c {
					// fmt.Printf("Unit wants to go to %d, %d", coord.x, coord.y)
					// fmt.Println()
					for backwardsMap[coord] != nil {
						// Traverse backwards until you get to the first step
						coord = *backwardsMap[coord]
					}
					return coord, newDepth
				}
			}
		}
	}

	return coordinate{}, -1
}

func (u *unit) attack(enemy *unit, graph [][]bool) {
	enemy.hp -= u.ap
	if enemy.hp <= 0 {
		graph[enemy.y][enemy.x] = true
	}
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

func allUnitsDead(units []*unit) bool {
	unitsAreAllDead := true
	for _, u := range units {
		if u.hp > 0 {
			unitsAreAllDead = false
			break
		}
	}
	return unitsAreAllDead
}
