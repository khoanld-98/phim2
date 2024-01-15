package auth

import (
	"web/service/auth"
	"web/store"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	db := store.ConnectDB()
	var registerClass auth.RegisterClass

	if !registerClass.Register(ctx, db) {
		return
	}

	ctx.JSON(200, gin.H{})
}
