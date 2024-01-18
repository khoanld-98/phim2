package models

import "gorm.io/gorm"

type Story struct {
	gorm.Model
	Name        string
	Description string
	Auth        string
	source      string
	Status      int
	AllowCrawl  bool
}
