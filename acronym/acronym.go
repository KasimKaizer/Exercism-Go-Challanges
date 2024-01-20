// acronym - package contains solution for Acronym exercise on Exercism
package acronym

import (
	"regexp"
	"strings"
)

// let re & r be evaluated at compile time.
var (
	// regex to get the first letter of the sentence, and any letter thats succeeds a whiteSpace or "-" or "_"
	re = regexp.MustCompile(`(^|\s|-|_)[a-zA-Z]`)
	// replacer to remove all whiteSpaces, "-" and "_"
	r = strings.NewReplacer(" ", "", "_", "", "-", "")
)

// Abbreviate - takes a phrase and turns it into its acronym.
func Abbreviate(s string) string {
	s = strings.Join(re.FindAllString(s, -1), "")
	return strings.ToUpper(r.Replace(s))
}
