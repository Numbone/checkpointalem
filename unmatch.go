package piscine

func Unmatch(arr []int) int {
	c := 0
	for _, l := range arr {
		for _, k := range arr {
			if l == k {
				c++
			}
		}
		if c%2 != 0 {
			return l
		}
	}
	return -1
}
