package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id   int    `json:"id" example:"1"`      //유저ID
	Name string `json:"name" example:"John"` //이름
	Age  int    `json:"age" example:"10"`    //나이
}

/* 아래 항목이 swagger에 의해 문서화 된다. */
// HelloHandler godoc
// @Summary 요약 기재
// @Description 상세한 설명 기재
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Router /api/v1/hello/{name} [get]
// @Success 200 {object} User
// @Failure 400
func HelloHandler(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"user": ""})
	} else {
		user := User{Id: 1, Name: name, Age: 20}
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}
