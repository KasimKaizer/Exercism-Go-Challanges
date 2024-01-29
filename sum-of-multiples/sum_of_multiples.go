// Package summultiples contains solution for Sum of Multiples on Exercism.
package summultiples

// SumMultiples returns the energy points player earned after completing a level.
// it takes the player's level (this is the limit) and an array of base values,
// and returns a sum of all multiples of those base values till the limit, excluding any
// repeating multiples.
func SumMultiples(limit int, divisors ...int) int {
	numbers := make([]bool, limit) // create an array till our limit.
	sum := 0

	for _, div := range divisors { // range over all base values.

		if div == 0 {
			continue // skip 0. 0 would lead to infinite loop.
		}

		for i := 1; (i * div) < limit; i++ {
			// if the current multiple has been added to sum before, continue to next value.
			if numbers[i*div] {
				continue
			}

			numbers[i*div] = true // set the position of the current multiple in numbers array to true.
			sum += (i * div)      // add the current multiple to sum.
		}
	}

	return sum
}
