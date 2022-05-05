package palindrome_number

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	newX := x
	res := 0
	for newX > 0 {
		remain := newX % 10
		newX = newX / 10

		res = (res * 10) + remain
	}

	return res == x
}

func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}

	xStr := strconv.Itoa(x)
	r := len(xStr)

	for i := 0; i < r/2; i++ {
		if xStr[i] == xStr[r-1-i] {
			continue
		}

		return false
	}

	return true
}
