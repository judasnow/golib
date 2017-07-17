package golib

import (
	"testing"
)


func TestIsValidEmail(t *testing.T) {
	match, _ := IsValidEmail("a@a.a")
	if match {
		t.Error("a@a.a 被判定为有效")
	}

	match2, _ := IsValidEmail("test@gmail.com")
	if !match2 {
		t.Error("test@gmail.com 被判定为无效")
	}
}