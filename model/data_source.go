package model

import "gorm.io/gorm"

type DataSource struct {
	gorm.Model
	Name        string `gorm:"size:10;unique;index" binding:"required"`
	Description string `gorm:"size:50"`
	CreateUser  string `gorm:"size:10;not null"`
	Data        string `gorm:"not null" binding:"required"`
	GroupID     uint
}
