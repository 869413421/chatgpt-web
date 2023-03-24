package bootstrap

import (
	"github.com/869413421/chatgpt-web/config"
	"github.com/869413421/chatgpt-web/pkg/logger"
	"github.com/869413421/chatgpt-web/pkg/model"
	"github.com/869413421/chatgpt-web/pkg/model/user"
	"gorm.io/gorm"
)

// SetupDB 启动数据库
func SetupDB() {
	//建立连接池
	db := model.ConnectDB()

	migration(db)

	insertAdmin()
}

// migration 迁移
func migration(db *gorm.DB) {
	err := db.AutoMigrate(&user.User{})
	if err != nil {
		logger.Danger("migration model error:", err)
	}
}

func insertAdmin() {
	cf := config.LoadConfig()
	if cf.AuthUser != "" {
		_, err := user.GetByName(cf.AuthUser)
		if err != nil && err != gorm.ErrRecordNotFound {
			logger.Danger("insert admin error:", err)
		}
		if err == gorm.ErrRecordNotFound {
			_, err = user.CreateUser(cf.AuthUser, cf.AuthPassword)
			if err != nil {
				logger.Danger("create admin error:", err)
			}
		}
	}
}
