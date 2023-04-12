package controllers

import (
	"context"
	"net/http"

	database "github.com/anirudhmpai/database/sqlc"
	"github.com/anirudhmpai/users"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	db  *database.Queries
	ctx context.Context
}

func NewUserController(db *database.Queries, ctx context.Context) UserController {
	return UserController{db, ctx}
}

func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(database.User)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": users.FilteredResponse(currentUser)}})
}
