// Package brackets contains tools to parse and match brackets
package brackets

var links = map[rune]rune{
	'}': '{',
	']': '[',
	')': '(',
}

// Bracket returns true or false based on if all the brackets in the input string are closed.
func Bracket(input string) bool {
	memory := make([]rune, 0)
	for _, char := range input {
		switch char {
		case '{', '(', '[':
			memory = append(memory, char)
		case '}', ')', ']':
			last := len(memory) - 1
			if last == -1 || memory[last] != links[char] {
				return false
			}
			memory = memory[:last]
		}
	}
	return len(memory) == 0
}
