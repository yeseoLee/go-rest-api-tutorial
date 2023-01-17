package controller

import (
	"encoding/json"
	"fmt"
	"go-rest-api/oauth"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func Login(c *gin.Context) {
	token := oauth.GetToken()
	url := oauth.GetLoginURL(token)
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Socal Login",
		"url":   url,
	})
}
func LoginCallback(c *gin.Context) {
	code := c.Query("code")
	token, err := oauth.OauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		c.JSON(403, gin.H{"Message": err.Error()})
		return
	}
	fmt.Printf("%+v\n", token)
	response, err := http.Get(oauth.OauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		c.JSON(403, gin.H{"Message": err.Error()})
		return
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(403, gin.H{"Message": err.Error()})
		return
	}
	fmt.Println(string(contents))
	jsonMap := make(map[string]interface{})
	json.Unmarshal(contents, &jsonMap)
	c.JSON(200, jsonMap)
}
