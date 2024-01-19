// armstrong package contains the solution for Armstrong Number exercise on Exercism.
package armstrong

import (
	"math"
)

// IsNumber takes a number and returns true or false based on if the number is an armstrong number or not.
func IsNumber(n int) bool {

	if n < 0 {
		return false // Numbers below 0 can not be armstrong numbers.
	}

	var sum int
	nLength := getNumLength(n) // get the number of digits in the passed number.

	for i := n; i > 0; i = i / 10 { // we divide the number by 10, as number is an integer, it drops the decimal value.

		sum += intPow(i%10, nLength) // we get the curret last digit of the number using modulo 10.
	}

	return n == sum
}

// getNumLength takes a number and returns the numbers of digits in that number
func getNumLength(n int) int {
	if n == 0 {
		return 1 // if the number is zero then return 1, this is here as our formula doesn't work with '0'
	}
	// we use (log base 10 number) + 1 to get the number of digits, as its converted back to integer, we don't have to
	// worry about decimal values.
	return int(math.Log10(float64(n))) + 1
}

// intPow is basically using math.pow function but works with integer values.
func intPow(num, pow int) int {
	return int(math.Pow(float64(num), float64(pow)))
}
