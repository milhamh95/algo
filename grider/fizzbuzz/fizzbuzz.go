package fizzbuzz

import (
	"fmt"
	"strconv"
)

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else if i%15 == 0 {
			fmt.Println("fizzbuzz")
		} else {
			fmt.Println(i)
		}
	}
}

func FizzBuzz2(n int) string {
	if n%15 == 0 {
		return "fizzbuzz"
	}

	if n%3 == 0 {
		return "fizz"
	}

	if n%5 == 0 {
		return "buzz"
	}

	return strconv.Itoa(n)
}
