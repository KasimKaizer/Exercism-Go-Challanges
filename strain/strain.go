// Package strain contains solution for the Strain exercise on Exercism.
package strain

// Keep takes a predicate and a slice of any type, and returns a slice of
// items from the passed list where predicate was true.
func Keep[T any](lst []T, fn func(T) bool) []T {
	var output []T
	for _, item := range lst {
		if !fn(item) {
			continue
		}
		output = append(output, item)
	}
	return output
}

// Discard takes a predicate and a slice of any type, and returns a slice of
// items from the passed list where predicate was false.
func Discard[T any](lst []T, fn func(T) bool) []T {
	// refactored this to follow DRY.
	// here we pass the list and predicate to Keep() which is a closure for fn
	// and returns opposite of what fn returns.
	return Keep(lst, func(x T) bool { return !fn(x) })
}
