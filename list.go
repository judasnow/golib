package golib

// 原地逆序一个 slice
func reverseList (items []interface{}) {
	for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]
	}
}
