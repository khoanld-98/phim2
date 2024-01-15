package controller

import (
	"net/http"
	"web/service/auth"
	"web/store"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var LoginClass auth.LoginClass
	db := store.ConnectDB()

	if _, err := LoginClass.Login(ctx, db); !err {
		return
	}

	token, _ := auth.GenerateJwt(&LoginClass.User)

	if isSuccess := LoginClass.StorePersonalToken(string(token), db); !isSuccess {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to store personal token",
		})

		return
	}

	ctx.JSON(200, gin.H{
		"user":  LoginClass.User,
		"token": token,
	})
}
