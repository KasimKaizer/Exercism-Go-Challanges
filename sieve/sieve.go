// sieve Package contains solution for the Sieve exercise on Exercism
package sieve

// Sieve returns a list of all the prime numbers till the given limit.
func Sieve(limit int) []int {

	// return nil as all numbers below 2 are non prime (composite) numbers.
	if limit < 2 {
		return nil
	}

	composite := make([]bool, limit+1) // make an array for all numbers till the given limit.

	composite[0] = true // 0 & 1 are composite numbers so can be marked as such.
	composite[1] = true

	var prime []int // slice of prime numbers which we will return.

	// we loop over all the numbers in the composite array.
	for num, isComposite := range composite {
		// if the current number is not composite.
		if !isComposite {
			prime = append(prime, num) // append that number to prime slice.
			for i := 2; i*num <= limit; i++ {
				// mark all the multiples of that number in composite array  as composite.
				composite[i*num] = true
			}
		}
	}

	return prime // return the slice of prime numbers.
}
