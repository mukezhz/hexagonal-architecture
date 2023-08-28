package application

import (
	"github.com/mukezhz/hexagonal-architecture/file/domain"
	"mime/multipart"
)

type FileUseCase struct {
	FileHandler domain.FileOutgoingPort
	Repository  domain.FileRepository
}

func NewFileUseCase(fileRepo domain.FileOutgoingPort, repository domain.FileRepository) *FileUseCase {
	return &FileUseCase{FileHandler: fileRepo, Repository: repository}
}

func (f *FileUseCase) Upload(file *multipart.FileHeader, dst string) (string, error) {
	// business logic to check weather file should be save or not
	_, err := f.FileHandler.Save(file, dst)
	if err != nil {
		return "", err
	}
	filePath, err := f.FileHandler.GetSignedURL(file, dst, nil)
	if err != nil {
		return filePath, err
	}
	return filePath, nil
}

func (f *FileUseCase) Save(file domain.FileMetadata) error {
	if err := f.Repository.Create(file); err != nil {
		return err
	}
	return nil
}
