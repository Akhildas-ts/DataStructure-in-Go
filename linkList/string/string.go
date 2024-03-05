package String

import (
	"strings"
)

func StringCount(val string, n int) string {
	result := ""

	for _, char := range val {
		if char >= 'a' && char <= 'z' {
			newNode := 'a' + (char-'a'+rune(n))%26
			result += string(newNode)
		} else if char >= 'A' && char <= 'Z' {
			newChar := 'A' + (char-'A'+rune(n))%26
			result += string(newChar)
		} else {
			result += string(char)
		}
	}

	return result
}

func LengthLastWord(s string) int {
	trimmed := strings.TrimSpace(s)

	// Split the string by spaces
	words := strings.Split(trimmed, " ")

	// Get the last word and return its length
	lastWord := words[len(words)-1]
	return len(lastWord)
}
