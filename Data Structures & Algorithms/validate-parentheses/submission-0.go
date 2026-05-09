func isValid(s string) bool {
	stack := make([]rune, 0, len(s))
	parentheses := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	
	for _, p := range s {
		if p == '(' || p == '{' || p == '[' {
			stack = append(stack, p)
			continue
		}
		
		if len(stack) == 0 || stack[len(stack)-1] != parentheses[p] {
			return false
		}
		
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}