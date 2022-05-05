package palindrome

func Palindrome(s string) bool {
	sLen := len(s)

	for i := 0; i < sLen/2; i++ {
		if s[i] == s[sLen-i-1] {
			continue
		}

		return false
	}

	return true
}
