package golib

import (
	"fmt"
	"net/http"
	"strings"
)

// 上传图片文件的时候可以用来判断图片的类型
func GuessImgType(b *[]byte) (string, error) {
	contentType := http.DetectContentType(*b)
	list := strings.Split(contentType, "/")
	if len(list) > 1 {
		fileType := list[len(list)-1]
		return fileType, nil
	} else {
		err := fmt.Errorf("不支持的文件格式 `%s`", contentType)
		return "", err
	}
}
