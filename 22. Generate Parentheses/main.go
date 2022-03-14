package main

var result []string

func generateParenthesis(n int) []string {
	result = []string{}

	dfs("", n, n)

	return result
}

func dfs(s string, l int, r int) {
	if l == 0 && r == 0 {
		result = append(result, s)
	}

	if r < l {
		return
	}

	if l > 0 {
		dfs(s + "(", l-1, r)
	}

	if r > 0 {
		dfs(s + ")", l, r-1)
	}

	return
}