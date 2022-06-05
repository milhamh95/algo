package factorial

func genFactorial1(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	return res
}

func genFactorial2(n int) int {
	if n < 2 {
		return n
	}

	return n * genFactorial2(n-1)
}
