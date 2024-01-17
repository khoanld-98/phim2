package models

import "gorm.io/gorm"

type Chap struct {
	gorm.Model
	MovieId int `gorm:"index"`
	Name    string
	Content string
}
