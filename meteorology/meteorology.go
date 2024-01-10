// triangle package contains tools to help us decern between different type of triangles
package triangle

import "math"

// Kind type refers to the triangle's type
type Kind int

// Defining various types of triangles
const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {

	// handle case where invalid input is provided.
	for _, side := range []float{a, b, c} {
		if side <= 0 || math.IsNaN(side) || math.IsInf(side, 0) {
			return NaT
		}
	}

	switch {
	// handle the case where a side is bigger then other two side combined.
	case a+b < c || c+b < a || a+c < b:
		return NaT // not a triangle

	// check for equilateral triangle.
	case a == b && b == c:
		return Equ // equilateral triangle

	// check for isosceles triangle .
	case a == b || b == c || c == a:
		return Iso // isosceles triangle

	// if non of the above case match then we know its scalene.
	default:
		return Sca // scalene triangle
	}
}
