package string

import (
	"unicode/utf8"
)

func EmojiFilter(content string) string {
	if content == "" {
		return ""
	}

	var newContent string
	for _, value := range content {
		_, size := utf8.DecodeRuneInString(string(value))
		if size <= 3 {
			newContent += string(value)
		}
	}
	return newContent
}
