package decimal_to_binary

import "strconv"

// decimal -> base 10
// decimal -> base 2
func decimalToBinary(num int) string {
	binaryStr := ""

	for num > 0 {
		remainder := num % 2
		num = num / 2

		binaryStr = strconv.Itoa(remainder) + binaryStr
	}

	return binaryStr
}
