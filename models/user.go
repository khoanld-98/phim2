package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username    string `form:"username" json:"Username"`
	Password    string `form:"password" json:"-"`
	Email       string `form:"email" json:"Email"`
	Address     string `form:"address" json:"Address"`
	PhoneNumber string `form:"phone_number" json:"phone_number"`
}
