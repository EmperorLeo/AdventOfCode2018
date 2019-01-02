package main

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
