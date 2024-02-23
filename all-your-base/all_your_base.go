// Package allyourbase contains solution for All Your Base  Exercise on Exercism.
package allyourbase

import (
	"errors"
	"math"
)

// errors for various invalid values.
var (
	errInvOutBase = errors.New("output base must be >= 2")
	errInvInBase  = errors.New("input base must be >= 2")
	errInvInput   = errors.New("all digits must satisfy 0 <= d < input base")
)

// ConvertToBase takes a value, its current base,  and base you want to convert that value into and
// converts the value to that base.
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {

	if inputBase < 2 {
		return nil, errInvInBase // input base can't be lower than 2
	}

	if outputBase < 2 {
		return nil, errInvOutBase // output base can't be lower than 2
	}

	inputB10, err := convToB10(inputBase, inputDigits) // convert the input value into base 10
	if err != nil {
		return nil, err
	}

	return convFromB10(outputBase, inputB10), nil // convert the base 10 value to output base.
}

// convToB10 takes a value and its base and returns that value represented in base10.
func convToB10(base int, val []int) (int, error) {
	base10Val := 0
	valLen := len(val) - 1
	for idx, num := range val { // loop through all the values in val.
		if num >= base || num < 0 {
			// if the current number is greater than or equal to its base then
			// input value if invalid.
			return 0, errInvInput
		}
		// multiply current number with its base ^ its place. and add it to base10val.
		// ex: if base was 6 and val was [2,3,4]
		// then to get base 10 value the formula would be: (2 * 6^2) + (3 * 6^1) + (4 *6^0)
		base10Val += num * intPow(base, valLen-idx)
	}
	return base10Val, nil
}

// convFromB10 takes a value in base 10 and the base the value should be converted into and converts
// the value to that base.
func convFromB10(base, b10Val int) []int {
	if b10Val == 0 {
		return []int{0} // if the value is 0 then return 0 inside a slice.
	}
	convVal := make([]int, 0)
	for b10Val > 0 { // loop for as long as base10 value is greater then 0.
		// append the reminder of the base10 value divided by the base we want to convert it into
		// to the front of the convVal slice.
		convVal = append([]int{(b10Val % base)}, convVal...)
		// update the base10 value by dividing it with the base we want to convert it into.
		b10Val /= base
	}
	return convVal
}

// intPow use math.Pow but takes and returns int values.
func intPow(num, exp int) int {
	return int(math.Pow(float64(num), float64(exp)))
}
