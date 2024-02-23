// Package perfect contains solution for Perfect Numbers exercise on Exercism.
package perfect

import "errors"

// Classification defines weather a number is perfect or not.
type Classification int

const (
	ClassificationPerfect   Classification = iota // classification for perfect number
	ClassificationAbundant                        // classification for Abundant number
	ClassificationDeficient                       // classification for Deficient number
)

// error for a negative or zero number.
var ErrOnlyPositive = errors.New("only positive non zero numbers are allowed")

// Classify takes a number and returns its calcification.
func Classify(n int64) (Classification, error) {
	if n < 1 {
		return 0, ErrOnlyPositive // perfect number can't be negative.
	}
	if n == 1 {
		return ClassificationDeficient, nil // 1 is Deficient
	}
	// sum would be the sum of all factor of n, as 1 is factor for all numbers we start with that.
	var sum int64 = 1
	for i := int64(2); i*i <= n; i++ { // we start with 2 till underroot n
		if (n % i) != 0 {
			continue // current number is not a factor, continue the loop.
		}
		sum += i           // add the factor to the sum.
		compFac := (n / i) // get the complementary factor of the current factor.
		if i != compFac {  // skip this condition if complementary factor  is same as factor.
			sum += compFac // add the complementary factor as well to the sum
		}
		if sum > n {
			return ClassificationAbundant, nil // check if sum is greater then n. if so, return early.
		}
	}
	if sum < n {
		return ClassificationDeficient, nil
	}
	return ClassificationPerfect, nil
}
