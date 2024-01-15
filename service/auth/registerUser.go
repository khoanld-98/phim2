package auth

import (
	"errors"
	"net/http"
	"time"
	"web/models"
	"web/validate"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterClass struct {
	UserName        string `json:"username" binding:"required,min=8"`
	Password        string `json:"password" binding:"required,min=8"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,min=8"`
}

func (r *RegisterClass) Register(ctx *gin.Context, db *gorm.DB) bool {
	err := ctx.ShouldBindJSON(r)
	var validateErr validator.ValidationErrors

	if err != nil && errors.As(err, &validateErr) {
		messageErr := validate.ResponseErrorValidate(validateErr)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": messageErr,
		})

		return err == nil
	}

	if inValid := r.Validate(db, ctx); !inValid {
		return inValid
	}

	errSave := make(chan bool)
	go r.SaveUser(db, errSave)

	return <-errSave
}

func (r *RegisterClass) SaveUser(db *gorm.DB, errSave chan<- bool) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(r.Password), 10)

	if err := db.Model(models.Users{}).
		Create(map[string]interface{}{
			"username":   r.UserName,
			"password":   hashPassword,
			"created_at": time.Now(),
			"updated_at": time.Now(),
		}).Error; err != nil {
		errSave <- false
	}

	errSave <- true
}

func (r *RegisterClass) Validate(db *gorm.DB, ctx *gin.Context) bool {
	var check int64
	db.Model(models.Users{}).Where("username = ?", r.UserName).Count(&check)

	if r.Password != r.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": map[string]string{
				"Password": "Password must same as confirm password",
			},
		})

		return false
	}

	if check > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": map[string]string{
				"UserName": "UserName has been existed",
			},
		})

		return false
	}

	return check == 0
}
