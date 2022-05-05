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
