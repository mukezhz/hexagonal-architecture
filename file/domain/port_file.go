package domain

import (
	"mime/multipart"
	"time"
)

/* ========== incoming ports ============= */
type FileIncomingPort interface {
	Upload(file *multipart.FileHeader, dst string) (string, error)
	Save(file FileMetadata) error
}

/* ========== outgoing ports ============= */
type FileOutgoingPort interface {
	Save(file *multipart.FileHeader, dst string) (string, error)
	SavePublicly(file *multipart.FileHeader, dst string) (string, error)
	GetSignedURL(file *multipart.FileHeader, dst string, expires *time.Time) (string, error)
}

type FileRepository interface {
	Create(file FileMetadata) error
}
