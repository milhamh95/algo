package validanagram

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapS := map[string]int{}
	for _, v := range s {
		mapS[string(v)]++
	}

	for _, v := range t {
		mapS[string(v)]--
		counter, ok := mapS[string(v)]
		if !ok {
			return false
		}

		if counter < 0 {
			return false
		}
	}

	return true
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapS := map[string]int{}
	mapT := map[string]int{}

	for _, v := range s {
		counter, ok := mapS[string(v)]
		if !ok {
			mapS[string(v)] = 1
			continue
		}

		mapS[string(v)] = counter + 1
	}

	for _, v := range t {
		counter, ok := mapT[string(v)]
		if !ok {
			mapT[string(v)] = 1
			continue
		}

		mapT[string(v)] = counter + 1
	}

	if len(mapS) != len(mapT) {
		return false
	}

	for k, v := range mapS {
		counter, ok := mapT[k]
		if !ok {
			return false
		}

		if v != counter {
			return false
		}
	}

	return true
}
