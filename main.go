package main

import (
	"fmt"
	v1 "go-rest-api/api/v1"
	"go-rest-api/db"
	"go-rest-api/docs"
	"go-rest-api/router"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	// Set DB
	mysql := db.MySQL{}
	dbconres := mysql.Connect(db.MakeMySQLConnectionString("127.0.0.1", 3306, "root", "123456", "sample-db"))
	if !dbconres {
		fmt.Println("mysql db fail to connect")
	} else {
		fmt.Println("mysql db connect")
	}
	db.SetDatabase(&mysql)

	// gin route and run
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	router.InitRoutes(r)

	// swagger 적용
	docs.SwaggerInfo.Title = "Swagger Example API"
	// 127.0.0.1:8080/docs/index.html
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1Group := r.Group("/api/v1")
	{
		v1Group.GET("/hello/:name", v1.HelloHandler)
	}
	r.Run("localhost:8080")
}
