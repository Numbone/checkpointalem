package piscine

func Compact(ptr *[]string) int {
	a := *ptr
	var arr []string

	for _, letter := range a {
		if letter != "" {
			arr = append(arr, letter)
		}
	}

	*ptr = arr
	return len(arr)
}
