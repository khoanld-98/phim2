package user

import (
	"web/service/user"

	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context) {
	var updateService user.UpdateClass

	updateService.Update(ctx)
}
