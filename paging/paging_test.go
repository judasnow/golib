package paging

import (
	"fmt"
	"testing"
)

func TestPaging(t *testing.T) {
	p1 := Paging(0, 0, 0)
	if p1.Count != 0 || p1.Limit != 0 || p1.Offset != 0 {
		t.Error("p1")
	}

	// 最后一页不够满页的情况
	p2 := Paging(11, 2, 1)
	if p2.Count != 6 || p2.Limit != 2 || p2.Offset != 0 {
		t.Error("p2")
	}
	p22 := Paging(11, 2, 2)
	if p22.Count != 6 || p22.Limit != 2 || p22.Offset != 2 {
		t.Error("p22")
	}
	p26 := Paging(11, 2, 6)
	if p26.Count != 6 || p26.Limit != 2 || p26.Offset != 10 {
		t.Error("p226")
	}

	// 最后一页满页的情况
	p3 := Paging(12, 2, 1)
	if p3.Count != 6 || p3.Limit != 2 || p3.Offset != 0 {
		t.Error("p3")
	}
}

func TestPaging2(t *testing.T) {
	// 测试分页循环
	var tasksCount = 4
	var perPage = 1
	var crtPage = 1

	for {
		pageInfo := Paging(tasksCount, perPage, crtPage)

		if crtPage > pageInfo.Count || crtPage <= 0 {
			return
		} else {
			// loop content
			crtPage = crtPage + 1
		}
	}
}

func TestPagingLoop(t *testing.T) {
	PagingLoop(4, 1, 1, func(limit int, offset int) bool {
		fmt.Printf("limit: %d, offset: %d\n", limit, offset)
		return true
	})
}

// 对 slice 进行分页显示
func TestPagingLoop2(t *testing.T) {
	x := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	SlicePagingLoop(x, 3, 1, func(subSlice []interface{}) bool {
		fmt.Println(subSlice)
		return true
	})
}

// 第一次返回 false 停止循环
func TestPagingLoopReutrnFalse(t *testing.T) {
	x := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	SlicePagingLoop(x, 3, 1, func(subSlice []interface{}) bool {
		fmt.Println(subSlice)
		return false
	})
}

func TestPagingLoopEmptySlice(t *testing.T) {
	x := []interface{}{}
	SlicePagingLoop(x, 3, 1, func(subSlice []interface{}) bool {
		fmt.Println(subSlice)
		return true
	})
}
