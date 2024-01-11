// prime package contains solution for the prime problem on Exercism
package prime

import "fmt"

// Nth take a number 'n' and returns the nth prime number.
func Nth(n int) (int, error) {

	// Check if n is less then 1, if so return error.
	if n < 1 {
		return 0, fmt.Errorf("prime can't be calculated for negative or zero: %d", n)
	}

	primeCount := 0
	num := 0

	for primeCount < n {
		num++

		// check if the current iterating number is prime, if so then increment primeCount.
		if isPrime(num) {
			primeCount++
		}
	}
	return num, nil
}

// Prime takes a number 'n' and return true or false based on if its prime or not.
func isPrime(n int) bool {
	// any number less then 2 is not a prime.
	if n < 2 {
		return false
	}

	// here we check if a number is prime by multiplying it with all numbers between 2 and sqrt(n)
	// and check if we find a multiple. we check only till sqrt(n) and not till 'n' because
	// if a number is not a prime then one if it factors has to be less then sqrt(n).
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
