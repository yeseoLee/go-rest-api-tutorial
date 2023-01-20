package router

import (
	"go-rest-api/controller"

	"github.com/gin-gonic/gin"
)

func setOauthRoutes(r *gin.Engine) {
	r.GET("/login", controller.Login)
	r.GET("/login/callback", controller.LoginCallback)
}
