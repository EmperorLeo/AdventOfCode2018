package main

import (
	"fmt"
)

type marble struct {
	value          int
	previous, next *marble
}

type game struct {
	currentMarble               *marble
	currentPlayer, totalPlayers int
	scores                      map[int]int
}

func day9() {
	input := readFile("../input/day9.txt")[0]
	var players, points int
	fmt.Sscanf(input, "%d players; last marble is worth %d points", &players, &points)

	// Play 2 games
	for _, numPoints := range []int{points, points * 100} {
		game := game{&marble{0, nil, nil}, 1, players, make(map[int]int, players)}
		for i := 1; i <= numPoints; i++ {
			game.takeTurn(i)
		}

		winner, score := game.getWinner()
		fmt.Printf("The winning score is %d, and the winning player is %d.\n", score, winner)
	}
}

func (g *game) takeTurn(round int) {
	if round == 1 {
		// Have to do pointer setup somehow
		g.currentMarble.next = g.currentMarble
		g.currentMarble.previous = g.currentMarble
	}

	if round%23 != 0 {
		leftMarble := g.currentMarble.next
		rightMarble := leftMarble.next
		newMarble := &marble{round, leftMarble, rightMarble}
		leftMarble.next = newMarble
		rightMarble.previous = newMarble

		g.currentMarble = newMarble
	} else {
		marble := g.currentMarble
		for i := 0; i < 7; i++ {
			marble = marble.previous
		}
		g.scores[g.currentPlayer] += (round + marble.value)
		g.currentMarble = marble.next
		marble.previous.next = marble.next
		marble.next.previous = marble.previous
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
