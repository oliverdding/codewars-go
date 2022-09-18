package kata

func ValidParentheses(parens string) bool {
	score := 0
	for i := 0; i < len(parens) && score >= 0; i++ {
		if parens[i] == '(' {
			score++
		} else if parens[i] == ')' {
			score--
		} else {
			return false
		}
	}
	return score == 0
}
