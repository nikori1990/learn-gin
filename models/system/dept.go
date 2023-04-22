package system

import (
	"gorm.io/gorm"
	"time"
)

type Dept struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	ParentId  uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Dept) TableName() string {
	return "dept"
}
