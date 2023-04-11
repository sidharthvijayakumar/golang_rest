package main

import (
	"github.com/anirudhmpai/albums"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", albums.GetAlbums)
	router.GET("/albums/:id", albums.GetAlbumByID)
	router.POST("/albums", albums.PostAlbums)
	router.DELETE("/albums/:id", albums.DeleteAlbumByID)
	router.Run("localhost:8083")
}
