package infrastructure

import (
	"github.com/mukezhz/hexagonal-architecture/file/domain"
	"gorm.io/gorm"
)

type MysqlRepository struct {
	Client *gorm.DB
}

func NewMysqlRepository(client *gorm.DB) *MysqlRepository {
	return &MysqlRepository{Client: client}
}

func (db *MysqlRepository) Create(data domain.FileMetadata) error {
	if err := db.Client.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (db *MysqlRepository) CreateExcel(data []domain.RouteStore) error {
	if err := db.Client.Create(data).Error; err != nil {
		return err
	}
	return nil
}
