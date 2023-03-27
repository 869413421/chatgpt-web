package user

import (
	"github.com/869413421/chatgpt-web/pkg/model"
	"github.com/869413421/chatgpt-web/pkg/password"
)

type User struct {
	model.BaseModel
	Name string `gorm:"column:name;type:varchar(255);not null;unique" valid:"name"`
	//Email    string `gorm:"column:email;type:varchar(255) not null;unique" valid:"email"`
	Password string `gorm:"column:password;type:varchar(255);not null" valid:"password"`
	// gorm:"-" 使用这个注解GORM读写会忽略这个字段
	//PasswordComfirm string `gorm:"-" valid:"password_comfirm"`
}

// ComparePassword 检查密码是否匹配
func (user *User) ComparePassword(_password string) bool {
	return password.CheckHash(_password, user.Password)
}
