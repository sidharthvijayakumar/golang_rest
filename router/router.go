package router

import (
	"os"

	"github.com/anirudhmpai/albums"
	"github.com/anirudhmpai/users"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/albums", albums.GetAlbums)
	router.GET("/albums/:id", albums.GetAlbumByID)
	router.POST("/albums", albums.PostAlbums)
	router.DELETE("/albums/:id", albums.DeleteAlbumByID)

	router.GET("/api/user", users.GetAllUser)
	router.GET("/api/user/:id", users.GetUser)
	router.POST("/api/new-user", users.CreateUser)
	router.PUT("/api/user/:id", users.UpdateUser)
	router.DELETE("/api/delete-user/:id", users.DeleteUser)

	router.Run(":" + os.Getenv("PORT"))

	return router
}
