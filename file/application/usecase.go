package application

import (
	"github.com/mukezhz/hexagonal-architecture/file/domain"
	"mime/multipart"
)

type FileUseCase struct {
	FileRepo domain.FileRepository
}

func NewFileUseCase(fileRepo domain.FileRepository) *FileUseCase {
	return &FileUseCase{
		FileRepo: fileRepo,
	}
}

func (f *FileUseCase) Upload(file *multipart.FileHeader, dst string) error {
	// business logic to check weather file should be save or not
	if err := f.FileRepo.Save(file, dst); err != nil {
		return err
	}
	return nil
}
