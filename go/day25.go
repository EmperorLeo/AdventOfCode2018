package main

import (
	"fmt"
	"math"
)

type constellation struct {
	id      int
	stars   []xyzr
	deleted bool
}

func day25() {
	lines := readFile("../input/day25.txt")
	stars := make([]xyzr, len(lines))
	unclaimedStarSet := make(map[xyzr]bool, len(lines))
	starConstellationMap := make(map[xyzr]*constellation, len(lines))
	constellations := []*constellation{}
	for i, line := range lines {
		star := xyzr{}
		fmt.Sscanf(line, "%d,%d,%d,%d", &star.x, &star.y, &star.z, &star.r)
		stars[i] = star
		unclaimedStarSet[star] = true
	}

	var constIdIncrementor int
	for i, s1 := range stars {
		for _, s2 := range stars[i:] {
			if s1 == s2 || !s1.inRangeOf(s2) {
				continue
			}
			const1 := starConstellationMap[s1]
			const2 := starConstellationMap[s2]
			if const1 != nil && const1 == const2 {
				continue
			}
			// fmt.Printf("Comparing s1 (%d,%d,%d,%d) to s2 (%d,%d,%d,%d)\n", s1.x, s1.y, s1.z, s1.r, s2.x, s2.y, s2.z, s2.r)
			if const1 == nil && const2 != nil {
				const2.add(s1)
				starConstellationMap[s1] = const2
			} else if const1 != nil && const2 == nil {
				const1.add(s2)
				starConstellationMap[s2] = const1
			} else if const1 != nil && const2 != nil {
				newConst := const1.mergeWith(const2)
				constIdIncrementor++
				newConst.id = constIdIncrementor
				for _, starInNewConst := range newConst.stars {
					starConstellationMap[starInNewConst] = &newConst
				}
				const1.deleted = true
				const2.deleted = true
				constellations = append(constellations, &newConst)
			} else {
				constIdIncrementor++
				newConst := constellation{constIdIncrementor, []xyzr{s1, s2}, false}
				starConstellationMap[s1] = &newConst
				starConstellationMap[s2] = &newConst
				constellations = append(constellations, &newConst)
			}
		}
	}
	var totalConstellations int
	for _, c := range constellations {
		if !c.deleted {
			c.print()
			fmt.Println()
			totalConstellations++
		}
	}
	for _, star := range stars {
		if starConstellationMap[star] == nil {
			totalConstellations++
		}
	}
	fmt.Printf("There are %d constellations.\n", totalConstellations)
}

func (star xyzr) inRangeOf(other xyzr) bool {
	return int(
		math.Abs(float64(star.x-other.x))+
			math.Abs(float64(star.y-other.y))+
			math.Abs(float64(star.z-other.z))+
			math.Abs(float64(star.r-other.r))) <= 3
}

func (c *constellation) add(star xyzr) {
	c.stars = append(c.stars, star)
}

func (c *constellation) mergeWith(other *constellation) constellation {
	merged := make([]xyzr, len(c.stars)+len(other.stars))
	for i, star := range c.stars {
		merged[i] = star
	}
	for i, star := range other.stars {
		merged[i+len(c.stars)] = star
	}
	return constellation{0, merged, false}
}

func (c *constellation) print() {
	fmt.Printf("Constellation %d:\n", c.id)
	for _, star := range c.stars {
		fmt.Printf("%d,%d,%d,%d\n", star.x, star.y, star.z, star.r)
	}
}
