package route

import (
	"web/controller"
	"web/controller/auth"
	"web/controller/story"
	"web/controller/user"
	"web/middleware"
	"web/validate"

	"github.com/gin-gonic/gin"
)

func RegisterRoute() {
	router := gin.Default()

	validate.Run()
	router.Use(middleware.Cors)
	router.Use(middleware.ParseMiddleWare)
	router.Static("/assets", "./assets")

	router.GET("init", controller.Init)

	router.POST("/login", controller.Login)
	router.POST("/register", auth.Register)
	router.GET("/logout", auth.Logout)
	router.GET("crawl-series", story.Crawl)
	router.GET("crawl-story", story.ScrawlStory)

	router.POST("user/update-information", user.Update)
	router.POST("user/change-password", user.ChangePassword)

	router.Run()
}
