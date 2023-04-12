package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"

	"github.com/anirudhmpai/config"
	"github.com/anirudhmpai/controllers"
	dbConn "github.com/anirudhmpai/database/sqlc"
	"github.com/anirudhmpai/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	server *gin.Engine
	db     *dbConn.Queries
	ctx    context.Context

	AuthController controllers.AuthController
	UserController controllers.UserController
	AuthRoutes     router.AuthRoutes
	UserRoutes     router.UserRoutes
	OtherRoutes    router.OtherRoutes
)

func init() {
	ctx = context.TODO()
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	conn, err := sql.Open(config.PostgreDriver, config.PostgresSource)
	if err != nil {
		log.Fatalf("could not connect to postgres database: %v", err)
	}

	db = dbConn.New(conn)

	fmt.Println("PostgreSQL connected successfully...")

	AuthController = *controllers.NewAuthController(db, ctx)
	UserController = controllers.NewUserController(db, ctx)
	AuthRoutes = router.NewAuthRoutes(AuthController, db)
	UserRoutes = router.NewUserRoutes(UserController, db)

	server = gin.Default()
	server.SetTrustedProxies(nil)

}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.Origin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	errConfig := godotenv.Load()

	if errConfig != nil {
		log.Fatalf("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := server.Group("/api")

	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Welcome to Golang with PostgreSQL"})
	})

	AuthRoutes.AuthRoute(router)
	UserRoutes.UserRoute(router)
	OtherRoutes.OtherRoute(router)

	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": fmt.Sprintf("Route %s not found", ctx.Request.URL)})
	})
	log.Fatal(server.Run(":" + config.Port))

	// r := router.Router()

	// fmt.Println("Starting server on the port " + port + "...")
	// // log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
	// r.Run(":" + port)
}
