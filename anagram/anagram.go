// anagram package contains tools to check anagram with given words.
package anagram

import "strings"

// stringMap takes a string and maps all its characters with there repetition.
func stringMap(str string) map[rune]int {

	strMap := make(map[rune]int, len(str))

	for _, char := range str {
		strMap[char]++
	}

	return strMap
}

// checkAnagram takes two strings and checks if they are Anagram of each other.
func checkAnagram(str1, str2 string) bool {
	str1Map := stringMap(str1)
	str2Map := stringMap(str2)

	for char, count := range str1Map {
		value, ok := str2Map[char]
		if !ok || value != count {
			return false
		}
	}
	return true

}

// Detect takes a subject and a list of candidates and returns all the candidates which are
// anagram of the subject.
func Detect(subject string, candidates []string) []string {

	lowSubject := strings.ToLower(subject)

	var filtered []string

	for _, candidate := range candidates {
		lowCandidate := strings.ToLower(candidate)

		// continue the loop if both are same words or the length of both words don't match.
		if lowSubject == lowCandidate || len(lowSubject) != len(lowCandidate) {
			continue
		}

		// check if lowercased of both words are anagram of each other.
		if checkAnagram(lowSubject, lowCandidate) {
			filtered = append(filtered, candidate)
		}

	}

	return filtered

}
