package validparentheses

func isValid(s string) bool {
	stack := make([]rune, len(s))
	top := 0

	parentheses := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, v := range s {
		switch v {
		case '(', '[', '{':
			stack[top] = parentheses[v]
			top++
		case ')', ']', '}':
			if top > 0 && stack[top-1] == v {
				top--
				continue
			}

			return false
		}
	}

	return top == 0
}
