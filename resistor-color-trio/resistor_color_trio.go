// Package resistorcolortrio contains various tools to calculate resistance of a resistance.
package resistorcolortrio

import (
	"fmt"
	"math"
)

// resistance types denotes resistance of a resistor.
type resistance int

const (
	kilo = 1e3 // couldn't figure out how to do this with iota.
	mega = 1e6
	giga = 1e9
	tera = 1e12
)

// resistanceMap links colors to the resistance value they represent.
var resistanceMap = map[string]resistance{
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

// string method represents the resistance type as a string.
func (r resistance) string() string {
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
		return resistance(0).string()
	}

	if colLen < 2 {
		// edge case, handle colors slice with only one color by retuning the resistance mapping
		// to that color.
		return resistanceMap[colors[0]].string()
	}

	res := resistance(0)
	for idx, color := range colors[:2] {
		res += resistanceMap[color] * intPow(10, 1-idx)
	}

	if colLen < 3 {
		// edge case, if there are only two colors in the colors array then return the
		// the current res value.
		return res.string()
	}

	res *= intPow(10, int(resistanceMap[colors[2]]))
	return res.string()
}

// intPow function calls math.Pow func on passed values, but it takes int values and returns
// type resistance value.
func intPow(num, exp int) resistance {
	return resistance(math.Pow(float64(num), float64(exp)))
}
