package router

import (
	"log"

	"github.com/anirudhmpai/albums"
	"github.com/anirudhmpai/users"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Router() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	// load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

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

	// router.Run(":" + os.Getenv("PORT"))
	return router
}
