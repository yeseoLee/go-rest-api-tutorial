package main

import (
	"go-rest-api/router"

	"github.com/gin-gonic/gin"
)

func main() {

	// gin route and run
	r := gin.Default()
	router.InitRoutes(r)
	r.Run("localhost:8080")

}
