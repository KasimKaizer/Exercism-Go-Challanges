// package stringset contains solution for Custom Set Exercise on Exercism.
package stringset

import (
	"fmt"
	"strings"
)

// Set is a custom set, it holds string values.
type Set map[string]bool

// New creates a new Set.
func New() Set {
	return make(Set)
}

// NewFromSlice takes a slice of string and returns a set with items of that slice.
func NewFromSlice(l []string) Set {
	newSet := New()
	for _, char := range l {
		newSet[char] = true
	}
	return newSet
}

// String method returns the set represented as a string.
func (s Set) String() string {
	var output strings.Builder
	idx, setLen := 0, len(s)-1
	output.WriteString("{")
	for val := range s {
		output.WriteString(fmt.Sprintf("\"%s\"", val))
		if idx == setLen {
			break
		}
		output.WriteString(", ")
		idx++
	}
	output.WriteString("}")
	return output.String()
}

// IsEmpty method returns true if the set is empty.
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

// Has method takes a string and returns true if that string is present inside the set.
func (s Set) Has(elem string) bool {
	_, ok := s[elem]
	return ok
}

// Add method takes a string and adds it to the set.
func (s Set) Add(elem string) {
	s[elem] = true

}

// Subset takes two sets and returns true if s2 is subset of s1, it returns false
// if s2 is not a subset of s1.
func Subset(s1, s2 Set) bool {
	for item := range s1 {
		if _, ok := s2[item]; !ok {
			return false
		}
	}
	return true
}

// Disjoint takes two sets and returns true if non of the elements of both sets are
// present in each other.
func Disjoint(s1, s2 Set) bool {
	for item := range s1 {
		if _, ok := s2[item]; ok {
			return false
		}
	}
	return true
}

// Equal takes two sets and returns true if they are equal, it returns false if they aren't
// equal.
func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	return Subset(s1, s2)
}

// Intersection takes two sets and returns a set of all shared elements
func Intersection(s1, s2 Set) Set {
	newSet := New()
	big, small := s1, s2
	if len(small) > len(big) {
		big, small = small, big
	}
	for item := range small {
		if _, ok := big[item]; !ok {
			continue
		}
		newSet[item] = true
	}
	return newSet
}

// Difference takes two sets and returns a set of all elements that are only in the first set.
func Difference(s1, s2 Set) Set {
	newSet := New()
	for item := range s1 {
		if _, ok := s2[item]; ok {
			continue
		}
		newSet[item] = true
	}
	return newSet
}

// Union takes two sets and returns a set of all elements in them.
func Union(s1, s2 Set) Set {
	newSet := New()
	for item, val := range s1 {
		newSet[item] = val
	}
	for item, val := range s2 {
		newSet[item] = val
	}
	return newSet
}
