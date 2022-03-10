package model

import (
	"gorm.io/datatypes"
	"time"
)

type MiddleLib struct {
	Id    		int64 `gorm:"primaryKey;column:id;autoIncrement" `
	Sn        	string
	Source      string
	Content    	datatypes.JSON  `gorm:"column:content" `
	IsQRcode    bool  `gorm:"column:is_qrcode" `
	CodeUrl     string `gorm:"column:code_url" `
	CreatedAt time.Time `gorm:"column:created_at;index;"`
	UpdatedAt time.Time  `gorm:"column:updated_at;index;"`    // 在创建时该字段值为零值或者在更新时，使用当前时间戳秒数填充
}
