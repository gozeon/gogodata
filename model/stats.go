package model

import "gorm.io/gorm"

type Stats struct {
	gorm.Model
	RequestMethod string
	RequestURI    string
	RequestIP     string
	ResponseCode  int
	ResponseTime  string
	GroupID       uint `gorm:"index"`
	DataSourceID  uint `gorm:"index"`
}
