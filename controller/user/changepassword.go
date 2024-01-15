package user

import (
	"web/service/user"

	"github.com/gin-gonic/gin"
)

func ChangePassword(ctx *gin.Context) {
	var ChangePasswordClass user.ChangePasswordClass

	ChangePasswordClass.Handle(ctx)
}
