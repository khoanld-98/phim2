package middleware

import (
	"encoding/json"
	"strings"
	"web/models"
	"web/service/auth"
	"web/store"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func ParseMiddleWare(ctx *gin.Context) {
	if beanToken := ctx.GetHeader("authorization"); beanToken == "" {

		ctx.Set("user", nil)
		ctx.Next()
	} else {
		token := strings.Split(beanToken, " ")
		var user models.Users
		db := store.ConnectDB()
		var personalAccessToken models.PersonalAccessToken
		db.Model(&personalAccessToken).Where("access_token = ?", token[len(token)-1]).First(&personalAccessToken)

		if (personalAccessToken != models.PersonalAccessToken{}) {
			parseValues := auth.ParserJwt(token[len(token)-1])

			jsonData, _ := json.Marshal(parseValues.Claims.(jwt.MapClaims)["User"])
			json.Unmarshal(jsonData, &user)
		}

		ctx.Set("user", user)
	}

	ctx.Next()
}
