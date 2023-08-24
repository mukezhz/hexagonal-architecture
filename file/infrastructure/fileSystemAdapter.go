package infrastructure

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type FileSystemAdapter struct {
}

func NewFileSystemAdapter() *FileSystemAdapter {
	return &FileSystemAdapter{}
}

func (repo *FileSystemAdapter) Save(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {

		}
	}(src)

	if err = os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {

		}
	}(out)

	_, err = io.Copy(out, src)
	return err
}
