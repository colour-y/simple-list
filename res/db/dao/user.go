package dao

import (
	"context"
	"gorm.io/gorm"
	"simplelist/res/db/model"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewDBClient(ctx)}
}

func (dao *UserDao) FindUserByUserName(userName string) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("user_name", userName).First(&user).Error
	return
}

func (dao *UserDao) FindUserByUserId(id int) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=", id).First(&user).Error

	return
}

func (dao *UserDao) CreateUser(user *model.User) (err error) {
	err = dao.DB.Model(&model.User{}).Create(user).Error
	return
}
