package user

import (
	"net/http"
	"web/helper"
	"web/models"
	"web/store"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type ChangePasswordClass struct {
	OldPassword     string `json:"OldPassword" binding:"required"`
	NewPassword     string `json:"NewPassword" binding:"required"`
	ConfirmPassword string `json:"ConfirmPassword" binding:"required"`
}

func (ChangePasswordClass *ChangePasswordClass) Handle(ctx *gin.Context) {
	db := store.ConnectDB()
	userMap := ctx.MustGet("user")
	user := helper.ConvertMapToUsersStruct(userMap)

	if (user == models.Users{}) {
		return
	}

	err := ctx.ShouldBindJSON(&ChangePasswordClass)

	if validate := helper.CheckValidate(ctx, err); validate == false {
		return
	}

	db.Model(models.Users{}).Where("username = ?", user.Username).First(&user)

	if check := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(ChangePasswordClass.OldPassword)); check != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": map[string]string{
				"OldPassword": "password in correct",
			},
		})
	}

	hasPassword, _ := bcrypt.GenerateFromPassword([]byte(ChangePasswordClass.NewPassword), 10)

	db.Model(&user).Update("password", hasPassword)
}
