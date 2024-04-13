// Package perfect contains solution for Perfect Numbers exercise on Exercism.
package perfect

import "errors"

// Classification defines weather a number is perfect or not.
type Classification int

const (
	ClassificationPerfect Classification = iota + 1
	ClassificationAbundant
	ClassificationDeficient
)

// error for a negative or zero number.
var ErrOnlyPositive = errors.New("only positive non zero numbers are allowed")

// Classify takes a number and returns its calcification.
func Classify(n int64) (Classification, error) {
	if n < 1 {
		return 0, ErrOnlyPositive
	}
	if n == 1 {
		return ClassificationDeficient, nil // 1 is Deficient
	}

	var sum int64 = 1
	for i := int64(2); i*i <= n; i++ {
		if (n % i) != 0 {
			continue
		}
		sum += i
		compFac := (n / i) // get the complementary factor of the current factor.
		if i != compFac {
			sum += compFac
		}
		if sum > n {
			return ClassificationAbundant, nil
		}
	}
	if sum < n {
		return ClassificationDeficient, nil
	}
	return ClassificationPerfect, nil
}
