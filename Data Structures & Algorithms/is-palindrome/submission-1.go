func isAlphanumeric(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b >= '0' && b <= '9'
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	s = strings.ToLower(s)

	l := 0
	r := len(s) - 1

	for r >= l {
		for !isAlphanumeric(s[l]) && l < len(s)-1 {
			l++
		}

		for !isAlphanumeric(s[r]) && r > 0 {
			r--
		}

		if r >= l && s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}