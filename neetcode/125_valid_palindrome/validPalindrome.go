package validpalindrome

import "strings"

func isPalindrome(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)
	sLen := len(s)

	for i := 0; i < sLen/2; i++ {
		if s[i] == s[sLen-1-i] {
			continue
		}

		return false
	}

	return true
}
