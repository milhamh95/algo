package reverseint

import (
	"strconv"
)

func ReverseInt(n int) int {
	tmpN := n
	if n < 0 {
		tmpN = tmpN * -1
	}
	stringTmpN := ""

	for tmpN > 0 {
		remain := tmpN % 10
		tmpN = tmpN / 10

		stringRemain := strconv.Itoa(remain)
		stringTmpN = stringTmpN + stringRemain
	}

	newN, _ := strconv.Atoi(stringTmpN)

	if n < 0 {
		newN = newN * -1
	}

	return newN
}

func ReverseInt2(n int) int {
	tmpN := n
	if n < 0 {
		tmpN = tmpN * -1
	}

	reversedN := 0

	for tmpN > 0 {
		remainder := tmpN % 10
		tmpN = tmpN / 10

		reversedN = (reversedN * 10) + remainder
	}

	if n < 0 {
		reversedN = reversedN * -1
	}

	return reversedN
}
