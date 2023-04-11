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
	r := router.Router()
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	fmt.Println("Starting server on the port " + os.Getenv("PORT") + "...")
	log.Fatal(http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), r))
}
