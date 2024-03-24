// Package resistorcolortrio contains various tools to calculate resistance of a resistance.
package resistorcolortrio

import (
	"fmt"
	"math"
)

// resistance types denotes resistance of a resistor.
type resistance int

const (
	kilo = 1e3
	mega = 1e6
	giga = 1e9
	tera = 1e12
)

// resistanceMap links colors to the resistance value they represent.
func resistanceMap(color string) resistance {
	switch color {
	case "black":
		return 0
	case "brown":
		return 1
	case "red":
		return 2
	case "orange":
		return 3
	case "yellow":
		return 4
	case "green":
		return 5
	case "blue":
		return 6
	case "violet":
		return 7
	case "grey":
		return 8
	case "white":
		return 9
	default:
		return -1
	}
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
	res := resistance(0)
	for i := 0; i < 2 && i < len(colors); i++ {
		res = (res * 10) + resistanceMap(colors[i])
	}

	if len(colors) >= 3 {
		res *= intPow(10, int(resistanceMap(colors[2])))
	}

	return res.string()
}

// intPow function calls math.Pow func on passed values, but it takes int values and returns
// type resistance value.
func intPow(num, exp int) resistance {
	return resistance(math.Pow(float64(num), float64(exp)))
}
