package wordy

import (
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`^(What is) (-?\d+)( (plus|minus|multiplied by|divided by){1} (-?\d+))*( )?(\?$)`)
var r = strings.NewReplacer("What is ", "", " by", "", "?", "")

func Answer(question string) (int, bool) {
	if !re.MatchString(question) {
		return 0, false
	}
	sanitized := strings.Split(r.Replace(question), " ")
	cal, err := strconv.Atoi(sanitized[0])
	if err != nil {
		return 0, false
	}
	for i := 1; i < len(sanitized)-1; i++ {
		val, err := strconv.Atoi(sanitized[i+1])
		if err != nil {
			continue
		}
		switch sanitized[i] {
		case "plus":
			cal += val
		case "minus":
			cal -= val
		case "multiplied":
			cal *= val
		case "divided":
			cal /= val
		}
	}
	return cal, true

}
