package print_star_based_on_num

import (
	"strconv"
)

// 1 * 3 4 * 6 * 8 9 *
// 11 * 13 14 * 16 * 18 19 *
func print(num int) string {
	str := ""

	for i := 1; i <= num; i++ {
		space := " "
		if i == num {
			space = ""
		}

		if i%5 == 2 ||
			i%5 == 0 {

			str = str + "*" + space
			continue
		}

		str = str + strconv.Itoa(i) + space
	}

	return str
}
