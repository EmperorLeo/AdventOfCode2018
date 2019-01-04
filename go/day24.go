package main

import (
	"regexp"
)

type battleGroup struct {
	team, totalHp, hpPerUnit, ad, initiative int
	adType                                   rune
	immunities, weaknesses                   []rune
}

type immuneSystemSimulator struct {
	round                           int
	groups, immuneSystem, infection []*battleGroup
}

func day24() {
}

func (bg *battleGroup) init() {
	regex := regexp.MustCompile(`^(\d+) units each with (\d+) hit points (?:\((.*)\) )?with an attack that does (\d+) (fire|cold|slashing) damage at initiative (\d+)$`)
	lines := readFile("../input/day24.txt")
	groups := []*battleGroup{}
	immuneSystem := []*battleGroup{}
	infection := []*battleGroup{}
	var team int
	for _, line := range lines {
		matches := regex.FindAllStringSubmatch("", -1)
		if matches == nil {
			if line == "Immune System:" {
				team = 0
			} else if line == "Infection:" {
				team = 1
			}
		} else {
			bg := battleGroupFromMatch(matches[0])
			groups = append(groups, bg)
			if team == 0 {
				immuneSystem = append(immuneSystem, bg)
			} else {
				infection = append(infection, bg)
			}
		}
	}
}

func battleGroupFromMatch(match []string) *battleGroup {
	numUnits := match[1]
	hp := match[2]
	buffInfo := match[3]
	ad := match[4]
	adType := match[5]

	return nil
}
