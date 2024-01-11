// lsproduct contains solution for the Largest Series product exercise on Exercism.
package lsproduct

import (
	"errors"

	"fmt"
)

// getProduct takes a strings of digits and returns there product.
func getProduct(series string) (int64, error) {

	product := int64(1)

	for _, num := range series {

		// handle case for non numeric character inside the provided string
		if num < '0' || num > '9' {
			return 0, fmt.Errorf("invalid character inside the provided input: %c", num)
		}
		product *= int64(num - '0')
	}
	return product, nil
}

// LargestSeriesProduct takes a string of digits and the span returns the biggest product inside
// the string according to that span.
func LargestSeriesProduct(digits string, span int) (int64, error) {

	// handle the case where the provided digit string is empty
	if digits == "" {
		return 0, errors.New("digits cannot be empty")
	}

	lengthOfDigits := len(digits)

	// handle the case where span is too small or too big
	if span < 1 || span > lengthOfDigits {
		return 0, errors.New("span provided cannot be less then 1 or above the number of digits")
	}

	var largestProduct int64

	for i, j := 0, span; i < lengthOfDigits; i, j = i+1, j+1 {

		// return early if we can't create any more series
		if j > lengthOfDigits {
			return largestProduct, nil
		}

		series := digits[i:j]

		currentProduct, err := getProduct(series)

		if err != nil {
			return 0, err
		}

		if largestProduct < currentProduct {
			largestProduct = currentProduct
		}

	}

	return largestProduct, nil

}
