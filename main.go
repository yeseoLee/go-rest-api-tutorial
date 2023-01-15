package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/router"

	"github.com/gin-gonic/gin"
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
	router.InitRoutes(r)
	r.Run("localhost:8080")
}
