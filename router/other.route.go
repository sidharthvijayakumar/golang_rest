package router

import (
	"github.com/anirudhmpai/controllers"
	database "github.com/anirudhmpai/database/sqlc"
	"github.com/gin-gonic/gin"
)

type OtherRoutes struct {
	otherController controllers.OtherController
	db              *database.Queries
}

func NewOtherRoutes(authController controllers.AuthController, db *database.Queries) AuthRoutes {
	return AuthRoutes{authController, db}
}

func (rc *OtherRoutes) OtherRoute(rg *gin.RouterGroup) {

	router := rg.Group("/other")

	router.GET("/push", rc.otherController.CreateNotification)
}
