// Package dndcharacter contains solution for D&D Character exercise on Exercism.
package dndcharacter

import (
	"math"
	"math/rand"
)

// Character represents a character with there stats in D&D.
type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier returns the constitution modifier for the passed score.
func Modifier(score int) int {
	// use maths.Floor method to round down the value and then convert it to int.
	return int(math.Floor((float64(score) - 10) / 2.0))
}

// Ability generates a random score for the players ability.
func Ability() int {
	smallestRoll, sum := 6, 0 // initialize smallestRoll as largest roll possible.
	for i := 0; i < 4; i++ {  // roll 4 times.
		curRoll := diceRoll()
		sum += curRoll // add all the resulting values to sum.
		// check if current roll was smaller, if yes then save current roll as smallest roll.
		if curRoll < smallestRoll {
			smallestRoll = curRoll
		}
	}
	return sum - smallestRoll // remove the smallest roll from the sum and return.
}

// diceRoll rolls a dice of 6 sides and returns the result.
func diceRoll() int {
	return rand.Intn(6) + 1
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	Const := Ability() // get the Constitution's score.
	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Const,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		// get Modifier using Constitution value and add 10 to it as players initial HitPoints.
		Hitpoints: (10 + Modifier(Const)),
	}
}
