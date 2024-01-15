package helper

import (
	"errors"
	"net/http"
	"web/validate"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CheckValidate(ctx *gin.Context, err error) bool {
	var validateErr validator.ValidationErrors

	if err != nil && errors.As(err, &validateErr) {
		messageErr := validate.ResponseErrorValidate(validateErr)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": messageErr,
		})

		return false
	}

	return true
}
