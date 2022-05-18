package decimal_to_binary

import "strconv"

// decimal -> base 10
// binary -> base 2
func decimalToBinary(num int) string {
	binaryStr := ""

	for num > 0 {
		remainder := num % 2
		num = num / 2

		binaryStr = strconv.Itoa(remainder) + binaryStr
	}

	return binaryStr
}

// decimal -> base 10
// binary -> base 2
func decimalToBinary2(num int) int {
	binary := 0
	place := 1

	for num > 0 {
		remainder := num % 2
		num = num / 2

		binary = (remainder * place) + binary

		place *= 10
	}

	return binary
}
