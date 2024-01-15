package auth

import (
	"fmt"
	"net/http"
	"web/helper"
	"web/models"
	"web/store"

	"github.com/gin-gonic/gin"
)

func Logout(ctx *gin.Context) {
	userMap := ctx.MustGet("user")
	fmt.Printf("%+v", userMap)

	if userMap != nil {
		user := helper.ConvertMapToUsersStruct(userMap)
		fmt.Printf("%+v", user)
		db := store.ConnectDB()
		db.Where("user_id = ?", user.ID).Delete(&models.PersonalAccessToken{})

		ctx.JSON(http.StatusOK, gin.H{})
		return
	}

	ctx.JSON(http.StatusNotAcceptable, gin.H{})
}
