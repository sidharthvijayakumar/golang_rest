package router

import (
	"github.com/anirudhmpai/albums"
	"github.com/anirudhmpai/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/albums", albums.GetAlbums)
	router.GET("/albums/:id", albums.GetAlbumByID)
	router.POST("/albums", albums.PostAlbums)
	router.DELETE("/albums/:id", albums.DeleteAlbumByID)

	router.GET("/api/user", middleware.GetAllUser)
	router.GET("/api/user/:id", middleware.GetUser)
	router.POST("/api/new-user", middleware.CreateUser)
	router.PUT("/api/user/:id", middleware.UpdateUser)
	router.DELETE("/api/delete-user/:id", middleware.DeleteUser)

	router.Run("localhost:8080")

	return router
}