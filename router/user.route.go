package router

import (
	"github.com/anirudhmpai/controllers"
	database "github.com/anirudhmpai/database/sqlc"
	"github.com/anirudhmpai/middleware"
	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userController controllers.UserController
	db             *database.Queries
}

func NewUserRoutes(userController controllers.UserController, db *database.Queries) UserRoutes {
	return UserRoutes{userController, db}
}

func (rc *UserRoutes) UserRoute(rg *gin.RouterGroup) {

	router := rg.Group("/users")
	router.GET("/me", middleware.DeserializeUser(rc.db), rc.userController.GetMe)
}
