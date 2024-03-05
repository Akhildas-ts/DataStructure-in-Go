package String

import (
	"fmt"
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
    lastSpace := strings.LastIndex(s, " ")
    if lastSpace == -1 {
        return len(s)
    }

    
    lastWord := s[lastSpace+1:]
    return len(lastWord)
}

