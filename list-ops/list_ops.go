// Package listops contains solution for the List Ops Exercise on Exercism.
package listops

// IntList is an abstraction of a list of integers which we can define methods on.
type IntList []int

// Foldl takes a function, a list, and initial accumulator, and folds each item into the
// accumulator from the left.
func (s IntList) Foldl(fn func(int, int) int, initial int) (cal int) {
	sLen := s.Length()

	if sLen == 0 {
		return initial // if there are no items in the list then return initial val.
	}

	cal = initial // reduce the initial val and the first value from list.

	for i := 0; i < sLen; i++ {
		cal = fn(cal, s[i]) // reduce subsequent values with cal.
	}

	return // named returns used to allow func to return a value if panic occurs.
}

// Foldr takes a function, a list, and initial accumulator, and folds each item into the
// accumulator from the right.
func (s IntList) Foldr(fn func(int, int) int, initial int) (cal int) {
	sLen := s.Length()

	if sLen == 0 {
		return initial // if there are no items in list then return initial val.
	}

	// as we are going from the right, reduce the initial val with the last value from the list.
	cal = initial

	for i := sLen - 1; i >= 0; i-- {
		cal = fn(s[i], cal) // reduce all preceding values with cal.
	}
	return // named returns used to allow func to return a value if panic occurs.
}

// Filter method takes a predicate and a list, return the list of all items for which predicate
// is True.
func (s IntList) Filter(fn func(int) bool) IntList {
	// make a new list, with the same size as our current list.
	// I forbid myself from using append() for this exercise.
	newS := make(IntList, s.Length())
	sLen := 0

	for _, num := range s { // range over all items in s.
		// if predicate doesn't return True for current item then move to next item.
		if !fn(num) {
			continue
		}
		newS[sLen] = num
		sLen++ // update len after adding the item.
	}
	// only return the list of size len, this allows us to prevent returning a bloated list with
	// zeroed values.
	return newS[:sLen]
}

// Length method returns number of items in s.
func (s IntList) Length() int {
	return len(s) // couldn't find a way around using len() func, complier magic at its best.
}

// Map takes a function and a list and returns the list of the results of applying function(item)
// on all items.
func (s IntList) Map(fn func(int) int) IntList {
	for idx, num := range s {
		s[idx] = fn(num) // just update the current item with the return value of the fn func.
	}
	return s
}

// Reverse takes a list and returns a list with all the original items, but in reversed order.
func (s IntList) Reverse() IntList {
	sLen := s.Length()
	newS := make(IntList, sLen)

	for i := sLen; i > 0; i-- {
		newS[sLen-i] = s[i-1] // adding last item of s to the first of newS and so on...
	}
	return newS
}

// Append method takes a list and adds all items in that list to the end of the s.
func (s IntList) Append(lst IntList) IntList {
	sLen := s.Length()

	// create a new list whose size is the sum of both list's size.
	new := make(IntList, sLen+len(lst))

	for idx, num := range s { // we add all elements from first list to the new list.
		new[idx] = num
	}

	for idx, num := range lst { // we add all elements from second list to the new list.
		new[sLen+idx] = num
	}
	return new
}

// Concat method takes a series of lists and combines all items in all lists into one flattened
// list. the returned list includes all elements from s.
func (s IntList) Concat(lists []IntList) IntList {
	for _, list := range lists { // we range over all lists
		// we use append method to add the current list to the end of s list and then update s.
		s = s.Append(list)
	}
	return s
}
