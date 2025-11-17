package models

import (
	"time"

	"gorm.io/gorm"
)

// 自定义 Model 结构体
type CustomModel struct {
	ID        uint           `gorm:"column:id;primarykey" json:"id" xlsx:"数据ID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
