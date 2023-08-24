package infrastructure

import (
	"errors"
	"fmt"
	"github.com/mukezhz/hexagonal-architecture/file/domain"
	"github.com/xuri/excelize/v2"
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

func (fsa *FileSystemAdapter) Save(file *multipart.FileHeader, dst string) error {
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

func (fsa *FileSystemAdapter) GetAll(filePath string) ([]domain.RouteStore, error) {
	routeStores := make([]domain.RouteStore, 0)
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	sheets := f.GetSheetList()
	for _, sheet := range sheets {
		rows, err := f.GetRows(sheet)

		if err != nil {
			return nil, err
		}

		const headerRow = 1
		rows = rows[headerRow:]

		for _, row := range rows {
			l := len(row)
			if l == 0 {
				// ignore if the line is empty
				continue
			} else if l == 2 {
				if len(row[0]) == 0 {
					fmt.Println("no route name")
					continue
				}
				routeStores = append(routeStores, domain.RouteStore{
					RouteName: row[0],
					Store:     row[1],
				})
			} else if l == 1 {
				return nil, errors.New("store code is empty")
			}
		}
	}
	return routeStores, nil
}
