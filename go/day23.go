package main

import (
	"fmt"
	"math"
)

const (
	// Start at 1000, increment back by 10 until you get an answer
	nanobotThreshold = 960
)

type nanobot struct {
	x, y, z, r int
}

type region struct {
	minX, maxX, minY, maxY, minZ, maxZ int
}

func day23() {
	lines := readFile("../input/day23.txt")
	var strongest nanobot
	nanobots := make([]nanobot, len(lines))
	for i, line := range lines {
		bot := nanobot{}
		fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d", &(bot.x), &(bot.y), &(bot.z), &(bot.r))
		nanobots[i] = bot
		if strongest.r < bot.r {
			strongest = bot
		}
	}

	// Part 1
	var reach int
	for _, other := range nanobots {
		if strongest.signalReaches(other) {
			reach++
		}
	}

	fmt.Printf("The strongest nanobot is (%d, %d, %d) with a radius of %d and a reach of %d.\n", strongest.x, strongest.y, strongest.z, strongest.r, reach)

	// Part 2
	// This probably is not right - YEP ITS NOT RIGHT (code left here for shame)
	// pointToCheck := xyz{0, 0, 0}
	// zeroZeroZero := xyz{0, 0, 0}
	// var pointToCheckMatches int
	// for p := 6; p >= 0; p-- {
	// 	pointToCheck = xyz{pointToCheck.x * 10, pointToCheck.y * 10, pointToCheck.z * 10}
	// 	oldPointToCheck := pointToCheck
	// 	pointToCheckMatches = 0
	// 	infinitesimalBots := make([]nanobot, len(nanobots))
	// 	for i, bot := range nanobots {
	// 		bot.x, bot.y, bot.z, bot.r = bot.x/int(math.Pow10(p)), bot.y/int(math.Pow10(p)), bot.z/int(math.Pow10(p)), bot.r/int(math.Pow10(p))
	// 		fmt.Printf("(%d, %d, %d) with r = %d\n", bot.x, bot.y, bot.z, bot.r)
	// 		infinitesimalBots[i] = bot
	// 	}
	// 	for x := 0; x <= 150; x++ {
	// 		for y := 0; y <= 150; y++ {
	// 			for z := 0; z <= 150; z++ {
	// 				coord := xyz{x + oldPointToCheck.x, y + oldPointToCheck.y, oldPointToCheck.z + z}
	// 				leobot := nanobot{x + oldPointToCheck.x, y + oldPointToCheck.y, oldPointToCheck.z + z, 0}
	// 				var count int
	// 				for _, bot := range infinitesimalBots {
	// 					if bot.signalReaches(leobot) {
	// 						count++
	// 					}
	// 				}
	// 				if pointToCheckMatches < count || (pointToCheckMatches == count && manhattanDistance(coord, zeroZeroZero) < manhattanDistance(pointToCheck, zeroZeroZero)) {
	// 					pointToCheck = coord
	// 					pointToCheckMatches = count
	// 				}
	// 			}
	// 		}
	// 	}
	// 	fmt.Printf("The point reached the most is (%d, %d, %d).  It is reached by %d nanobots.\n", pointToCheck.x, pointToCheck.y, pointToCheck.z, pointToCheckMatches)
	// 	fmt.Println("End stage.")
	// }

	// Part 2
	// Ok, here is the real code!!
	// This code may not work if I'm only checking one sub-octree, but if I check all with more than 932 regions I could make some headway
	// I found a point in the bad code up north that was affected by 932
	// There are edge cases where this code would not work, but I think given the high amount of intersecting regions, it will
	var minX, minY, minZ, maxX, maxY, maxZ int
	// Get the bounds of the big region first
	for _, bot := range nanobots {
		if bot.x-bot.r < minX {
			minX = bot.x - bot.r
		}
		if bot.y-bot.r < minY {
			minY = bot.y - bot.r
		}
		if bot.z-bot.r < minZ {
			minZ = bot.z - bot.r
		}
		if bot.x+bot.r > maxX {
			maxX = bot.x + bot.r
		}
		if bot.y+bot.r > maxY {
			maxY = bot.y + bot.r
		}
		if bot.z+bot.r > maxZ {
			maxZ = bot.z + bot.r
		}
	}
	fullRegion := region{minX, maxX, minY, maxY, minZ, maxZ}
	bestPoint, bestCount := fullRegion.getBestPointAndNanobotCountInRegion(nanobots)

	fmt.Printf("The best point that I have discovered is (%d, %d, %d).  It is affected by %d nanobots.\n", bestPoint.x, bestPoint.y, bestPoint.z, bestCount)
	fmt.Printf("%d + %d + %d = %d\n", bestPoint.x, bestPoint.y, bestPoint.z, bestPoint.x+bestPoint.y+bestPoint.z)
}

// I learned how to name methods from the best in college
func (r region) getBestPointAndNanobotCountInRegion(nanobots []nanobot) (xyz, int) {
	count := r.getInRangeOf(nanobots)

	if r.isPoint() {
		return xyz{r.minX, r.minY, r.minZ}, count
	}

	if count < nanobotThreshold {
		// Do not keep recursing if this region isn't affected by more than the threshold that I have already determined exists
		return xyz{}, -1
	}

	regions := r.splice()

	var bestPoint xyz
	bestCount := -1

	for _, subRegion := range regions {
		p, c := subRegion.getBestPointAndNanobotCountInRegion(nanobots)
		if c > bestCount || (c == bestCount && c != -1 && manhattanDistance(p, xyz{0, 0, 0}) < manhattanDistance(bestPoint, xyz{0, 0, 0})) {
			bestPoint = p
			bestCount = c
		}
	}

	return bestPoint, bestCount
}

func (nb nanobot) signalReaches(other nanobot) bool {
	// fmt.Printf("Distance is %d, and radius is %d.\n", distance, nb.r)
	return manhattanDistance(xyz{nb.x, nb.y, nb.z}, xyz{other.x, other.y, other.z}) <= nb.r
}

func manhattanDistance(p1, p2 xyz) int {
	return int(math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y)) + math.Abs(float64(p1.z-p2.z)))
}

func (r region) splice() []region {
	regions := make([]region, 8)
	xSplit := (r.maxX - r.minX) / 2
	ySplit := (r.maxY - r.minY) / 2
	zSplit := (r.maxZ - r.minZ) / 2
	// Note: using this formula, sometimes the regions will overlap at their edges, but for the purpose of this problem, it is fine.
	// This is because I'm doing a recursive-style search, not ever comparing these regions to each other.
	regions[0] = region{r.minX, r.minX + xSplit, r.minY, r.minY + ySplit, r.minZ, r.minZ + zSplit}
	regions[1] = region{r.maxX - zSplit, r.maxX, r.minY, r.minY + ySplit, r.minZ, r.minZ + zSplit}
	regions[2] = region{r.minX, r.minX + xSplit, r.maxY - ySplit, r.maxY, r.minZ, r.minZ + zSplit}
	regions[3] = region{r.maxX - zSplit, r.maxX, r.maxY - ySplit, r.maxY, r.minZ, r.minZ + zSplit}
	regions[4] = region{r.minX, r.minX + xSplit, r.minY, r.minY + ySplit, r.maxZ - zSplit, r.maxZ}
	regions[5] = region{r.maxX - zSplit, r.maxX, r.minY, r.minY + ySplit, r.maxZ - zSplit, r.maxZ}
	regions[6] = region{r.minX, r.minX + xSplit, r.maxY - ySplit, r.maxY, r.maxZ - zSplit, r.maxZ}
	regions[7] = region{r.maxX - zSplit, r.maxX, r.maxY - ySplit, r.maxY, r.maxZ - zSplit, r.maxZ}
	return regions
}

func (r region) getInRangeOf(nanobots []nanobot) int {
	var count int
	for _, bot := range nanobots {
		p := xyz{bot.x, bot.y, bot.z}
		if r.containsPoint(p) {
			// bot is inside of the region, so definitely affects it
			count++
			continue
		}
		// Walk each point to the boundary of the cube, and see if the radius can hit that point
		remainingInfluence := bot.r
		if p.x < r.minX {
			remainingInfluence -= (r.minX - p.x)
		} else if p.x > r.maxX {
			remainingInfluence -= (p.x - r.maxX)
		}
		if p.y < r.minY {
			remainingInfluence -= (r.minY - p.y)
		} else if p.y > r.maxY {
			remainingInfluence -= (p.y - r.maxY)
		}
		if p.z < r.minZ {
			remainingInfluence -= (r.minZ - p.z)
		} else if p.z > r.maxZ {
			remainingInfluence -= (p.z - r.maxZ)
		}
		if remainingInfluence >= 0 {
			count++
		}
	}
	return count
}

func (r region) isPoint() bool {
	return r.minX == r.maxX && r.minY == r.minY && r.minZ == r.minZ
}

func (r region) containsPoint(p xyz) bool {
	return p.x >= r.minX && p.x <= r.maxX && p.y >= r.minY && p.y <= r.maxY && p.z >= r.minZ && p.z <= r.maxZ
}
