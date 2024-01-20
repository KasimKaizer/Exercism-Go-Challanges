// cipher - package contains the solution for the Simple Cipher problem on Exercism.
package cipher

import (
	"math/rand"
	"strings"
)

// lowLetters contains all lowercase letters in the alphabet.
const (
	lowLetters = "abcdefghijklmnopqrstuvwxyz"
	lowerCaseA = 97
	lowerCaseZ = 122
)

// shift - defines the value to shift characters by.
type shift int

// vigenere - defines the key to encode / decode a message with.
type vigenere string

// NewCaesar - creates a new shift which shifts values based on ceaser cipher.
func NewCaesar() Cipher {
	return shift(3)
}

// NewShift - creates a new shift that shifts the chars based on the distance provided.
func NewShift(distance int) Cipher {

	if distance < -25 || distance > 25 || distance == 0 {
		return nil
	}
	return shift(distance)
}

// Encode - takes a phrase and encodes it based on the char shift.
func (c shift) Encode(input string) string {
	return shiftCipher(input, c)
}

// Encode - takes a encoded phrase and decodes it based on the character shift.
func (c shift) Decode(input string) string {
	return shiftCipher(input, -c) // we call shiftCipher function with the inverse value of shift.
}

// NewVigenere - creates a new instance of vigenere with the key provided.
// If no key is provided then it creates a random key with length of 100.
func NewVigenere(key ...string) Cipher {

	lenSlice := len(key)
	if lenSlice > 1 || key[0] == "" {
		return nil // return nil if more then one key was provided.
	}

	if lenSlice == 0 {
		return vigenere(keyGen()) // return a randomly generated key if no key was provided.
	}

	count := 0
	for _, char := range []byte(key[0]) {
		if char < lowerCaseA || char > lowerCaseZ {
			return nil // return nil if an invalid character was in string.
		}
		count += int(char) - lowerCaseA //  sum the positions of characters in key.
	}

	if count == 0 {
		return nil // if count is zero then we know all characters in key are 'a'.
	}

	return vigenere(key[0])
}

// Encode - takes a phrase and encodes it with the help of the key in the vigenere.
// formula: C = (P + (K * 1))%26 where C = cipher, P = initial text, K = key string.
func (v vigenere) Encode(input string) string {
	return vigenereCipher(1, input, v) // we pass "1" here as key value has to be positive.
}

// Decode - takes an encoded phrase and decodes it with the help of the key in the vigenere.
// formula: P = (C + (K * -1))%26 where C = cipher, P = initial text, K = key string. if P is negative then P+=26.
func (v vigenere) Decode(input string) string {
	return vigenereCipher(-1, input, v) // we pass "-1" here as key value has to be negative.
}

// shiftCipher - shifts the position of the characters in the string according to the shift value provided.
func shiftCipher(str string, shift shift) string {
	str = strings.ToLower(str)
	var output strings.Builder

	for _, char := range []byte(str) {
		if char < lowerCaseA || char > lowerCaseZ {
			continue // skip any character that is not a lowercase letter.
		}
		output.WriteByte(shiftByte(char, shift))
	}
	return output.String()
}

// keyGen - creates a key which is 100 characters long and which consists of only random lower case letters.
func keyGen() string {
	k := make([]byte, 100)
	for i := range k {
		k[i] = lowLetters[rand.Intn(26)]
	}
	return string(k)
}

// vigenereCipher - takes a phrase and encodes/decodes it based on Vigenere Cipher method using the given inv and key.
func vigenereCipher(inv int, str string, key vigenere) string {
	str = strings.ToLower(str)
	var output strings.Builder
	count := 0

	for _, char := range []byte(str) {
		if char < lowerCaseA || char > lowerCaseZ {
			continue // skip any character that is not a lowercase letter.
		}
		keyByte := key[count%len(key)]
		output.WriteByte(shiftCes(charPos(char), charPos(keyByte)*inv))
		count++
	}
	return output.String()
}

// charPos - gets the position of the character in the alphabet.
// Example: a=0, b=1, c=2, d=3, e=4.....
func charPos(r byte) int {
	return int(r - byte(lowerCaseA))
}

// shiftCes - takes two shifts, adds them and returns a character shifted by the sum of those two shifts.
func shiftCes(x, y int) byte {

	r := (x + y) % 26 // characters position in alphabet can't go beyond 26.

	if r < 0 { // we roll over any position below zero here, Example: -1=25, -2=24, -3=23....
		// r + 26 + 97, where 97 is 'a' in bytes and 26 is the last position in alphabet.
		return byte(r + 26 + lowerCaseA)
	}
	return byte(r + lowerCaseA) // return the position added to the first character of alphabet aka 'a'.
}

// shiftByte - takes a character and a shift value and shifts it based on the shift value.
func shiftByte(r byte, s shift) byte {
	// we get the position of r in alphabet and send it to shiftCes with the shift value provided.
	return shiftCes(charPos(r), int(s))
}
