package domain

import (
	"gorm.io/gorm"
)

type FileMetadata struct {
	gorm.Model
	Filename string `gorm:"not null"`
	UUID     string `gorm:"unique;not null"`
}
