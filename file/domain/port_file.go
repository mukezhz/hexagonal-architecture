package domain

import "mime/multipart"

/* ========== incoming ports ============= */
type FileIncomingPort interface {
	Upload(file *multipart.FileHeader, dst string) error
	Save(file FileMetadata) error
}

/* ========== outgoing ports ============= */
type FileOutgoingPort interface {
	Save(file *multipart.FileHeader, dst string) error
}

type FileRepository interface {
	Create(file FileMetadata) error
}
