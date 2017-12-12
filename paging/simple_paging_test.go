package paging

import (
	"testing"
	"fmt"
)

func TestGetSimplePagingInfo(t *testing.T) {
	fmt.Print(GetSimplePagingInfo(1, 20))
}
