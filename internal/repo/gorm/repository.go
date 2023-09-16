package gorm

import (
	"nonelandBackendInterview/configs"
	"nonelandBackendInterview/internal/entity"

	"gorm.io/gorm"
)

type repository struct {
	db     *gorm.DB
	config *configs.Config
}

func NewRepository(db *gorm.DB, config *configs.Config) entity.Repository {
	return &repository{
		db:     db,
		config: config,
	}
}
