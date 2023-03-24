package password

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/869413421/chatgpt-web/pkg/logger"
)

// Hash 进行加密
func Hash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		logger.Danger(err, "hash password error")
	}

	return string(bytes)
}

//CheckHash 检查密码和hash是否匹配
func CheckHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println(err)
	return err == nil
}

// IsHashed 检查密码和hash是否已经加密
func IsHashed(str string) bool {
	return len(str) == 60
}
