package tests

func isOdd(i int) bool {
	return i%2 != 0
}

func startsWith(letter string) func(string) bool {
	return func(s string) bool {
		if len(s) == 0 {
			return false
		}
		return string(s[0]) == letter
	}
}
