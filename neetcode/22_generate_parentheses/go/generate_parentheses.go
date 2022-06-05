package generate_parentheses

func generateParenthesis(n int) []string {
	result := []string{}

	// https: //nanxiao.gitbooks.io/golang-101-hacks/content/posts/pass-slice-as-a-function-argument.html
	// https://www.pixelstech.net/article/1607859246-The-hidden-risk-of-passing-slice-as-function-parameter
	// https://medium.com/@alexkw/slices-as-arguments-in-golang-584a7f84f24

	generate(&result, "", 0, 0, n)
	return result
}

func generate(result *[]string, s string, open, close, n int) {
	if open == n && close == n {
		*result = append(*result, s)
		return
	}

	if open < n {
		generate(result, s+"(", open+1, close, n)
	}

	if close < open {
		generate(result, s+")", open, close+1, n)
	}
}
