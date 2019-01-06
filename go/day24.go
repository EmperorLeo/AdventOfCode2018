package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	boost = 35
)

type battleGroup struct {
	id, target, team, totalHp, hpPerUnit, ad, initiative int
	adType                                               string
	immunities, weaknesses                               []string
}

type battleGroups []*battleGroup

type immuneSystemSimulator struct {
	round, infectionLeft, immuneSystemLeft int
	groups, immuneSystem, infection        battleGroups
	groupMap                               map[int]*battleGroup
	groupChosenMap                         map[int]bool
}

// Note: a lot of the code that calculates damage uses floats, because of an incorrect assumption
// made that a single unit in a group could end the round damaged (not at full HP.)
func day24() {
	simulator := immuneSystemSimulator{}
	simulator.init()
	for !simulator.winnerFound() {
		simulator.incrementRound()
	}
	finalScore, winningTeam := simulator.getWinnerScore()
	fmt.Printf("Winning army (%d) has %d units left.\n", winningTeam, finalScore)
}

func (iss *immuneSystemSimulator) init() {
	regex := regexp.MustCompile(`^(\d+) units each with (\d+) hit points (?:\((.*)\) )?with an attack that does (\d+) (fire|cold|slashing|bludgeoning|radiation) damage at initiative (\d+)$`)
	lines := readFile("../input/day24.txt")
	groups := []*battleGroup{}
	immuneSystem := []*battleGroup{}
	infection := []*battleGroup{}
	groupMap := map[int]*battleGroup{}
	var team int
	var id int
	for _, line := range lines {
		matches := regex.FindAllStringSubmatch(line, -1)
		if matches == nil {
			if line == "Immune System:" {
				team = 0
			} else if line == "Infection:" {
				team = 1
			}
		} else {
			id++
			bg := battleGroupFromMatch(matches[0], team, id)
			groups = append(groups, bg)
			groupMap[bg.id] = bg
			if team == 0 {
				immuneSystem = append(immuneSystem, bg)
			} else {
				infection = append(infection, bg)
			}
			fmt.Printf("Group %d on team %d has %d units each with %d hp.  Does %d %s damage.\n", bg.id, bg.team, bg.totalHp/bg.hpPerUnit, bg.hpPerUnit, bg.ad, bg.adType)
		}
	}
	fmt.Println()
	iss.groups = groups
	iss.immuneSystem = immuneSystem
	iss.infection = infection
	iss.round = 0
	iss.infectionLeft = len(infection)
	iss.immuneSystemLeft = len(immuneSystem)
	iss.groupMap = groupMap
}

func (iss *immuneSystemSimulator) incrementRound() {
	iss.round++
	iss.targetSelectionPhase()
	iss.attackPhase()
}

func (iss *immuneSystemSimulator) targetSelectionPhase() {
	iss.immuneSystem.sortGroups()
	iss.infection.sortGroups()
	iss.groupChosenMap = map[int]bool{}
	for _, is := range iss.immuneSystem {
		is.chooseTargetFrom(iss.infection, iss.groupChosenMap)
	}
	for _, inf := range iss.infection {
		inf.chooseTargetFrom(iss.immuneSystem, iss.groupChosenMap)
	}
}

func (iss *immuneSystemSimulator) attackPhase() {
	iss.groups.sortGroupsAttack()
	fmt.Printf("Round %d beginning.\n", iss.round)
	for _, group := range iss.groups {
		target := iss.groupMap[group.target]
		if target == nil || target.totalHp <= 0 || group.totalHp <= 0 {
			continue
		}

		damage := group.damageTo(target)
		unitsToKill := damage / target.hpPerUnit
		numTargetUnits := math.Ceil(float64(target.totalHp) / float64(target.hpPerUnit))
		target.totalHp -= unitsToKill * target.hpPerUnit
		if target.totalHp <= 0 { // Critical hit!! Target should never have less than 0 hp due to the rules, but just in case!
			target.totalHp = 0 // ded
			if target.team == 0 {
				iss.immuneSystemLeft--
			} else {
				iss.infectionLeft--
			}
		}
		unitsLeft := math.Ceil(float64(target.totalHp) / float64(target.hpPerUnit))
		unitsKilled := numTargetUnits - unitsLeft
		fmt.Printf("Group %d on team %d did %d %s damage to group %d on team %d, killing %f units.\n", group.id, group.team, damage, group.adType, target.id, target.team, unitsKilled)
	}
	fmt.Println()
}

func (iss *immuneSystemSimulator) winnerFound() bool {
	return iss.immuneSystemLeft == 0 || iss.infectionLeft == 0
}

func (iss *immuneSystemSimulator) getWinnerScore() (int, int) {
	var score, team int
	for _, g := range iss.groups {
		score += int(math.Ceil(float64(g.totalHp) / float64(g.hpPerUnit)))
	}
	if iss.immuneSystemLeft == 0 {
		team = 1
	} else {
		team = 0
	}
	return score, team
}

func (g battleGroups) sortGroups() {
	sort.SliceStable(g, func(i, j int) bool {
		if g[i].effectivePower() > g[j].effectivePower() {
			return true
		} else if g[i].effectivePower() < g[j].effectivePower() {
			return false
		} else {
			return g[i].initiative > g[j].initiative
		}
	})
}

func (g battleGroups) sortGroupsAttack() {
	sort.SliceStable(g, func(i, j int) bool {
		return g[i].initiative > g[j].initiative
	})
}

func (g battleGroup) effectivePower() int {
	return g.ad * int(math.Ceil(float64(g.totalHp)/float64(g.hpPerUnit)))
}

func (g *battleGroup) chooseTargetFrom(bgs battleGroups, groupChosenMap map[int]bool) {
	if g.totalHp <= 0 {
		g.target = 0
		return
	}
	var target *battleGroup
	var damageToTarget, targetEffectivePower, targetInitiative int
	for _, enemy := range bgs {
		if enemy.totalHp <= 0 || groupChosenMap[enemy.id] {
			continue
		}
		damage := g.damageTo(enemy)
		if damage > damageToTarget || (damageToTarget == damage && (enemy.effectivePower() > targetEffectivePower || (enemy.effectivePower() == targetEffectivePower && enemy.initiative > targetInitiative))) {
			target = enemy
			damageToTarget = damage
			targetEffectivePower = enemy.effectivePower()
			targetInitiative = enemy.initiative
		}
	}
	if target != nil && damageToTarget > 0 {
		g.target = target.id
		groupChosenMap[target.id] = true
	} else {
		g.target = 0
	}
}

func (g *battleGroup) damageTo(enemy *battleGroup) int {
	damage := g.effectivePower()
	// fmt.Printf("Group %d with base damage %d is up against an enemy (group %d) with %v weaknesses.\n", g.id, damage, enemy.id, enemy.weaknesses)
	for _, weakness := range enemy.weaknesses {
		if weakness == g.adType {
			damage *= 2
			break
		}
	}
	for _, immunity := range enemy.immunities {
		if immunity == g.adType {
			damage = 0
			break
		}
	}
	return damage
}

func battleGroupFromMatch(match []string, team int, id int) *battleGroup {
	numUnits, _ := strconv.Atoi(match[1])
	hp, _ := strconv.Atoi(match[2])
	buffInfo := match[3]
	ad, _ := strconv.Atoi(match[4])
	if team == 0 {
		ad += boost
	}
	adType := match[5]
	initiative, _ := strconv.Atoi(match[6])

	weaknesses, immunities := []string{}, []string{}

	if len(buffInfo) > 0 {
		buffRegex := regexp.MustCompile(`^(weak|immune) to (.*)$`)
		modifiers := strings.Split(buffInfo, "; ")
		for _, mod := range modifiers {
			buffMatch := buffRegex.FindAllStringSubmatch(mod, -1)[0]
			modType := buffMatch[1]
			modAdTypes := strings.Split(buffMatch[2], ", ")
			for _, modAdType := range modAdTypes {
				if modType == "weak" {
					weaknesses = append(weaknesses, modAdType)
				} else {
					immunities = append(immunities, modAdType)
				}
			}
		}
	}

	return &battleGroup{id, 0, team, numUnits * hp, hp, ad, initiative, adType, immunities, weaknesses}
}
