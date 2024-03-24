// Package secret contains tools derive secret messages.
package secret

var secretCode = [...]string{
	"wink",
	"double blink",
	"close your eyes",
	"jump",
}

// Handshake convert the provided code into a secret message.
func Handshake(code uint) []string {
	message := make([]string, 0)
	for idx, text := range secretCode {
		pos := uint(1 << idx)
		if pos&code == pos {
			message = append(message, text)
		}
	}
	if code&(1<<4) == 1<<4 {
		reverse(message)
	}
	return message
}

// reverse reverses the provided slice.
func reverse(message []string) {
	msgLen := len(message)
	for i, j := 0, msgLen-1; i < msgLen/2; i, j = i+1, j-1 {
		message[i], message[j] = message[j], message[i]
	}
}
