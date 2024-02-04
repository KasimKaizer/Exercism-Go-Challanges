package allyourbase

import (
	"errors"
	"math"
)

var (
	errInvOutBase = errors.New("output base must be >= 2")
	errInvInBase  = errors.New("input base must be >= 2")
	errInvInput   = errors.New("all digits must satisfy 0 <= d < input base")
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {

	if inputBase < 2 {
		return nil, errInvInBase
	}

	if outputBase < 2 {
		return nil, errInvOutBase
	}

	inputB10, err := convToB10(inputBase, inputDigits)
	if err != nil {
		return nil, err
	}

	return convFromB10(outputBase, inputB10), nil
}

// convToB10 will convert a value from any base to base 10
func convToB10(base int, val []int) (int, error) {
	base10Val := 0
	valLen := len(val) - 1
	for idx, num := range val {
		if num >= base || num < 0 {
			return 0, errInvInput
		}
		base10Val += num * intPow(base, valLen-idx)
	}
	return base10Val, nil
}

func convFromB10(base, b10Val int) []int {
	if b10Val == 0 {
		return []int{0}
	}
	convVal := make([]int, 0)
	for b10Val > 0 {
		i := b10Val % base
		convVal = append([]int{i}, convVal...)
		b10Val /= base
	}
	return convVal
}

func intPow(num, exp int) int {
	return int(math.Pow(float64(num), float64(exp)))
}
