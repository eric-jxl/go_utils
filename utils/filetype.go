package utils

import (
	"net/http"
	"os"
)

func GetFileContentType(out *os.File) (string, error) {

	// 只需要前 512 个字节就可以了
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "文件类型错误", err
	}

	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

type FileSet struct {
	Name    string
	Size    int64
	Dir     bool
	Mode    interface{}
	ModTime interface{}
}

func GetFileInfo(filepath string) (*FileSet, error) {
	file, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}
	return &FileSet{
		Name:    file.Name(),
		Size:    file.Size(),
		Dir:     file.IsDir(),
		Mode:    file.Mode(),
		ModTime: file.ModTime(),
	}, nil
}
