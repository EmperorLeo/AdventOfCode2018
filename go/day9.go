package main

import (
	"fmt"
)

type game struct {
	currentMarble int
	currentPlayer int
	totalPlayers  int
	circle        []int
	scores        map[int]int
}

func day9() {
	input := readFile("../input/day9.txt")[0]
	var players, points int
	fmt.Sscanf(input, "%d players; last marble is worth %d points", &players, &points)
	// Part 1
	fastGame := game{0, 1, players, []int{0}, make(map[int]int, players)}
	for i := 1; i <= points; i++ {
		fastGame.takeTurn(i)
	}

	winner, score := fastGame.getWinner()
	fmt.Printf("The winning score is %d, and the winning player is %d.\n", score, winner)

	// Part 2
	slowGame := game{0, 1, players, []int{0}, make(map[int]int, players)}
	for i := 1; i <= points*100; i++ {
		slowGame.takeTurn(i)
	}

	winner, score = slowGame.getWinner()
	fmt.Printf("The winning score is %d, and the winning player is %d.\n", score, winner)
}

func (g *game) takeTurn(round int) {
	if round%23 != 0 {
		// Two marble should go 2 positions clockwise accounting for new length of circle
		pos := g.currentMarble + 2
		// Go back to the beginning of the "circle" if pos is too big
		if pos >= len(g.circle) {
			pos = pos % (len(g.circle))
		}
		largerCircle := make([]int, len(g.circle)+1)
		copy(largerCircle[:pos], g.circle[:pos])
		largerCircle[pos] = round
		copy(largerCircle[pos+1:], g.circle[pos:])
		g.circle = largerCircle
		g.currentMarble = pos
	} else {
		pos := g.currentMarble - 7
		for pos < 0 {
			pos += len(g.circle)
		}
		g.scores[g.currentPlayer] += (round + g.circle[pos])
		g.circle = append(g.circle[:pos], g.circle[pos+1:]...)
		if pos >= len(g.circle) {
			pos = 0
		}
		g.currentMarble = pos
	}

	// Finish the turn by moving to the next player
	if g.currentPlayer == g.totalPlayers {
		g.currentPlayer = 1
	} else {
		g.currentPlayer++
	}
}

func (g *game) getWinner() (int, int) {
	var player, max int
	for p, i := range g.scores {
		if i > max {
			max = i
			player = p
		}
	}

	return player, max
}
