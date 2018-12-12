package main

import (
	"fmt"
)

const (
	initialState = "##...#...###.#.#..#...##.###..###....#.#.###.#..#....#..#......##..###.##..#.##..##..#..#.##.####.##"
	fiftyBillion = 50000000000
)

type pot struct {
	id        int
	plant     bool
	nextPlant bool
	left      *pot
	right     *pot
}

type machine struct {
	referencePot *pot
	stateMap     map[string]bool
	plantMap     map[bool]rune
}

func day12() {
	// Part 1
	stateMachine := machine{}
	stateMachine.init()
	for i := 1; i <= 120; i++ {
		fmt.Print(stateMachine.print())
		fmt.Println()
		stateMachine.proceedNextGen()
	}
	fmt.Printf("The total sum of all the pot IDs that contain a plant is %d.", stateMachine.getTotal())
	fmt.Println()

	// Part 2
	crazyMachine := machine{}
	crazyMachine.init()
	// The pattern should be established after 120 runs
	iterations := 120
	for i := 1; i <= iterations; i++ {
		crazyMachine.proceedNextGen()
	}
	iterationsLeft := fiftyBillion - iterations
	totalPlants := crazyMachine.getTotalPlants()
	crazyMachineTotal := crazyMachine.getTotal() + (iterationsLeft * totalPlants) // Each iteration adds totalPlants number of points, because each plant shifts one to the left

	fmt.Printf("The total sum of all the pot IDs in the crazy machine that contain a plant is %d.", crazyMachineTotal)
	fmt.Println()
}

func (s *machine) init() {
	lines := readFile("../input/day12.txt")
	s.stateMap = make(map[string]bool)
	for _, line := range lines {
		var init string
		var next string
		fmt.Sscanf(line, "%s => %s", &init, &next)
		fmt.Printf("%s maps to %s", init, next)
		fmt.Println()
		s.stateMap[init] = next == "#"
	}
	s.plantMap = make(map[bool]rune, 2)
	s.plantMap[true] = '#'
	s.plantMap[false] = '.'
	s.referencePot = &pot{len(initialState) - 1, initialState[len(initialState)-1] == '#', false, nil, nil}
	rightPot := s.referencePot
	var newPot *pot
	for i := len(initialState) - 2; i >= 0; i-- {
		// Initialize new pot
		newPot = &pot{i, initialState[i] == '#', false, nil, rightPot}
		// Setup pointer to the new pot
		rightPot.left = newPot
		// Shift pot over 1 for next iteration
		rightPot = newPot
	}
	// Gotta give myself a buffer
	s.referencePot = s.referencePot.addEmptyRightPot().addEmptyRightPot().addEmptyRightPot().addEmptyRightPot().addEmptyRightPot()
	newPot.addEmptyLeftPot().addEmptyLeftPot().addEmptyLeftPot().addEmptyLeftPot().addEmptyLeftPot()
}

func (s *machine) proceedNextGen() {
	pot := s.referencePot
	for pot != nil {
		pot.determineNextGen(s)
		pot = pot.left
	}

	pot = s.referencePot
	for pot != nil {
		pot.plant = pot.nextPlant
		if pot.left == nil && (pot.plant || pot.right.plant || pot.right.right.plant) {
			// Add 2 new pots to the left for more buffering
			pot.addEmptyLeftPot().addEmptyLeftPot()
		}
		pot = pot.left
	}

	// BTW, I am assuming that none of these pots can be null, since I already added 5 buffer pots to each side at the beginning
	if s.referencePot.plant || s.referencePot.left.plant || s.referencePot.left.left.plant {
		s.referencePot = s.referencePot.addEmptyRightPot().addEmptyRightPot()
	}
}

func (s *machine) getTotal() int {
	pot := s.referencePot
	var total int
	for pot != nil {
		if pot.plant {
			total += pot.id
		}
		pot = pot.left
	}

	return total
}

func (s *machine) getTotalPlants() int {
	pot := s.referencePot
	var total int
	for pot != nil {
		if pot.plant {
			total++
		}
		pot = pot.left
	}

	return total
}

func (s *machine) print() string {
	pot := s.referencePot
	output := ""
	for pot != nil {
		if pot.plant {
			output = "#" + output
		} else {
			output = "." + output
		}
		pot = pot.left
	}
	return output
}

func (p *pot) addEmptyRightPot() *pot {
	p.right = &pot{p.id + 1, false, false, p, nil}
	return p.right
}

func (p *pot) addEmptyLeftPot() *pot {
	p.left = &pot{p.id - 1, false, false, nil, p}
	return p.left
}

func (p *pot) determineNextGen(s *machine) {
	stateString := fmt.Sprintf("%s%s%s%s%s",
		string(s.plantMap[p.left != nil && p.left.left != nil && p.left.left.plant]),
		string(s.plantMap[p.left != nil && p.left.plant]),
		string(s.plantMap[p.plant]),
		string(s.plantMap[p.right != nil && p.right.plant]),
		string(s.plantMap[p.right != nil && p.right.right != nil && p.right.right.plant]))
	p.nextPlant = s.stateMap[stateString]
}
