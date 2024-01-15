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

type LoginClass struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	User     models.Users
}

func (c *LoginClass) Login(ctx *gin.Context, db *gorm.DB) (models.Users, bool) {
	if validate := c.Validator(ctx); validate {
		return c.User, false
	}

	if login := c.CompareDB(db); login {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "login failed",
		})

		return c.User, false
	}

	return c.User, true
}

func (c *LoginClass) Validator(ctx *gin.Context) bool {
	err := ctx.ShouldBind(&c)

	var validateErr validator.ValidationErrors
	if err != nil && errors.As(err, &validateErr) {
		messageErr := validate.ResponseErrorValidate(validateErr)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": messageErr,
		})

		return true
	}

	return false
}

func (c *LoginClass) CompareDB(db *gorm.DB) bool {
	if err := db.Model(&models.Users{}).
		Where("username =?", c.UserName).
		First(&c.User).
		Error; err != nil {

		return true
	}

	err := bcrypt.CompareHashAndPassword([]byte(c.User.Password), []byte(c.Password))

	return err != nil
}

func (c *LoginClass) StorePersonalToken(token string, db *gorm.DB) bool {
	db.Unscoped().Where("user_id = ?", c.User.ID).Delete(&models.PersonalAccessToken{})

	personAccessToken := models.PersonalAccessToken{
		UserId:      int(c.User.ID),
		AccessToken: token,
		LoginAt:     time.Now(),
	}

	if err := db.Create(&personAccessToken).Error; err != nil {
		return false
	}

	return true
}
