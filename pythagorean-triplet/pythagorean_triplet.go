// Package pythagorean contains solution for exercise on Exercism.
package pythagorean

import "math"

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	output := make([]Triplet, 0)
	for a := min; a < max; a++ {
		for b := a + 1; b < max; b++ {
			c, ok := perfectSquare(a*a + b*b)
			if !ok {
				continue
			}
			if c > max {
				return output
			}
			output = append(output, Triplet{a, b, c})
		}
	}
	return output
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	output := make([]Triplet, 0)
	for _, tri := range Range(1, p) {
		if (tri[0] + tri[1] + tri[2]) == p {
			output = append(output, tri)
		}
	}
	return output
}

// perfectSquare returns underroot x and if x is a perfect square root.
func perfectSquare(x int) (int, bool) {
	y := int(math.Sqrt(float64(x)))
	return y, (y * y) == x

}
