package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	database "github.com/anirudhmpai/database/sqlc"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

type OtherController struct {
	db  *database.Queries
	ctx context.Context
}

func NewOtherController(db *database.Queries, ctx context.Context) OtherController {
	return OtherController{db, ctx}
}

// ? CreateNotification creates a notification
func (uc *OtherController) CreateNotification(ctx *gin.Context) {
	opt := option.WithCredentialsFile("firebase.json")
	configFirebase := &firebase.Config{ProjectID: "golangrest"}

	app, err := firebase.NewApp(context.Background(), configFirebase, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	// Obtain a messaging.Client from the App.
	c := context.Background()
	client, err := app.Messaging(c)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// This registration token comes from the client FCM SDKs.
	registrationToken := "d8ih4W1tRLSC8Z8pwy5fA9:APA91bGhMCbV5DuoDD0CzFpRjdjtXkrvMpr1PUWg7PmnlA0cSyBr75_hkT_jLPsSSJVln9I80mis9pObAvEvgoBwOW3UMlnqboKylh3eD_IuY6RU7xBJ6hwUpRqMhMwLebwWGf3Nstf2"

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"title": "Welcome to golang push Notifications",
			"body":  "via railway app",
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
	// send the middleware.Response
	ctx.IndentedJSON(http.StatusOK, response)
}
