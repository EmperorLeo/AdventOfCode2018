package main

import (
	"fmt"
	"sort"
)

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

func day16() {
	game := cavern{}
	game.init()
	for !game.winnerFound() {
		game.incrementTurn()
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
	for _, unit := range c.units {
		// TODO
	}
}

func (c *cavern) winnerFound() bool {
	return len(c.elves) == 0 || len(c.goblins) == 0
}

func (c *cavern) getWinner() (rune, int) {
	var winner rune
	if len(c.elves) == 0 {
		winner = 'E'
	} else {
		winner = 'G'
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
