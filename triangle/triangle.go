// triangle package contains tools to help us decern between different type of triangles
package triangle

// Kind type is an alias for string used to return the triangle's type
type Kind string

// Defining the various types of triangles
const (
	NaT Kind = "NaT" // not a triangle
	Equ Kind = "Equ" // equilateral
	Iso Kind = "Iso" // isosceles
	Sca Kind = "Sca" // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	switch {
	// handle the case where any given side is less than or equal to zero.
	case a <= 0 || b <= 0 || c <= 0:
		return NaT // not a triangle

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
