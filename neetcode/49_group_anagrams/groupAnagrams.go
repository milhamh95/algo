package group_anagrams

import (
	"sort"
)

// m n log n
func groupAnagrams1(strs []string) [][]string {
	res := [][]string{}
	if len(strs) == 0 {
		return res
	}

	m := make(map[string][]string)

	// rune -> int32
	// represent unicode
	for _, v := range strs {
		wordRune := []rune(v)
		sort.Slice(wordRune, func(i int, j int) bool {
			return wordRune[i] < wordRune[j]
		})

		sortedWordString := string(wordRune)
		m[sortedWordString] = append(m[sortedWordString], v)
	}

	for _, v := range m {
		res = append(res, v)
	}

	return res
}

// m n log n
// n log n each word -> total word = m
func groupAnagrams2(strs []string) [][]string {
	res := [][]string{}
	if len(strs) == 0 {
		return res
	}

	// store key = sorted word in res
	// value = index of sorted word in res
	m := make(map[string]int)

	// rune -> int32
	// represent unicode
	for _, v := range strs {
		wordRune := []rune(v)
		sort.Slice(wordRune, func(i int, j int) bool {
			return wordRune[i] < wordRune[j]
		})

		sortedWordString := string(wordRune)

		idx, ok := m[sortedWordString]
		if !ok {
			m[sortedWordString] = len(res)
			res = append(res, []string{v})
			continue
		}

		res[idx] = append(res[idx], v)
	}

	return res
}
