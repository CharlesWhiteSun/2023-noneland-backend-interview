package gorm

import (
	"nonelandBackendInterview/internal/entity"
	"nonelandBackendInterview/internal/repo/model"
)

func (repo *repository) GetUsers() (users []entity.User, err error) {
	datas := []*model.User{}
	err = repo.db.Model(&model.User{}).Find(&datas).Error
	if err != nil {
		return nil, err
	}

	users = make([]entity.User, len(datas))
	for i, v := range datas {
		item := model.UserModelToEntity(v)
		if item != nil {
			users[i] = *item
		}
	}

	return nil, err
}
