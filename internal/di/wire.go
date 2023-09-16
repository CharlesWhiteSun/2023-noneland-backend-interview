//go:build wireinject
// +build wireinject

package di

import (
	"nonelandBackendInterview/configs"
	"nonelandBackendInterview/internal/db"
	"nonelandBackendInterview/internal/entity"
	repo "nonelandBackendInterview/internal/repo/gorm"
	"sync"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var cg *configs.Config
var configOnce sync.Once

func NewConfig() *configs.Config {
	configOnce.Do(func() {
		cg = configs.NewConfig()
	})

	return cg
}

var database *gorm.DB
var dbOnce sync.Once

func NewDB() *gorm.DB {
	dbOnce.Do(func() {
		database = db.NewDb()
	})

	return database
}

func NewRepo() (entity.Repository, error) {
	panic(wire.Build(repo.NewRepository, NewDB, NewConfig))
}
