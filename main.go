package main

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}

	return generate("", n, n)
}

func generate(s string, left, right int) []string {

	// closes () if left open, and returns []string{s}
	// the recursion will end here
	if left == 0 {
		for i := 0; i < right; i++ {
			s += ")"
		}

		return []string{s}
	}

	result := []string{}

	// open parentheses first
	if left > 0 {
		l := generate(s+"(", left-1, right)
		result = append(result, l...)
	}

	// close parentheses if right remaining > left remaining
	if right > left {
		r := generate(s+")", left, right-1)
		result = append(result, r...)
	}

	return result
}
