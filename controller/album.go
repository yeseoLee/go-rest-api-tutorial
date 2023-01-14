package controller

import (
	"fmt"
	"go-rest-api/entity"
	"go-rest-api/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	albums, err := model.GetAllAlbums()
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}

// GetAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "please check id value"})
		return
	}

	album, err := model.GetAlbumById(id)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

// PostAlbum adds an album from JSON received in the request body.
func PostAlbum(c *gin.Context) {
	var newAlbum entity.Album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "please check value"})
		return
	}

	// Add the new album to DB
	if err := model.InsertAlbum(&newAlbum); err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "fail to create album"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func PutAlbum(c *gin.Context) {
	var renewAlbum entity.Album

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "please check id value"})
		return
	}

	if err = c.BindJSON(&renewAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "please check request body"})
		return
	}
	renewAlbum.ID = int64(id)

	// Update album to DB
	if err := model.UpdateAlbum(&renewAlbum); err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusCreated, renewAlbum)
}

func DeleteAlbum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "please check id value"})
		return
	}

	// Delete album from DB
	err = model.DeleteAlbum(id)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "delete album complete"})
}
