package user

import (
	"errors"
	"net/http"
	"web/helper"
	"web/models"
	"web/store"
	"web/validate"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm/clause"
)

type UpdateClass struct {
	User        models.Users
	Username    string `json:"Username" binding:"required"`
	Email       string `json:"Email" binding:"required"`
	Address     string `json:"Address" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

func (updateClass *UpdateClass) Update(ctx *gin.Context) {
	err := ctx.ShouldBindJSON(&updateClass)
	var validateErr validator.ValidationErrors
	db := store.ConnectDB()

	check := updateClass.SetUser(ctx)

	if check == false {
		return
	}

	if err != nil && errors.As(err, &validateErr) {
		messageErr := validate.ResponseErrorValidate(validateErr)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": messageErr,
		})

		return
	}

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"address", "email", "phone_number", "username"}),
	}).Create(&updateClass.User)

	ctx.Set("user", updateClass.User)

	ctx.JSON(http.StatusOK, gin.H{
		"user": updateClass.User,
	})
}

func (updateClass *UpdateClass) SetUser(ctx *gin.Context) bool {
	userMap := ctx.MustGet("user")
	user := helper.ConvertMapToUsersStruct(userMap)

	if (user == models.Users{}) {
		return false
	}

	user.Address = updateClass.Address
	user.Email = updateClass.Email
	user.PhoneNumber = updateClass.PhoneNumber

	updateClass.User = user

	return true
}
