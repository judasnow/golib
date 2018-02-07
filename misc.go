package golib

func prependIntSlice(s []int, e int) []int {
	return append([]int{e}, s...)
}

func Xor(a bool, b bool) bool {
	return a && !b || !a && b
}

func ThreeInputXOR(a bool, b bool, c bool) bool {
	return (a && !b && !c) || (!a && b && !c) || (!a && !b && c)
}
