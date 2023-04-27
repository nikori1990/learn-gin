package base

import (
	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt *LocalTime     `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt *LocalTime     `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
