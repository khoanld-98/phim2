package models

import (
	"strconv"
	"web/store"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	err := godotenv.Load(".env")
	db = store.ConnectDB()
	if err != nil {
		panic("file khong ton tai")
	}

	db = store.ConnectDB()
}

func Paginate(ctx *gin.Context, limit int, customPage int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(ctx.Query("page"))

		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(ctx.Query("limit"))

		if pageSize <= 0 {
			pageSize = 30
		}

		if limit != 0 {
			pageSize = limit
		}

		if customPage > 0 {
			page = customPage
		}

		offset := (page - 1) * pageSize

		return db.Offset(offset).Limit(pageSize)
	}
}
