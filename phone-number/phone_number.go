// Package phonenumber contains the solution for Phone Number exercise on Exercism.
package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	// regex to detect all the digits in a string.
	// We compile our regex here to prevent any errors during runtime.
	re               = regexp.MustCompile(`\d`)
	errInvalidNumber = errors.New("the number provided is invalid")
)

// Number takes a phone number in any format and cleans it up.
func Number(phoneNumber string) (string, error) {
	number := re.FindAllString(phoneNumber, 12) // find all digits in the passed string.
	numLen := len(number)

	if numLen > 11 || numLen < 10 {
		return "", errInvalidNumber // a number can't be above 11 digits or below 10 digits in NANP.
	}

	if numLen == 11 {
		if number[0] != "1" {
			return "", errInvalidNumber // first digit of a 11 digit number can't be anything but '1' in NANP.
		}
		number = number[1:]
	}
	if number[0] == "1" || number[0] == "0" {
		return "", errInvalidNumber // Area Code can't start with 0 or 1 in NANP.
	}

	if number[3] == "1" || number[3] == "0" {
		return "", errInvalidNumber // Exchange code can't start with 0 or 1 in NANP.
	}
	return strings.Join(number, ""), nil
}

// AreaCode takes a phone number in any formate and returns its area code.
func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return number[:3], nil
}

// Format takes an unformatted phone number and formats it.
func Format(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", number[:3], number[3:6], number[6:]), nil
}
