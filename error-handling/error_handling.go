// Package erratum contains solution for Error Handling exercise on Exercism.
package erratum

import (
	"errors"
)

// Use executes forb method on the returned value of opener.
func Use(opener ResourceOpener, input string) (err error) {
	recourse, err := opener()
	if errors.As(err, &TransientError{}) { // recurse on transient error.
		return Use(opener, input)
	}
	if err != nil {
		return
	}
	defer recourse.Close()
	defer func() { // handle panic case
		r, ok := recover().(error)
		if !ok {
			return
		}
		if errors.As(r, &FrobError{}) {
			recourse.Defrob(r.(FrobError).defrobTag)
		}
		err = r
	}()

	recourse.Frob(input)
	return
}
