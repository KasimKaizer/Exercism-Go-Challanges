// Package yacht contains solution for the Yacht exercise on Exercism.
package yacht

import (
	"slices"
)

const (
	invalidScore     = 0
	straightScore    = 30
	yachtScore       = 50
	lilStraightStart = 1
	bigStraightStart = 2
)

// map names of numbers to there numeric values.
var nameToNum = map[string]int{
	"ones":   1,
	"twos":   2,
	"threes": 3,
	"fours":  4,
	"fives":  5,
	"sixes":  6,
}

// Score takes the result of dice rolls and the category, it returns the score based on the
// Yacht game's scoring.
func Score(dice []int, category string) int {
	if num, ok := nameToNum[category]; ok {
		return numberFreq(dice, num)
	}
	switch category {
	case "full house":
		return checkHouse(dice)
	case "four of a kind":
		return checkFour(dice)
	case "little straight":
		return lilStraight(dice)
	case "big straight":
		return bigStraight(dice)
	case "choice":
		return choice(dice)
	case "yacht":
		return yachtCheck(dice)
	}
	return invalidScore
}

// numberFreq returns the score for the dice rolls based on provided number's category of yacht game.
func numberFreq(rolls []int, number int) int {
	score := 0

	for _, num := range rolls {
		if num != number {
			continue
		}
		score += number
	}
	return score
}

// lilStraight checks and returns the score for the dice rolls based on little straight category of yacht game.
func lilStraight(rolls []int) int {
	newRolls := copyAndSort(rolls) // prevent mutation of the original slice.
	if newRolls[0] != lilStraightStart {
		return invalidScore
	}
	return straight(newRolls)
}

// bigStraight checks and returns the score for the dice rolls based on Big straight category of yacht game.
func bigStraight(rolls []int) int {
	newRolls := copyAndSort(rolls) // prevent mutation go original slice.
	if newRolls[0] != bigStraightStart {
		return invalidScore
	}
	return straight(newRolls)
}

// straight checks for a straight, if any number in the rolls array except the first number is not the
// successor of previous number then it returns invalidScore, otherwise it returns the straightScore.
func straight(rolls []int) int {

	for i := 1; i < len(rolls); i++ {
		if rolls[i-1]+1 != rolls[i] {
			return invalidScore
		}
	}

	return straightScore
}

// checkHouse checks and returns the score for the dice rolls based on full house category of yacht game.
func checkHouse(rolls []int) int {
	newRolls := copyAndSort(rolls) // prevent mutation for original slice.
	if newRolls[0] == newRolls[len(newRolls)-1] {
		return invalidScore
	}
	twoPart, threePart := newRolls[:2], newRolls[2:]
	if newRolls[2] != newRolls[3] {
		threePart, twoPart = newRolls[:3], newRolls[3:]
	}
	twoS, ok := freqCheck(twoPart, 2)
	if !ok {
		return invalidScore
	}
	threeS, ok := freqCheck(threePart, 3)
	if !ok {
		return invalidScore
	}
	return twoS + threeS
}

// freqCheck checks if all the numbers till the provided length in the slice are same,
// if they are then it returns the sum of them all.
func freqCheck(rolls []int, length int) (int, bool) {
	sum := 0
	for i := 0; i < length; i++ {
		if rolls[i] != rolls[0] {
			return invalidScore, false
		}
		sum += rolls[i]
	}
	return sum, true
}

// yachtCheck checks and returns the score for the dice rolls based on yacht category of yacht game.
func yachtCheck(rolls []int) int {
	if _, ok := freqCheck(rolls, len(rolls)); !ok {
		return invalidScore
	}
	return yachtScore
}

// checkFour checks and returns the score for the dice rolls based on four of a kind category of yacht game.
func checkFour(rolls []int) int {
	newRolls := copyAndSort(rolls) // prevent mutation for original slice.

	start := 0
	for newRolls[start] != newRolls[start+1] {
		start++
		if start > 1 {
			return invalidScore
		}
	}
	score, ok := freqCheck(newRolls[start:], 4)
	if !ok {
		return invalidScore
	}
	return score
}

// choice checks and returns the score for the dice rolls based on choice category of yacht game.
func choice(rolls []int) int {
	sum := 0
	for _, roll := range rolls {
		sum += roll
	}
	return sum
}

// copyAndSort returns a sorted copy of the provided slice.
func copyAndSort(slice []int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	slices.Sort(newSlice)
	return newSlice
}
