package utils

func IntRange(n int) []int {
	vals := make([]int, n+1)
	for i := range vals {
		vals[i] = i
	}
	return vals
}
