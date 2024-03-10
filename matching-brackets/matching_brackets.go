package brackets

var links = map[rune]rune{
	'}': '{',
	']': '[',
	')': '(',
}

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
