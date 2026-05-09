func evalRPN(tokens []string) int {
	stack := []int{}
	
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		
		if err == nil {
			stack = append(stack, val)
		} else {
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			
			result := 0
			switch token {
				case "+": result = a + b
				case "-": result = a - b
				case "*": result = a * b
				case "/": result = a / b
			}
			stack = append(stack, result)
		}
	}
	
	return stack[0]
}