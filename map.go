package piscine

func Map(f func(int) bool, a []int) []bool {
	boolean := []bool{}
	for _, s := range a {
		boolean = append(boolean, f(s))
	}
	return boolean
}
