// prime package contains solution for Prime Factor exercise on Exercism.
package prime

// Factors - takes a number and returns all of its prime factors.
func Factors(n int64) []int64 {
	var output []int64

	// there are no prime numbers less then 2, so return an empty slice.
	if n < 2 {
		return output
	}

	// as long as the number is devisable by 2, append 2 to the output string.
	for n%2 == 0 {
		output = append(output, 2)
		n /= 2
	}

	// as now 2 and all its multiples (even numbers) are out of the picture, we iterate through all
	// the odd numbers till SquareRoot(n).
	for i := int64(3); i*i <= n; i = i + 2 {
		// if one of those numbers is devisable by n, divide the number from n till its no longer
		// devisable by n.
		for n%i == 0 {
			output = append(output, i) // append 'i' to output each time its successfully devisable.
			n /= i
		}
	}

	// if after all the above the number is still not '1' then the leftover n is a prime number and
	// our last prime factor.
	if n > 1 {
		output = append(output, n) // append it to the output.
	}
	return output
}
