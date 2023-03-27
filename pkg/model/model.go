package model

import (
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	gloger "gorm.io/gorm/logger"

	"github.com/869413421/chatgpt-web/pkg/logger"
	"github.com/869413421/chatgpt-web/pkg/types"
)

// BaseModel 主模型
type BaseModel struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement;not null"`
	CreatedAt time.Time `gorm:"column:created_at;index"`
	UpdatedAt time.Time `gorm:"column:updated_at;index"`
}

// GetStringID 获取主键字符串
func (model BaseModel) GetStringID() string {
	return types.UInt64ToString(model.ID)
}

// CreatedAtDate 获取创建时间
func (model BaseModel) CreatedAtDate() string {
	return model.CreatedAt.Format("2006-01-02")
}

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := "chat.db"
	var err error
	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: gloger.Default.LogMode(gloger.Info),
	})
	if err != nil {
		logger.Danger("open sqlite error:", err)
	}
	return DB
}
