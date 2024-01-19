// phonenumber - contains the solution for Phone Number exercise on Exercism.
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

// Number - takes a phone number in any format and cleans it up.
func Number(phoneNumber string) (string, error) {
	number := re.FindAllString(phoneNumber, -1) // find all digits in the passed string.
	numLen := len(number)

	switch numLen {
	case 11: // cases for if the number of digits in the passed string is 11
		if number[0] != "1" {
			// first digit can't be anything but '1' in american phone system.
			return "", errInvalidNumber
		}

		if number[1] == "1" || number[1] == "0" {
			return "", errInvalidNumber // second digit can't be 0 or 1.
		}

		if number[4] == "1" || number[4] == "0" {
			return "", errInvalidNumber // 5th digit can't be 0 or 1.
		}

		return strings.Join(number[1:], ""), nil

	case 10: // cases for if the number of digits in the passed string is 10
		if number[0] == "1" || number[0] == "0" {
			return "", errInvalidNumber // first digit can't be 0 or 1.
		}

		if number[3] == "1" || number[3] == "0" {
			return "", errInvalidNumber // 4th digit can't be 0 or 1.
		}
		return strings.Join(number, ""), nil

	}
	// if number of digits is not 10 or 11 then the number is invalid.
	return "", errInvalidNumber
}

// AreaCode - takes a phone number in any formate and returns its area code.
func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return number[:3], nil
}

// Format - takes an unformatted phone number and formats it.
func Format(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", number[:3], number[3:6], number[6:]), nil
}
