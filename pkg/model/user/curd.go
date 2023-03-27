package user

import (
	"github.com/869413421/chatgpt-web/pkg/model"
)

// GetByName 根据名称获取用户
func GetByName(name string) (user *User, err error) {
	user = &User{}
	err = model.DB.Where("name = ?", name).First(user).Error
	return
}

// CreateUser 创建用户
func CreateUser(name, password string) (user *User, err error) {
	user = &User{}
	user.Name = name
	user.Password = password
	result := model.DB.Create(user)
	err = result.Error
	return
}
