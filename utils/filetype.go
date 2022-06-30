package utils

import (
	"net/http"
	"os"
)

func GetFileContentType(out *os.File) (string,error) {

	// 只需要前 512 个字节就可以了
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "文件类型错误", err
	}

	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func Stepbystep(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return Stepbystep(n-1) + Stepbystep(n-2)
}
