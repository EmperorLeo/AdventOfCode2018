package main

import (
	"errors"
	"fmt"
)

// -1 is left, 0 is straight, 1 is right (for next direction)
// 0 is up, 1 is right, 2 is down, 3 is left for direction
type cart struct {
	coords                   coordinates
	direction, nextDirection int
}

type track struct {
	char rune
}

type coordinates struct {
	x, y int
}

type trackSystem struct {
	tracks [][]track
	carts  []cart
}

func day13() {
	// Part 1
	system := trackSystem{}
	system.init()
	x, y := -1, -1
	for x < 0 && y < 0 {
		// BUG: My input did not have this problem, but technically the carts only move one at a time
		system.incrementFrame()
		coords, err := system.lookForCollision()
		if err == nil {
			x = coords.x
			y = coords.y
		}
	}
	fmt.Printf("Coordinates of first crash are (%d, %d).", x, y)
	fmt.Println()

	// Part 2
	smartSystem := trackSystem{}
	smartSystem.init()
	x, y = -1, -1
	for x < 0 && y < 0 {
		smartSystem.incrementFrame()
		smartSystem.removeAtCollisions()
		if len(smartSystem.carts) == 1 {
			x = smartSystem.carts[0].coords.x
			y = smartSystem.carts[0].coords.y
		}
	}
	fmt.Printf("Coordinates of last remaining cart are (%d, %d).", x, y)
	fmt.Println()
}

func (t *trackSystem) init() {
	lines := readFile("../input/day13.txt")
	t.tracks = make([][]track, len(lines))
	t.carts = []cart{}
	for y, line := range lines {
		t.tracks[y] = make([]track, len(line))
		for x, char := range line {
			trackType := char
			coords := coordinates{x, y}
			switch trackType {
			case '^':
				trackType = '|'
				t.carts = append(t.carts, cart{coords, 0, -1})
			case 'v':
				trackType = '|'
				t.carts = append(t.carts, cart{coords, 2, -1})
			case '<':
				trackType = '-'
				t.carts = append(t.carts, cart{coords, 3, -1})
			case '>':
				trackType = '-'
				t.carts = append(t.carts, cart{coords, 1, -1})
			default:
				// nothing
			}
			t.tracks[y][x] = track{trackType}
		}
	}
}

func (t *trackSystem) incrementFrame() {
	for i := 0; i < len(t.carts); i++ {
		t.carts[i].moveCart(t.tracks)
	}
}

func (t *trackSystem) lookForCollision() (coordinates, error) {
	coordsMap := make(map[coordinates]bool, len(t.carts))
	var coords coordinates
	for _, cart := range t.carts {
		if coordsMap[cart.coords] {
			return cart.coords, nil
		}
		coordsMap[cart.coords] = true
	}
	return coords, errors.New("No coordinates found")
}

func (t *trackSystem) removeAtCollisions() {
	coordsMap := make(map[coordinates]int, len(t.carts))
	for _, cart := range t.carts {
		coordsMap[cart.coords]++
	}
	remainingCarts := []cart{}
	for _, cart := range t.carts {
		if coordsMap[cart.coords] < 2 {
			remainingCarts = append(remainingCarts, cart)
		}
	}
	t.carts = remainingCarts
}

func (c *cart) moveCart(tracks [][]track) {
	if c.direction == 0 {
		c.moveUp(tracks)
	} else if c.direction == 1 {
		c.moveRight(tracks)
	} else if c.direction == 2 {
		c.moveDown(tracks)
	} else {
		c.moveLeft(tracks)
	}

	track := tracks[c.coords.y][c.coords.x]
	switch track.char {
	case '/':
		c.processSlash()
	case '\\':
		c.processBackslash()
	case '+':
		c.processCrossroads()
	}
}

func (c *cart) moveUp(tracks [][]track) {
	c.coords.y--
}

func (c *cart) moveRight(tracks [][]track) {
	c.coords.x++
}

func (c *cart) moveDown(tracks [][]track) {
	c.coords.y++
}

func (c *cart) moveLeft(tracks [][]track) {
	c.coords.x--
}

func (c *cart) processCrossroads() {
	c.direction += c.nextDirection
	c.nextDirection++

	if c.nextDirection > 1 {
		c.nextDirection = -1
	}
	if c.direction < 0 {
		c.direction += 4
	} else if c.direction > 3 {
		c.direction -= 4
	}
}

func (c *cart) processSlash() {
	switch c.direction {
	case 0:
		c.direction = 1
	case 1:
		c.direction = 0
	case 2:
		c.direction = 3
	case 3:
		c.direction = 2
	}
}

func (c *cart) processBackslash() {
	switch c.direction {
	case 0:
		c.direction = 3
	case 1:
		c.direction = 2
	case 2:
		c.direction = 1
	case 3:
		c.direction = 0
	}
}
