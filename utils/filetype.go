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
type File struct{}

type FileInfer interface {
	GetFileInfo(string) (*FileSet, error)
	FileSize(filepath string) int64  //获取文件大小
	FileName(filepath string) string //获取文件名称
}

type FileSet struct {
	Name    string
	Size    int64
	Dir     bool
	Mode    interface{}
	ModTime interface{}
}

func (f *File) GetFileInfo(filepath string) (*FileSet, error) {
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

func (f *File) FileSize(filepath string) int64 {
	file, _ := f.GetFileInfo(filepath)
	return file.Size
}

func (f *File) FileName(filepath string) string {
	file, _ := f.GetFileInfo(filepath)
	return file.Name
}
