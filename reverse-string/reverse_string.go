// reverse - Package  contains method String.
package reverse

// Reverse - function has logic to reverse a given string.
func Reverse(s string) string {
	sRune := []rune(s)
	sRuneLen := len(sRune)

	for i := 0; i < (sRuneLen / 2); i++ {
		sRune[i], sRune[sRuneLen-i-1] = sRune[sRuneLen-i-1], sRune[i]
	}

	return string(sRune)
}
