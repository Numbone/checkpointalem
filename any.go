package piscine

func Any(f func(string) bool, a []string) bool {
	for _, s := range a {
		if f(s) == true {
			return true
		}
	}
	return false
}

func IsNumeric(s string) bool {
	h := []rune(s)
	for i := 0; i <= len(h)-1; i++ {
		if (h[i] >= 0) && (h[i] <= 47) || (h[i] >= 58) && (h[i] <= 127) {
			return false
		}
	}
	return true
}
