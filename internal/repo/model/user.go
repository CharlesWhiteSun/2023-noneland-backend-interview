package model

import "nonelandBackendInterview/internal/entity"

type User struct {
	Name string `gorm:"type:varchar(255)"`
}

func UserModelToEntity(input *User) *entity.User {
	return &entity.User{
		Name: input.Name,
	}
}

func UserEntityToModel(input *entity.User) *User {
	return &User{
		Name: input.Name,
	}
}
