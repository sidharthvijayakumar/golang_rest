package main

import (
	"log"

	"github.com/anirudhmpai/router"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// r :=
	router.Router()

	// fmt.Println("Starting server on the port " + os.Getenv("PORT") + "...")
	// log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
