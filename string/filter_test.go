package string

import (
	"testing"
)

func TestEmojiFilter(t *testing.T) {
	if EmojiFilter("xxoo😀") != "xxoo" {
		t.Error("fail")
	}
}
