package fibonacci

// imperative
// n = 2, 1
// n = 3 => 2
//[0,1,1,2,3]
func generateFibonacci1(n int) int {
	res := []int{0, 1}

	for i := 1; i < n; i++ {
		total := res[i] + res[i-1]
		res = append(res, total)
	}

	return res[n]
}

func generateFibonacci2(n int) int {
	// jika n < 2, misal angka 0 & 1
	// maka kembalikan saja angka 0 & 1 (angka dasar untuk membangun fibonacci)
	// base case
	if n < 2 {
		return n
	}

	// recursive case
	return generateFibonacci2(n-1) + generateFibonacci2(n-2)
}
