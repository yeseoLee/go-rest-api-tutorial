package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setSwaggerRoutes(r *gin.Engine) {
	// swagger 적용 (127.0.0.1:8080/docs/index.html)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
