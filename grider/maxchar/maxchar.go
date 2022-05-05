package maxchar

func MaxChar(s string) string {
	mapChar := map[string]int{}

	for _, v := range s {
		counter, ok := mapChar[string(v)]
		if ok {
			mapChar[string(v)] = counter + 1
			continue
		}

		mapChar[string(v)] = 1
	}

	totalChar := 0
	maxChar := ""

	for k, v := range mapChar {
		if v >= totalChar {
			totalChar = v
			maxChar = k
		}
	}

	return maxChar
}
