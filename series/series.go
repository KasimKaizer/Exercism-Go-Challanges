// Package series contains solution for the Series Exercise on Exercism
package series

// All takes a span and a string returns all the contiguous substrings of length span in that string
// in the order that they appear.
func All(n int, s string) []string {
	strLen := len(s)

	if strLen < n {
		return nil // return nil is the span is higher then length of the string.
	}

	var output []string
	// only loop till the last possible character which can create a series of length span.
	for i := 0; i < strLen-n+1; i++ {
		output = append(output, s[i:i+n])
	}

	return output
}

// UnsafeFirst takes a string and a span and returns the first substring with that span.
// if the string provided is smaller then span, it returns the string.
func UnsafeFirst(n int, s string) string {
	output, _ := First(n, s)
	return output
}

// First similar to UnsafeFirst function takes a string and a span and returns the first substring
// with that span. if the string provided is smaller then span, it returns the string and `false`.
func First(n int, s string) (string, bool) {
	if len(s) < n {
		return s, false
	}

	return s[:n], true
}
