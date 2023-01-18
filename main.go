package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/router"

	"github.com/gin-gonic/gin"
)

// @title Example API
// @description 검색창에 버전을 입력하면 해당 버전의 api가 노출됩니다.
// @description 예시) v1.json
// @description 현재 사용가능한 버전 : ###versionList###
// @description 버전 없이 호출 시 v1 함수가 호출 됩니다.
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
	r.Run("localhost:8080")
}
