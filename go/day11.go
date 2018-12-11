package main

import (
	"fmt"
	"math"
)

const (
	// SerialNumber - This apparently needs a comment
	SerialNumber = 6042
)

type cell struct {
	powerLevel int
	rackID     int
}

type powerGrid struct {
	cells  [][]cell
	serial int
}

func day11() {
	cells := make([][]cell, 300)
	for i := 0; i < 300; i++ {
		cells[i] = make([]cell, 300)
	}

	grid := powerGrid{cells, SerialNumber}
	grid.initGrid()
	s, x, y := grid.findLargestTotalPowerSquare()
	fmt.Printf("Region with largest power is (%d, %d, %d).", x, y, s)
	fmt.Println()
}

func (g *powerGrid) initGrid() {
	for y := 0; y < len(g.cells); y++ {
		row := g.cells[y]
		for x := 0; x < len(row); x++ {
			row[x].rackID = (x + 1) + 10
			intermediate := row[x].rackID * ((row[x].rackID * (y + 1)) + g.serial)
			row[x].powerLevel = (cutOffThousandsPlacePlus(intermediate) / 100) - 5
		}
	}
}

func (g *powerGrid) findLargestTotalPowerSquare() (int, int, int) {
	var sizeOfLargest, largestX, largestY, largestPower int
	largestPower = math.MinInt32
	for s := 1; s <= len(g.cells); s++ {
		for y := 0; y < len(g.cells)-(s-1); y++ {
			for x := 0; x < len(g.cells)-(s-1); x++ {
				var powerLevelArea int
				for a := 0; a < s; a++ {
					for b := 0; b < s; b++ {
						powerLevelArea += g.cells[y+a][x+b].powerLevel
					}
				}
				if powerLevelArea > largestPower {
					largestX = x + 1
					largestY = y + 1
					largestPower = powerLevelArea
					sizeOfLargest = s
				}
			}
		}
	}
	return sizeOfLargest, largestX, largestY
}

func cutOffThousandsPlacePlus(number int) int {
	return number - ((number / 1000) * 1000)
}
