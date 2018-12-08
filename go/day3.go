package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// Claim needs a comment cause go linter said so
type Claim struct {
	id     int
	x      int
	y      int
	width  int
	height int
}

func day3() {
	lines := readFile("../input/day3.txt")
	claims := make([]Claim, len(lines))
	regex := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)
	var maxWidth, maxHeight int
	for i, line := range lines {
		matches := regex.FindAllStringSubmatch(line, -1)
		c := claimFromMatch(matches[0])
		claims[i] = c
		if c.height+c.y > maxHeight {
			maxHeight = c.height + c.y
		}
		if c.width+c.x > maxWidth {
			maxWidth = c.width + c.x
		}
	}
	fabricClaims := make([][]int, maxHeight)
	for i := 0; i < maxHeight; i++ {
		fabricClaims[i] = make([]int, maxWidth)
	}

	var numConflicted int

	// Part 1
	for _, claim := range claims {
		for x := claim.x; x < claim.x+claim.width; x++ {
			for y := claim.y; y < claim.y+claim.height; y++ {
				if fabricClaims[y][x] == 0 {
					fabricClaims[y][x] = claim.id
				} else if fabricClaims[y][x] > 0 {
					numConflicted++
					fabricClaims[y][x] = -1
				}
			}
		}
	}

	fmt.Println(fmt.Sprintf("Number of conflicted sqr inches: %d", numConflicted))

	// Part 2
	for _, claim := range claims {
		var conflicted bool
		for x := claim.x; x < claim.x+claim.width; x++ {
			for y := claim.y; y < claim.y+claim.height; y++ {
				if fabricClaims[y][x] != claim.id {
					conflicted = true
				}
			}
		}

		if !conflicted {
			fmt.Println(fmt.Sprintf("ID of non-conflicting claim: %d", claim.id))
			break
		}
	}
}

func claimFromMatch(match []string) Claim {
	id, _ := strconv.Atoi(match[1])
	x, _ := strconv.Atoi(match[2])
	y, _ := strconv.Atoi(match[3])
	width, _ := strconv.Atoi(match[4])
	height, _ := strconv.Atoi(match[5])

	return Claim{id, x, y, width, height}
}
