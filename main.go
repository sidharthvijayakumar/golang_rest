package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/anirudhmpai/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := router.Router()

	fmt.Println("Starting server on the port " + os.Getenv("PORT") + "...")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
