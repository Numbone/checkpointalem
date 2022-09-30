package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	c := 0
	k := len(a) - 1
	for i := range a[:k] {
		if f(a[i], a[i+1]) <= 0 {
			c++
		}
	}
	if c == k || c == 0 {
		return true
	}
	return false
}
