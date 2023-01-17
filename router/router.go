package router

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.Engine) {
	setAlbumRoutes(r)
	setOauthRoutes(r)
}
