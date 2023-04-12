package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/gin-contrib/cors"
	"google.golang.org/api/option"

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

	opt := option.WithCredentialsFile("firebase.json")
	configFirebase := &firebase.Config{ProjectID: "golangrest"}

	app, err := firebase.NewApp(context.Background(), configFirebase, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// This registration token comes from the client FCM SDKs.
	registrationToken := "YOUR_REGISTRATION_TOKEN"

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
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

	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": fmt.Sprintf("Route %s not found", ctx.Request.URL)})
	})
	log.Fatal(server.Run(":" + config.Port))

	// r := router.Router()

	// fmt.Println("Starting server on the port " + port + "...")
	// // log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
	// r.Run(":" + port)
}
