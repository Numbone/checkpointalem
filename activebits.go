package piscine

func ActiveBits(n int) int {
	count := 0
	var arr []int
	for n != 0 {
		arr = append(arr, n%2)
		n /= 2
	}
	for _, digit := range arr {
		if digit == 1 {
			count++
		} else {
			continue
		}
	}
	return count
}
