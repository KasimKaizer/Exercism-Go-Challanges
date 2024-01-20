// Package armstrong contains the solution for Armstrong Number exercise on Exercism.
package armstrong

// IsNumber takes a number and returns true or false based on if the number is an armstrong number or not.
func IsNumber(n int) bool {

	var sum int
	nLength := numLength(n) // get the number of digits.

	for i := n; i > 0; i = i / 10 { // we divide the number by 10, as number is an integer, it drops the decimal value.

		sum += intPow(i%10, nLength) // we get the current last digit of the number using modulo 10.
	}

	return n == sum
}

func numLength(n int) int {
	l := 0
	for n > 0 {
		l++
		n /= 10
	}
	return l
}

// intPow returns number^pow for int numbers.
func intPow(num, pow int) int {
	p := num
	for i := 1; i < pow; i++ {
		p *= num
	}
	return p
}
