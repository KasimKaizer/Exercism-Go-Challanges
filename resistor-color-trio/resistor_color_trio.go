// Package resistorcolortrio contains solution for Resistor Color Trio Exercise on Exercism.
package resistorcolortrio

import (
	"fmt"
	"math"
)

// resistor types denotes resistance of a resistor.
type resistor int

// metric ohms values.
const (
	kilo = 1e3 // couldn't figure out how to do this with iota.
	mega = 1e6
	giga = 1e9
	tera = 1e12
)

// resistance map links colors to the resistance they represent.
var resistance = map[string]resistor{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// string method represents the resistor type as a string.
func (r resistor) string() string {
	switch {
	case r >= tera:
		return fmt.Sprintf("%d teraohms", r/tera)
	case r >= giga:
		return fmt.Sprintf("%d gigaohms", r/giga)
	case r >= mega:
		return fmt.Sprintf("%d megaohms", r/mega)
	case r >= kilo:
		return fmt.Sprintf("%d kiloohms", r/kilo)
	}
	return fmt.Sprintf("%d ohms", r)
}

// Label takes a list of colors and returns the resistance value they represent.
func Label(colors []string) string {
	colLen := len(colors)

	if colLen == 0 {
		// edge case, handle empty colors slice by returning 0.
		return resistor(0).string()
	}

	if colLen < 2 {
		// edge case, handle colors slice with only one color by retuning the resistance mapping
		// to that color.
		return resistance[colors[0]].string()
	}

	res := resistor(0)
	for idx, color := range colors[:2] { // loop through first two colors.
		// add the resistance of current color multiplied by 10 ^ the color's position
		// to the final res value.
		res += resistance[color] * intPow(10, 1-idx)
	}

	if colLen < 3 {
		// edge case, if there are only two colors in the colors array then return the
		// the current res value.
		return res.string()
	}

	// add the amount of zeros represented by the third color of the colors array to res value.
	res *= intPow(10, int(resistance[colors[2]]))
	return res.string()
}

// intPow function calls math.Pow func on passed values, but it takes int values and returns
// type resistor value.
func intPow(num, exp int) resistor {
	return resistor(math.Pow(float64(num), float64(exp)))
}
