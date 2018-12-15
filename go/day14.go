package main

import (
	"fmt"
	"strconv"
)

const (
	puzzleInput = 505961
)

type recipe struct {
	score, placement int
	next             *recipe
}

type recipeBoard struct {
	elfone *recipe
	elftwo *recipe
	first  *recipe
	last   *recipe
}

func day14() {
	curElfTwo := recipe{7, 2, nil}
	curElfOne := recipe{3, 1, &curElfTwo}
	recipeBoard := recipeBoard{&curElfOne, &curElfTwo, &curElfOne, &curElfTwo}

	for recipeBoard.last.placement < puzzleInput+10 {
		// recipeBoard.print()
		recipeBoard.generateMore()
	}

	// Part 1
	futureTen := ""
	curRecipe := recipeBoard.first
	for len(futureTen) < 10 {
		if curRecipe.placement > puzzleInput {
			futureTen += strconv.Itoa(curRecipe.score)
		}
		curRecipe = curRecipe.next
	}
	fmt.Printf("Future 10 recipes are %s.", futureTen)
	fmt.Println()

	// Part 2
	curRecipe = recipeBoard.first
	for {
		recipeToMatch := curRecipe
		sequence := strconv.Itoa(puzzleInput)
		matches := true
		for _, c := range sequence {
			score, _ := strconv.Atoi(string(c))
			matches = matches && recipeToMatch.score == score
			if recipeToMatch.next == nil {
				recipeBoard.generateMore()
			}
			recipeToMatch = recipeToMatch.next
		}

		if matches {
			fmt.Printf("%d recipes before matching sequence.", curRecipe.placement-1)
			fmt.Println()
			break
		}

		curRecipe = curRecipe.next
	}
}

func (b *recipeBoard) generateMore() {
	sum := b.elfone.score + b.elftwo.score

	tens := sum / 10
	ones := sum % 10
	if tens > 0 {
		b.appendRecipe(tens)
	}
	b.appendRecipe(ones)

	curElfOne := b.elfone
	curElfTwo := b.elftwo

	curElfOneScore := curElfOne.score
	for i := 0; i <= curElfOneScore; i++ {
		if curElfOne.next == nil {
			curElfOne = b.first
		} else {
			curElfOne = curElfOne.next
		}
	}

	curElfTwoScore := curElfTwo.score
	for i := 0; i <= curElfTwoScore; i++ {
		if curElfTwo.next == nil {
			curElfTwo = b.first
		} else {
			curElfTwo = curElfTwo.next
		}
	}

	b.elfone = curElfOne
	b.elftwo = curElfTwo
}

func (b *recipeBoard) appendRecipe(score int) {
	newRecipe := &recipe{score, b.last.placement + 1, nil}
	b.last.next = newRecipe
	b.last = newRecipe
}

func (b *recipeBoard) print() {
	recipes := ""
	curRecipe := b.first
	for curRecipe != nil {
		if curRecipe.placement > puzzleInput {
			recipes += strconv.Itoa(curRecipe.score)
		}
		curRecipe = curRecipe.next
	}
	fmt.Print(recipes)
	fmt.Println()
}
