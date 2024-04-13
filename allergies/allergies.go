// Package allergies contains tools to detect allergies a person might have.
package allergies

// allergyList is list of all possible allergies
var allergyList = []string{"cats", "pollen", "chocolate", "tomatoes", "strawberries", "shellfish", "peanuts", "eggs"}

// allergyList links allergies to their allergies values.
func allergiesToNum(allergy string) uint {
	switch allergy {
	case "cats":
		return 128
	case "pollen":
		return 64
	case "chocolate":
		return 32
	case "tomatoes":
		return 16
	case "strawberries":
		return 8
	case "shellfish":
		return 4
	case "peanuts":
		return 2
	case "eggs":
		return 1
	}
	return 0
}

// Allergies takes the allergies number and returns a list of the allergies that number
// represents.
func Allergies(allergies uint) []string {
	output := make([]string, 0)
	for _, allergy := range allergyList {
		if !AllergicTo(allergies, allergy) {
			continue
		}
		output = append(output, allergy)
	}
	return output
}

// AllergicTo takes allergies number and a specific allergy and returns true/false if the input
// allergy is represented by that allergies number.
func AllergicTo(allergies uint, allergen string) bool {
	allergyNum := allergiesToNum(allergen)
	return (allergies & allergyNum) == allergyNum
}

/*
// Just for fun, different approach without doing any bit-shifting.

import "slices"

type allergy struct {
	name  string
	score uint
}

var allergyList = []allergy{
	{"cats", 128},
	{"pollen", 64},
	{"chocolate", 32},
	{"tomatoes", 16},
	{"strawberries", 8},
	{"shellfish", 4},
	{"peanuts", 2},
	{"eggs", 1},
}

func Allergies(allergies uint) []string {
	allergies %= allergyList[0].score * 2

	output := make([]string, 0)
	for _, item := range allergyList {
		if item.score > allergies {
			continue
		}
		allergies -= item.score
		output = append(output, item.name)
	}
	return output
}

func AllergicTo(allergies uint, allergen string) bool {
	list := Allergies(allergies)
	return slices.Contains(list, allergen)
}
*/
