package application

import (
	"github.com/mukezhz/hexagonal-architecture/file/domain"
	"mime/multipart"
)

type FileUseCase struct {
	FileSystem domain.FileOutgoingPort
	Repository domain.FileRepository
}

func NewFileUseCase(fileRepo domain.FileOutgoingPort, repository domain.FileRepository) *FileUseCase {
	return &FileUseCase{FileSystem: fileRepo, Repository: repository}
}

func (f *FileUseCase) Upload(file *multipart.FileHeader, dst string) error {
	// business logic to check weather file should be save or not
	if err := f.FileSystem.Save(file, dst); err != nil {
		return err
	}
	return nil
}

func (f *FileUseCase) Save(file domain.FileMetadata) error {
	if err := f.Repository.Create(file); err != nil {
		return err
	}
	return nil
}
