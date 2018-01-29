package golib

func prependIntSlice(s []int, e int) []int {
	return append([]int{e}, s...)
}
