package is_prime

func checkIsPrime1(n int) bool {
	isPrime := true
	for i := 2; i < n; i++ {
		if n%i == 0 {
			isPrime = false
		}
	}

	return isPrime
}

func checkIsPrime1(n int) bool {
	isPrime := true
	for i := 2; i < n; i++ {
		if n%i == 0 {
			isPrime = false
		}
	}

	return isPrime
}
