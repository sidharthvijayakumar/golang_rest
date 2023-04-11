package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anirudhmpai/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe("golangrest.up.railway.app:8080", r))
}
