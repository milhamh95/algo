package palindrome2

import "strings"

func Palindrome(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)
	sLen := len(s)

	for i := 0; i < sLen/2; i++ {
		if s[i] == s[sLen-i-1] {
			continue
		}

		return false
	}

	return true
}
