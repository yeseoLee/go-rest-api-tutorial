package router

import (
	"go-rest-api/controller"

	"github.com/gin-gonic/gin"
)

func setAlbumRoutes(r *gin.Engine) {
	r.GET("/albums", controller.GetAlbums)
	r.GET("/albums/:id", controller.GetAlbumByID)
	r.POST("/albums", controller.PostAlbum)
	r.PUT("/albums/:id", controller.PutAlbum)
	r.DELETE("/albums/:id", controller.DeleteAlbum)
}
