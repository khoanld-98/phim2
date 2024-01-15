package models

import (
	"time"

	"gorm.io/gorm"
)

type PersonalAccessToken struct {
	gorm.Model
	UserId      int `gorm:"index"`
	AccessToken string
	LoginAt     time.Time
}
