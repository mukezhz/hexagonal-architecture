package domain

import "mime/multipart"

type FilePort interface {
	Upload(file *multipart.FileHeader, dst string) error
}

type FileRepository interface {
	Save(file *multipart.FileHeader, dst string) error
}
