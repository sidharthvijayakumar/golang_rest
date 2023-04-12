package router

import (
	"github.com/anirudhmpai/controllers"
	database "github.com/anirudhmpai/database/sqlc"
	"github.com/anirudhmpai/middleware"
	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	authController controllers.AuthController
	db             *database.Queries
}

func NewAuthRoutes(authController controllers.AuthController, db *database.Queries) AuthRoutes {
	return AuthRoutes{authController, db}
}

func (rc *AuthRoutes) AuthRoute(rg *gin.RouterGroup) {

	router := rg.Group("/auth")
	router.POST("/register", rc.authController.SignUpUser)
	router.POST("/login", rc.authController.SignInUser)
	router.GET("/refresh", rc.authController.RefreshAccessToken)
	router.GET("/logout", middleware.DeserializeUser(rc.db), rc.authController.LogoutUser)
}
