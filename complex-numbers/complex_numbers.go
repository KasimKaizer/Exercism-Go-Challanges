// Package complexnumbers contains solution for the Complex Numbers exercise on Exercism.
package complexnumbers

import "math"

// Number defines a complex number.
type Number struct {
	real float64
	imag float64
}

// NewNumber returns a new complex number with given real and imaginary values.
func NewNumber(real, imag float64) Number {
	return Number{real: real, imag: imag}
}

// Real method returns the real part of the complex number.
func (n Number) Real() float64 {
	return n.real
}

// Imaginary method returns the imaginary part of the complex number.
func (n Number) Imaginary() float64 {
	return n.imag
}

// Add method adds the provided complex number to the complex
// number the method is being called on and returns the result.
func (n1 Number) Add(n2 Number) Number {
	n1.real += n2.real
	n1.imag += n2.imag
	return n1
}

// Subtract method subtracts the provided complex number with the complex
// number the method is being called on and returns the result.
func (n1 Number) Subtract(n2 Number) Number {
	n1.real -= n2.real
	n1.imag -= n2.imag
	return n1
}

// Multiply method multiples the provided complex number with the complex
// number the method is being called on and returns the result.
// formula: (a + bi) (c + di) = (ac - bd) + (ad + bc)i
func (n1 Number) Multiply(n2 Number) Number {
	realN := n1.real*n2.real - n1.imag*n2.imag
	imagN := n1.real*n2.imag + n1.imag*n2.real
	return NewNumber(realN, imagN)
}

// Times method multiples the complex number by the factor provided.
func (n Number) Times(factor float64) Number {
	n.real *= factor
	n.imag *= factor
	return n
}

// Divide method divides the provided complex number with the complex
// number the method is being called on and returns the result.
// formula: (a + i * b) / (c + i * d) =
// (a * c + b * d)/(c^2 + d^2) + (b * c - a * d)/(c^2 + d^2) * i
func (n1 Number) Divide(n2 Number) Number {
	devisor := (n2.real * n2.real) + (n2.imag * n2.imag)
	realN := (n1.real*n2.real + n1.imag*n2.imag) / devisor
	imagN := (n1.imag*n2.real - n1.real*n2.imag) / devisor
	return NewNumber(realN, imagN)
}

// Conjugate method returns the conjugate of the complex number it's being
// called on.
func (n Number) Conjugate() Number {
	n.imag = -n.imag
	return n
}

// Abs method returns the absolute values of the complex number it's being
// called on.
func (n Number) Abs() float64 {
	return math.Sqrt((n.real * n.real) + (n.imag * n.imag))
}

// Exp returns the base-e exponential values of the complex number it's
// being called on.
// formula: e^(a + i * b) = e^a * e^(i * b);  e^(i * b) = cos(b) + i * sin(b)
func (n Number) Exp() Number {
	expA := math.Exp(n.real)
	cosB := math.Cos(n.imag)
	sinB := math.Sin(n.imag)
	return NewNumber((expA * cosB), (expA * sinB))
}
