// atbash contains solution for the atbash cipher problem on Exercism.

// first version

package atbash

import "strings"

// Atbash takes a string and returns an encrypted string which is encrypted based on Atbash Cipher.
func Atbash(s string) string {

	s = strings.ToLower(s)
	var output strings.Builder
	var count int

	// we range over the whole string, we skip over any char which is not a letter or a number.
	for _, char := range s {

		// if character is a letter, shift it according to abash cipher, append it to output and
		// increment the count.
		if char >= 'a' && char <= 'z' {
			output.WriteRune('a' + 'z' - char)
			count++
			addSpace(&output, count)
		}

		// if character is a number, just append it to the output and increment the count.
		if char >= '0' && char <= '9' {
			output.WriteRune(char)
			count++
			addSpace(&output, count)
		}
	}
	return strings.TrimSpace(output.String())
}

// addSpace takes a count and a string and appends a space to the string based on the count.
func addSpace(str *strings.Builder, count int) {
	if count%5 == 0 {
		str.WriteRune(' ')
	}
}
