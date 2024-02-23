// Package allergies contains solution for Allergies exercise on exercism.
package allergies

// allergyList is map linking allergies to their allergies values.
var allergyList = map[string]uint{
	"cats":         128,
	"pollen":       64,
	"chocolate":    32,
	"tomatoes":     16,
	"strawberries": 8,
	"shellfish":    4,
	"peanuts":      2,
	"eggs":         1,
}

// Allergies takes the allergies number and returns a list of the allergies that number
// represents.
func Allergies(allergies uint) []string {
	output := make([]string, 0)
	for allergy := range allergyList {
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
	allergyNum := allergyList[allergen]
	return (allergies & allergyNum) == allergyNum
}

/*
// Just for fun, different approach without doing any bit-shifting.
// This approach according to my testing is twice as fast when just creating Allergies list.
// but not so fast when retrieving result from AllergicTo function.

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
